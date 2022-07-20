# go-simple_app

Simple web server that responds on port 8888 with a message and the contents of a JSON file.

You can override the default JSON file with the `JSON_FILE` environment variable.

## Kubernetes deployment file

The file `deployment-simple-app.yaml` can be used to deploy this app in k8s using the [Vault Agent Sidecar Injector](https://www.vaultproject.io/docs/platform/k8s/injector) to generate the JSON file. Check the [sidecar annotations reference](https://www.vaultproject.io/docs/platform/k8s/injector/annotations) to understand the annotations used.

This is useful for testing Kubernetes authentication when setting up [HashiCorp Vault](https://vaultproject.io/).
