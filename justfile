set dotenv-load

export runtime := env('CONTAINER_RUNTIME', 'docker')
export registry := env('IMAGE_REGISTRY')

default:
  just --list

# Push images to the $IMAGE_REGISTRY
push-images: build-images
  {{runtime}} push {{registry}}/instaslice-operator
  {{runtime}} push {{registry}}/instaslice-daemonset

# Build images using the $IMAGE_REGISTRY
build-images:
  IMAGE_REGISTRY={{registry}} make images

