package instaslice

import (
	"context"

	clientset "github.com/openshift/instaslice-operator/pkg/generated/clientset/versioned"

	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

type InstasliceController struct {
	namespace          string
	instasliceClient   *clientset.Clientset
	instasliceInformer cache.SharedInformer
}

func NewInstasliceController(namespace string, operatorClient *clientset.Clientset, instasliceInformer cache.SharedInformer, eventRecorder events.Recorder) factory.Controller {
	c := &InstasliceController{
		namespace:          namespace,
		instasliceClient:   operatorClient,
		instasliceInformer: instasliceInformer,
	}

	return factory.New().
		WithSync(c.sync).
		WithInformersQueueKeysFunc(c.nameToKey, instasliceInformer).
		ToController("InstasliceController", eventRecorder)
}

func (c *InstasliceController) sync(ctx context.Context, syncCtx factory.SyncContext) error {
	instasliceName := syncCtx.QueueKey()

	klog.V(2).InfoS("Instaslice Sync", "queue_key", instasliceName)

	slice, err := c.instasliceClient.OpenShiftOperatorV1alpha1().Instaslices(c.namespace).Get(ctx, instasliceName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	klog.V(2).InfoS("Slice", "slice", slice)

	return nil
}

// queueKeysRuntimeForObj is an adapter on top of queueKeysForObj to be used in
// factory.Controller queueing functions
func (c *InstasliceController) nameToKey(obj runtime.Object) []string {
	metaObj, ok := obj.(metav1.ObjectMetaAccessor)
	if !ok {
		klog.Errorf("the object is not a metav1.ObjectMetaAccessor: %T", obj)
		return []string{}
	}
	return []string{metaObj.GetObjectMeta().GetName()}
}
