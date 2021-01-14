# Exercise 2
This exercise uses go-swagger to generate the http service for this project
to generate the server in the case of swagger.yaml changes run
```bash
swagger generate server -f swagger.yaml
```

## Prerequisites
You'll need a few tools to run the demo locally. They are listed below.

### Chocolately
In Windows machine you need to install [Chocolately](https://chocolatey.org/install)

### Docker
If your development machine runs windows, install the
[Docker for desktop windows application](https://docs.docker.com/docker-for-windows/). If you're on Mac install the
[Mac version](https://docs.docker.com/docker-for-mac/). When you're running Linux you'll need to install and run
[Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/). Skaffold should automatically detect which of
these you're using deploy accordingly. If it doesn't you can add a you can set the config value `deploy.kubeContext` in
`skaffold.yaml` to either `docker-for-desktop` or `minikube`, depending on which one you're using.

* Windows
```bash
choco install minikube
```

### Skaffold
[Skaffold](https://skaffold.dev/docs/install/) will build and deploy the required containers to a local cluster.

* Windows
```bash
choco install -y skaffold
```

### Kustomize
Skaffold will use [Kustomize](https://github.com/kubernetes-sigs/kustomize/blob/master/docs/INSTALL.md) for
deployments. It needs to be installed on your machine.


# demo 
Once minikube is running
```bash
minikube run
```
Run
```bash
skaffold dev --port-forward
```
This will kick off skaffold and forward the services to your local machine. 
The demo service will be available locally on the port listed in the output.

Example api call
```bash
curl http://localhost:8080/base64/decode -d "$(echo hello world | base64)" -H "Content-type: text/plain"
```

