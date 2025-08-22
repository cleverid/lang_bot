package user

import(
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"
)

type Client struct {
    UserClient
    config ClientConfig
    connection *grpc.ClientConn
    opts []grpc.DialOption
}

func NewClient(config ClientConfig) (*Client, error) {
	client := &Client{
        config: config,
    }
	client.opts = append(client.opts, 
        grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
	return client, nil
}

func (c *Client) Start() (err error) {
	c.connection, err = grpc.NewClient(c.config.Host, c.opts...)
    if err != nil {
        return err
    }
	c.UserClient = NewUserClient(c.connection)
    return nil
}

func (c *Client) Stop() (err error) {
    return c.connection.Close()
}
