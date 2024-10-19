### Prerequisites:
- Go
- Helm
- Tilt
- K3d
- Docker

1. Setup the k3d cluster + local image registry
```
task bootstrap
```

2. Setup the code + application into the cluster + enable hot reload for testing
```
task dev
```

3. Clean up all resources from this project
```
task clean
```
