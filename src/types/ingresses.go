package types

// Ingresses struct
type Ingresses struct {
	Name string `json: "name"`
	Namespace string `json: "namespace"`
}

// List Ingresses Response
type ListIngressesResponse struct {
	Ingresses []Ingresses `json: "pods"`
}
