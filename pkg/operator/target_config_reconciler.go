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

	instasliceoperatorv1clientset "github.com/openshift/instaslice-operator/pkg/generated/clientset/versioned/typed/instasliceoperator/v1"
	operatorclientv1informers "github.com/openshift/instaslice-operator/pkg/generated/informers/externalversions/instasliceoperator/v1"

	"github.com/openshift/instaslice-operator/pkg/operator/operatorclient"
	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
)

const (
	WebhookCertificateSecretName  = "webhook-server-cert"
	WebhookCertificateName        = "lws-serving-cert"
	CertManagerInjectCaAnnotation = "cert-manager.io/inject-ca-from"
)

type TargetConfigReconciler struct {
	targetImage                string
	operatorClient             instasliceoperatorv1clientset.InstasliceOperatorInterface
	dynamicClient              dynamic.Interface
	discoveryClient            discovery.DiscoveryInterface
	instasliceoperatorClient   *operatorclient.InstasliceOperatorSetClient
	kubeClient                 kubernetes.Interface
	apiextensionClient         *apiextclientv1.Clientset
	eventRecorder              events.Recorder
	kubeInformersForNamespaces v1helpers.KubeInformersForNamespaces
	secretLister               v1.SecretLister
	namespace                  string
	resourceCache              resourceapply.ResourceCache
}

func NewTargetConfigReconciler(
	targetImage string,
	namespace string,
	operatorConfigClient instasliceoperatorv1clientset.InstasliceOperatorInterface,
	operatorClientInformer operatorclientv1informers.InstasliceOperatorInformer,
	kubeInformersForNamespaces v1helpers.KubeInformersForNamespaces,
	instasliceoperatorClient *operatorclient.InstasliceOperatorSetClient,
	dynamicClient dynamic.Interface,
	discoveryClient discovery.DiscoveryInterface,
	kubeClient kubernetes.Interface,
	apiExtensionClient *apiextclientv1.Clientset,
	eventRecorder events.Recorder,
) factory.Controller {
	c := &TargetConfigReconciler{
		operatorClient:             operatorConfigClient,
		dynamicClient:              dynamicClient,
		instasliceoperatorClient:   instasliceoperatorClient,
		kubeClient:                 kubeClient,
		discoveryClient:            discoveryClient,
		apiextensionClient:         apiExtensionClient,
		eventRecorder:              eventRecorder,
		kubeInformersForNamespaces: kubeInformersForNamespaces,
		secretLister:               kubeInformersForNamespaces.SecretLister(),
		targetImage:                targetImage,
		namespace:                  namespace,
		resourceCache:              resourceapply.NewResourceCache(),
	}

	return factory.New().WithInformers(
		// for the operator changes
		operatorClientInformer.Informer(),
		// for the deployment and its configmap and secret
		kubeInformersForNamespaces.InformersFor(namespace).Apps().V1().Deployments().Informer(),
		kubeInformersForNamespaces.InformersFor(namespace).Core().V1().ConfigMaps().Informer(),
		kubeInformersForNamespaces.InformersFor(namespace).Core().V1().Secrets().Informer(),
	).ResyncEvery(time.Minute*5).
		WithSync(c.sync).
		WithSyncDegradedOnError(instasliceoperatorClient).
		ToController("TargetConfigController", eventRecorder)
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
