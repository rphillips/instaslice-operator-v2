# Building

## Create an .env file

The contents of the file overload the IMAGE_REGISTRY environment variable.

```
IMAGE_REGISTRY=quay.io/some.user
```

## Building and Deploying to Kind

```bash
just push-images create-kind apply-cert-manager apply-crds
```

