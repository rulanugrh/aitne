## Aitne
Tools for management k8s in CLI and management k8s with API resources

## Usage
Must export env for kubeconfig path
```
export KUBECONFIG_PATH=$HOME/.kube/config
```

To display command you can use this
```bash
$ go run cmd/cmd.go help
```

If want to get data about k8s fiture like deployment, daemon, etc. You can use this command
```bash
$ go run cmd/cmd.go get deployment
```

If want to get data by name you can use this
```bash
$ go run cmd/cmd.go catch deployment -name=sample-deployment
```

If want to delete by name you can use this
```bash
$ go run cmd/cmd.go delete deployment -name=sample-deployment
```

## Running API Services
```bash
$ go run api/api.go
```

> This project still working for better usage
