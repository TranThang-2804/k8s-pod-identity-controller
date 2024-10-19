### Prerequisites:
- Go
- Tilt
- Docker

> Tools installed by running ```task install``` if they are not installed
> - Helm
> - K3d

### Command

1. Install required tools
```
task install
```

2. Setup the k3d cluster + local image registry
```
task bootstrap
```

3. Setup the code + application into the cluster + enable hot reload for testing
```
task dev
```

4. Clean up all resources from this project
```
task clean
```
