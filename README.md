# Building

## Install imagebuilder

```
go install github.com/openshift/imagebuilder/cmd/imagebuilder@latest
```

## Create an .env file

The contents of the file overload the IMAGE_REGISTRY environment variable.

```
IMAGE_REGISTRY=quay.io/some.user
```

## Building and Deploying to Kind

```bash
OSS=1 just push-images create-kind apply-cert-manager apply-crds
```


## Building and Deploying to any other cluster

```bash
just push-images apply-cert-manager apply-crds
```

## Environment Flags

* OSS=1 build with upstream images, otherwise build with OpenShift images
* IMAGE_REGISTRY=quay.io/some.user : overloads the image registry destination

