# Infra-gw

The ingra-gw service provides us with higher level access to the kubernetes cluster. E.g.listing/creating/deleting resources

## Run 

`go run main.go`

### Flags 

```shell
-kubeconfig string
    Path to kubeconfig (~/.kube/config if empty)
-listenAddress string
    Address to bind web server to (default "127.0.0.1:4444")
-inCluster bool
    Use in-cluster mode - use mounted Kubernetes credentials instead of kube config (default false)
```

## Endpoint

* ### GET

    - [X] `/ping`               - Returns service status
    - [X] `/pods[?ns="foo"]`    - Lists all pods for namespace, for all if empty
    - [X] `/deployments[?ns="foo"]` - Lists all deployments for namespace, for all if empty
    - [X] `/services[?ns="foo"]` Lists all services for namespace, for all if empty
    - [X] `/ingresses[?ns="foo]` Lists all ingresses for namespace, for all if empty
    - [X] `/namespaces` Lists all namespaces  

* ### POST

    - [X] `/namespaces` Creates new namespace with name
        ```json
        {
            "name": "namespace-name"
        }
        ```
    - [X] `/apps/mysql` Creates new mysql instance in provided options
        ```json
        {
            "name": "mysql-name",
            "namespace": "namespace-name"
        }
        ```
    - [X] `apps/redis` Creates new redis instance with provided options
        ```json
        {
            "name": "redis-name",
            "namespace": "namespace-name"
        }
        ```

* ### DELETE
    - [X] `/namespaces"` Deletes namespaces by name
        ```json
        {
            "name": "namespace-name"
        }
        ```


## In-Cluster
In this mode, `k8s-gw` gets credentials from he mounted service account.
Make sure to create a service account and assign it the necessary RBAC permissions (see k8s/
dir). This mode is enabled by a `-inCluster` flag.

## Docker

Run following command to build docker image

```shell
docker build -t infra-k8s-gw -f docker/Dockerfile . 
```

To run docker container use following command

```shell
docker run -it -d --rm --name k8s-infra-gw -p 4444:4444 infra-k8s-gw
```
