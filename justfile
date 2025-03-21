set dotenv-load

export runtime := env('CONTAINER_RUNTIME', 'docker')
export registry := env('IMAGE_REGISTRY')
export imgtag := env('IMAGE_TAG', 'latest')

export kind_cluster := "slice-cluster"
export kind_image := "kindest/node:v1.31.2"

default:
  just --list

# Push images to the $IMAGE_REGISTRY
images-push: images-build
  #!/usr/bin/env bash
  
  if [[ "{{runtime}}" == "podman" ]]; then
    echo "Copying image from podman root: operator"
    {{runtime}} pull docker-daemon:{{registry}}/instaslice-operator:{{imgtag}}
    echo "Copying image from podman root: daemonset"
    {{runtime}} pull docker-daemon:{{registry}}/instaslice-daemonset:{{imgtag}}
    echo "Copying image from podman root: webhook"
    {{runtime}} pull docker-daemon:{{registry}}/instaslice-webhook:{{imgtag}}
  fi

  {{runtime}} push {{registry}}/instaslice-operator:{{imgtag}}
  {{runtime}} push {{registry}}/instaslice-daemonset:{{imgtag}}
  {{runtime}} push {{registry}}/instaslice-webhook:{{imgtag}}

# Build images using the $IMAGE_REGISTRY
images-build:
  IMAGE_REGISTRY={{registry}} make images

# Create a local kind cluster
kind-create:
  #!/usr/bin/env bash

  set -e -o pipefail

  # check to see if the cluster is running
  if kind get kubeconfig --name {{kind_cluster}} &>/dev/null; then
    echo "kind cluster {{kind_cluster}} is running"
    exit 0
  fi


  kind create cluster -n {{kind_cluster}} --image {{kind_image}}
  kubectl cluster-info --context kind-{{kind_cluster}}

# Delete the local kind cluster
kind-delete:
  #!/usr/bin/env bash

  set -e -o pipefail

  kind delete cluster -n {{kind_cluster}}

# Apply CRDs
apply-crds:
  #!/usr/bin/env bash

  set -e -o pipefail
  
  find deploy/*.yaml | grep -v 09_ | xargs -n 1 kubectl apply -f
  sleep 2
  kubectl apply -f deploy/09_*
  sleep 2

# install cert-manager
apply-cert-manager:
  kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.17.0/cert-manager.yaml

wait-for-operator:
  kubectl wait --for=condition=ready pod -l app=instaslice-operator -n instaslice-system --timeout=300s

# run e2e's
e2e-kind path="./test/e2e": kind-create apply-cert-manager apply-crds wait-for-operator && kind-delete
  #!/usr/bin/env bash

  echo "Run E2E's on Kind"

  exit 0
