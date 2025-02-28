package instaslice

import (
	"context"

	operatorclientv1alpha1informers "github.com/openshift/instaslice-operator/pkg/generated/informers/externalversions/instasliceoperator/v1alpha1"
	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/klog/v2"
)

type InstasliceController struct {
	instasliceInformer operatorclientv1alpha1informers.InstasliceInformer
}

func NewInstasliceController(instasliceInformer operatorclientv1alpha1informers.InstasliceInformer, eventRecorder events.Recorder) factory.Controller {
	c := &InstasliceController{
		instasliceInformer: instasliceInformer,
	}

	return factory.New().
		WithSync(c.sync).
		WithInformersQueueKeysFunc(c.nameToKey, instasliceInformer.Informer()).
		ToController("InstasliceController", eventRecorder)
}

func (c *InstasliceController) sync(ctx context.Context, syncCtx factory.SyncContext) error {
	klog.V(2).InfoS("Instaslice Sync", "queue_key", syncCtx.QueueKey())
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
