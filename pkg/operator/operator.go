package operator

import (
	"context"
	"os"
	"time"

	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"github.com/openshift/library-go/pkg/operator/loglevel"
	"github.com/openshift/library-go/pkg/operator/v1helpers"

	apiextclientv1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"

	operatorconfigclient "github.com/openshift/instaslice-operator/pkg/generated/clientset/versioned"
	operatorclientinformers "github.com/openshift/instaslice-operator/pkg/generated/informers/externalversions"
	instaslicecontroller "github.com/openshift/instaslice-operator/pkg/operator/controllers/instaslice"
	"github.com/openshift/instaslice-operator/pkg/operator/operatorclient"
)

var operatorNamespace = "instaslice-system"

func RunOperator(ctx context.Context, cc *controllercmd.ControllerContext) error {
	kubeClient, err := kubernetes.NewForConfig(cc.ProtoKubeConfig)
	if err != nil {
		return err
	}

	dynamicClient, err := dynamic.NewForConfig(cc.ProtoKubeConfig)
	if err != nil {
		return err
	}

	apiextensionClient, err := apiextclientv1.NewForConfig(cc.KubeConfig)
	if err != nil {
		return err
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(cc.KubeConfig)
	if err != nil {
		return err
	}

	operatorConfigClient, err := operatorconfigclient.NewForConfig(cc.KubeConfig)
	if err != nil {
		return err
	}
	operatorConfigInformers := operatorclientinformers.NewSharedInformerFactory(operatorConfigClient, 10*time.Minute)

	namespace := cc.OperatorNamespace
	if namespace == "openshift-config-managed" {
		// we need to fall back to our default namespace rather than library-go's when running outside the cluster
		namespace = operatorNamespace
	}

	kubeInformersForNamespaces := v1helpers.NewKubeInformersForNamespaces(kubeClient, "", namespace)

	instasliceClient := &operatorclient.InstasliceOperatorSetClient{
		Ctx:               ctx,
		SharedInformer:    operatorConfigInformers.OpenShiftOperator().V1alpha1().InstasliceOperators().Informer(),
		Lister:            operatorConfigInformers.OpenShiftOperator().V1alpha1().InstasliceOperators().Lister(),
		OperatorClient:    operatorConfigClient.OpenShiftOperatorV1alpha1(),
		OperatorNamespace: namespace,
	}

	targetConfigReconciler := NewTargetConfigReconciler(
		os.Getenv("RELATED_IMAGE_OPERAND_IMAGE"),
		namespace,
		operatorConfigClient.OpenShiftOperatorV1alpha1().InstasliceOperators(namespace),
		operatorConfigInformers.OpenShiftOperator().V1alpha1().InstasliceOperators(),
		kubeInformersForNamespaces,
		instasliceClient,
		dynamicClient,
		discoveryClient,
		kubeClient,
		apiextensionClient,
		cc.EventRecorder,
	)

	logLevelController := loglevel.NewClusterOperatorLoggingController(instasliceClient, cc.EventRecorder)
	instasliceController := instaslicecontroller.NewInstasliceController(operatorConfigInformers.OpenShiftOperator().V1alpha1().Instaslices(), cc.EventRecorder)

	klog.Infof("Starting informers")
	operatorConfigInformers.Start(ctx.Done())
	kubeInformersForNamespaces.Start(ctx.Done())

	klog.Infof("Starting log level controller")
	go logLevelController.Run(ctx, 1)
	klog.Infof("Starting target config reconciler")
	go targetConfigReconciler.Run(ctx, 1)
	klog.Infof("Starting Instaslice Controller")
	go instasliceController.Run(ctx, 1)

	<-ctx.Done()
	return nil
}
