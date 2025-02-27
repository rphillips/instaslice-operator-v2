package instaslice

import (
	"context"

	operatorclientv1alpha1informers "github.com/openshift/instaslice-operator/pkg/generated/informers/externalversions/instasliceoperator/v1alpha1"
	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"

	"k8s.io/klog/v2"
)

type InstasliceController struct {
	instasliceInformer operatorclientv1alpha1informers.InstasliceInformer
}

func NewInstasliceController(instasliceInformer operatorclientv1alpha1informers.InstasliceInformer, eventRecorder events.Recorder) factory.Controller {
	c := &InstasliceController{
		instasliceInformer: instasliceInformer,
	}

	return factory.New().WithInformers(instasliceInformer.Informer()).
		WithSync(c.sync).
		ToController("InstasliceController", eventRecorder)
}

func (c *InstasliceController) sync(ctx context.Context, syncCtx factory.SyncContext) error {
	klog.V(2).Info("Instaslice Sync")
	return nil
}
