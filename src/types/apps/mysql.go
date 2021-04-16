package types

type CreateMySQLRequest struct {
	Name string `json: "name"`
	Namespace string `json: "namespace"`
}

type CreateMySQLResponse struct {
	Status int `json: "status"`
	Message string `json: "message"`
}
