load('ext://restart_process', 'docker_build_with_restart')

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/cloud-role-controller ./'

local_resource(
  'cloud-role-controller-compile',
  compile_cmd,
  deps=['./main.go'],
)

docker_build_with_restart(
  'k3d-localregistry.localhost:12345/cloud-role-identity-image',
  '.',
  entrypoint=['/app/build/cloud-role-controller'],
  dockerfile='local/Dockerfile',
  only=[
    './build',
  ],
  live_update=[
    sync('./build', '/app/build'),
  ],
)

k8s_yaml('local/deployment.yaml')
k8s_resource('cloud-role-identity-controller', port_forwards=8080,
             resource_deps=['cloud-role-controller-compile'])
