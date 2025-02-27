package operator

import (
	"context"
	"fmt"
	"time"

	apiextclientv1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"

	instasliceoperatorv1alphaclientset "github.com/openshift/instaslice-operator/pkg/generated/clientset/versioned/typed/instasliceoperator/v1alpha1"
	operatorclientv1alpha1informers "github.com/openshift/instaslice-operator/pkg/generated/informers/externalversions/instasliceoperator/v1alpha1"

	"github.com/openshift/instaslice-operator/pkg/operator/operatorclient"
	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
)

const (
	WebhookCertificateSecretName  = "webhook-server-cert"
	WebhookCertificateName        = "instaslice-serving-cert"
	CertManagerInjectCaAnnotation = "cert-manager.io/inject-ca-from"
)

type TargetConfigReconciler struct {
	apiextensionClient         *apiextclientv1.Clientset
	discoveryClient            discovery.DiscoveryInterface
	dynamicClient              dynamic.Interface
	eventRecorder              events.Recorder
	instasliceoperatorClient   *operatorclient.InstasliceOperatorSetClient
	instasliceInformer         operatorclientv1alpha1informers.InstasliceInformer
	kubeClient                 kubernetes.Interface
	kubeInformersForNamespaces v1helpers.KubeInformersForNamespaces
	namespace                  string
	operatorClient             instasliceoperatorv1alphaclientset.InstasliceOperatorInterface
	resourceCache              resourceapply.ResourceCache
	secretLister               v1.SecretLister
	targetImage                string
}

func NewTargetConfigReconciler(
	targetImage string,
	namespace string,
	operatorConfigClient instasliceoperatorv1alphaclientset.InstasliceOperatorInterface,
	operatorClientInformer operatorclientv1alpha1informers.InstasliceOperatorInformer,
	kubeInformersForNamespaces v1helpers.KubeInformersForNamespaces,
	instasliceoperatorClient *operatorclient.InstasliceOperatorSetClient,
	dynamicClient dynamic.Interface,
	discoveryClient discovery.DiscoveryInterface,
	kubeClient kubernetes.Interface,
	apiExtensionClient *apiextclientv1.Clientset,
	eventRecorder events.Recorder,
) factory.Controller {
	c := &TargetConfigReconciler{
		apiextensionClient:         apiExtensionClient,
		discoveryClient:            discoveryClient,
		dynamicClient:              dynamicClient,
		eventRecorder:              eventRecorder,
		instasliceoperatorClient:   instasliceoperatorClient,
		kubeClient:                 kubeClient,
		kubeInformersForNamespaces: kubeInformersForNamespaces,
		namespace:                  namespace,
		operatorClient:             operatorConfigClient,
		resourceCache:              resourceapply.NewResourceCache(),
		secretLister:               kubeInformersForNamespaces.SecretLister(),
		targetImage:                targetImage,
	}

	return factory.New().WithInformers(
		// for the operator changes
		operatorClientInformer.Informer(),
		// for the deployment and its configmap, secret, daemonsets.
		kubeInformersForNamespaces.InformersFor(namespace).Apps().V1().Deployments().Informer(),
		kubeInformersForNamespaces.InformersFor(namespace).Apps().V1().DaemonSets().Informer(),
		kubeInformersForNamespaces.InformersFor(namespace).Core().V1().ConfigMaps().Informer(),
		kubeInformersForNamespaces.InformersFor(namespace).Core().V1().Secrets().Informer(),
	).ResyncEvery(time.Minute*5).
		WithSync(c.sync).
		WithSyncDegradedOnError(instasliceoperatorClient).
		ToController("InstasliceOperatorController", eventRecorder)
}

func (c *TargetConfigReconciler) sync(ctx context.Context, syncCtx factory.SyncContext) error {
	found, err := isResourceRegistered(c.discoveryClient, schema.GroupVersionKind{
		Group:   "cert-manager.io",
		Version: "v1",
		Kind:    "Issuer",
	})
	if err != nil {
		return fmt.Errorf("unable to check cert-manager is installed: %w", err)
	}

	if !found {
		return fmt.Errorf("please make sure that cert-manager is installed on your cluster")
	}

	_, err = c.operatorClient.Get(ctx, operatorclient.OperatorConfigName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("unable to get operator configuration %s/%s: %w", c.namespace, operatorclient.OperatorConfigName, err)
	}

	return err
}

func isResourceRegistered(discoveryClient discovery.DiscoveryInterface, gvk schema.GroupVersionKind) (bool, error) {
	apiResourceLists, err := discoveryClient.ServerResourcesForGroupVersion(gvk.GroupVersion().String())
	if err != nil {
		if apierrors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	for _, apiResource := range apiResourceLists.APIResources {
		if apiResource.Kind == gvk.Kind {
			return true, nil
		}
	}
	return false, nil
}
