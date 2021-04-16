package types

//Deployments struct
type Deployments struct {
	Name string `json: "name"`
	Namespace string `json: "namespace"`
}

// List Deployment Response
type ListDeploymentsResponse struct {
	Deployments []Deployments `json: "deployments"`
}
