package types

// Pods struct
type Pods struct {
	Name string `json: "name"`
	Namespace string `json: "namespace"`
}

// List Pods Response
type ListPodsResponse struct {
	Pods []Pods `json: "pods"`
}
