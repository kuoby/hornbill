package mongoose

import "context"

type Server struct {
	UUID string `json:"uuid"`
}

// Servers returns a list of the servers' information that are present on Mongoose.
// Using dummy data for this function for now.
func (c *Client) Servers(ctx context.Context) ([]Server, error) {
	return []Server{
		{
			UUID: "2003df06-8a35-4953-86ec-73d87bbc0bac",
		},
	}, nil
}

// Server returns a server information that is present on Mongoose.
// Using dummy data for this function for now.
func (c *Client) Server(ctx context.Context) (Server, error) {
	return Server{
		UUID: "2003df06-8a35-4953-86ec-73d87bbc0bac",
	}, nil
}
