package types

// Namespaces struct
type Namespaces struct {
	Name string `json: "name"`
}

// List Namespaces struct
type ListNamespacesResponse struct {
	Namespaces []Namespaces `json: "namespaces"`
}

// Create Namespaces Request
type CreateNamespacesRequest struct {
	Name string `json: "name"`
}

// Create Namespaces Response
type CreateNamespacesResponse struct {
	Message string `json: "message"`
	Status int `json: "status"`
}

//Delete Namespaces Reqest
type DeleteNamespacesRequest struct {
	Name string `json: "name"`
}

//Delete Namespaces Response
type DeleteNamespacesResponse struct {
	Message string `json: "message"`
	Status int `json: "status"`
}
