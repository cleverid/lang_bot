package grpc

type Service struct {
	Name      string   `json:"name"`
	Clients   []Client `json:"clients,omitempty"`
	Contracts Contract
}
type Client struct {
	Service string `json:"service"`
}
type Contract struct {
	GRPC GRPC
}
type GRPC struct {
	Files []string
}
