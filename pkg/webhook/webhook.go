package webhook

import (
	"encoding/json"
	"net/http"

	admissionctl "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	corev1 "k8s.io/api/core/v1"
)

const (
	URI         string = "/mutate-pod"
	WebhookName string = "instaslice-webhook"
)

// Webhook interface
type Webhook interface {
	// Authorized will determine if the request is allowed
	Authorized(request admissionctl.Request) admissionctl.Response
	// GetURI returns the URI for the webhook
	GetURI() string
	// Name is the name of the webhook
	Name() string
}

type InstasliceWebhook struct{}

func NewWebhook() Webhook {
	return &InstasliceWebhook{}
}

// GetURI implements Webhook interface
func (s *InstasliceWebhook) GetURI() string { return URI }

// Name implements Webhook interface
func (s *InstasliceWebhook) Name() string { return WebhookName }

// Validate if the incoming request even valid
func (s *InstasliceWebhook) Validate(req admissionctl.Request) bool {
	return req.Kind.Kind == "Pod" || req.Kind.Kind == "SharedSecret" || req.Kind.Kind == "SharedConfigMap"
}

// Authorized implements Webhook interface
func (s *InstasliceWebhook) Authorized(request admissionctl.Request) admissionctl.Response {
	var ret admissionctl.Response

	pod, err := s.renderPod(request)
	if err != nil {
		log.Error(err, "couldn't render a Pod from the incoming request")
		ret = admissionctl.Errored(http.StatusBadRequest, err)
		ret.UID = request.AdmissionRequest.UID
		return ret
	}

	mutatePod, err := s.mutatePod(pod)
	if err != nil {
		log.Error(err, "could not mutate pod")
		ret = admissionctl.Errored(http.StatusBadRequest, err)
		ret.UID = request.AdmissionRequest.UID
		return ret
	}

	ret = admissionctl.PatchResponseFromRaw(request.Object.Raw, mutatePod)
	ret.UID = request.AdmissionRequest.UID
	return ret
}

func (s *InstasliceWebhook) mutatePod(pod *corev1.Pod) ([]byte, error) {
	mutatedPod := pod.DeepCopy()
	// TODO mutate pod
	return json.Marshal(mutatedPod)
}

// renderPod decodes an *corev1.Pod from the incoming request.
// If the request includes an OldObject (from an update or deletion), it will be
// preferred, otherwise, the Object will be preferred.
func (s *InstasliceWebhook) renderPod(request admissionctl.Request) (*corev1.Pod, error) {
	var err error
	decoder := admissionctl.NewDecoder(scheme)
	pod := &corev1.Pod{}
	if len(request.OldObject.Raw) > 0 {
		err = decoder.DecodeRaw(request.OldObject, pod)
	} else {
		err = decoder.DecodeRaw(request.Object, pod)
	}

	return pod, err
}
