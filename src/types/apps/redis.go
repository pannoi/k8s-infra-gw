package types

// Create Redis Request
type CreateRedisRequest struct {
	Name string `json: "name"`
	Namespace string `json: "namespace"`
}

// Create Redis Response
type CreateRedisResponse struct {
	Status int `json: "status"`
	Message string `json: "message"`
}
