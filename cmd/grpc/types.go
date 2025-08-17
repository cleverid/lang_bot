package grpc

type Service struct {
	Clients []Client `json:"clients,omitempty"`
}

type Client struct {
	Name string `json:"name"`
}
