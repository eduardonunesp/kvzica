package client

import (
	"github.com/eduardonunesp/kvzika/pkg/proto"
	"google.golang.org/grpc"
)

type Client struct {
	client proto.KeyValueServiceClient
	conn   *grpc.ClientConn
}
