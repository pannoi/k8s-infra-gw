package types

type Services struct {
	Name string `json: "name"`
	Namespace string `json: "namespace"`
}

// List Services Response
type ListServicesResponse struct {
	Services []Services `json: "services"`
}
