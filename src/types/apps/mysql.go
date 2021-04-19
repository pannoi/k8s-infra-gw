package types

//Create MySQL Request
type CreateMySQLRequest struct {
	Name string `json: "name"`
	Namespace string `json: "namespace"`
}

// Create MySQL Response
type CreateMySQLResponse struct {
	Status int `json: "status"`
	Message string `json: "message"`
	DatabaseUsername string `json: "username"`
	DatabasePassword map[string][]byte `json: "password"`
}
