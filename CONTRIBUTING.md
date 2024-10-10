### Prerequisites:
- Go
- Helm
- Tilt
- Skaffold
- K3d
- Docker

1. Setup the k3d cluster + local image registry
```
task create-cluster
```

2. Setup the code + application into the cluster + enable hot reload for testing
```
task dev
```
