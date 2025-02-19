package search

import (
	"context"

	"github.com/paulja/go-arch/proto/search/v1"
	"github.com/paulja/go-arch/web/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SearchClient struct {
	ctx    context.Context
	conn   *grpc.ClientConn
	client search.SearchServiceClient
}

func NewSearchClient() *SearchClient {
	return &SearchClient{
		ctx: context.Background(),
	}
}

func (c *SearchClient) Connect() error {
	conn, err := grpc.NewClient(
		config.GetSeachAddr(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	c.conn = conn
	c.client = search.NewSearchServiceClient(conn)
	return nil
}

func (c *SearchClient) Close() error {
	return c.conn.Close()
}

func (c *SearchClient) FindUsers(exp string) ([]string, error) {
	resp, err := c.client.FindUsers(c.ctx, &search.FindUsersRequest{
		Expression: exp,
	})
	if err != nil {
		return nil, err
	}
	return resp.Users, nil
}
