package client

import (
	"context"
	"fmt"
	"log"

	"github.com/eduardonunesp/kvzika/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(host string, port int) (*Client, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: proto.NewKeyValueServiceClient(conn),
		conn:   conn,
	}, err
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) SetKeyValue(key, value []byte) error {
	_, err := c.client.SetKeyValue(
		context.Background(),
		&proto.KeyValueRequest{
			Key:   key,
			Value: value,
		},
	)
	return err
}

func (c *Client) GetValue(key []byte) ([]byte, error) {
	resp, err := c.client.GetValue(
		context.Background(),
		&proto.KeyRequest{
			Key: key,
		},
	)
	if err != nil {
		return nil, err
	}

	log.Println(resp)

	return resp.Value, nil
}
