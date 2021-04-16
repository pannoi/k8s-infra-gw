# Infra-gw

The ingra-gw service provides us with higher level access to the kubernetes cluster. E.g.listing/creating/deleting resources

## Run 

`go run main.go` or `make run`

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
    - [X] `/apps/mysql` Creates new mysql instance
        ```json
        {
            "name" : "mysql-name",
            "namespace": "namespace-name"
        }
        ```

* ### DELETE
    - [X] `/namespaces/{namespace_name}"` Deletes namespaces by name


## In-Cluster
In this mode, `k8s-gw` gets credentials from he mounted service account.
Make sure to create a service account and assign it the necessary RBAC permissions (see k8s/
dir). This mode is enabled by a `-inCluster` flag.

