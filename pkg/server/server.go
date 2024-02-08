package server

import (
	"context"
	"fmt"
	"net"

	"github.com/eduardonunesp/kvzika/pkg/db"
	"github.com/eduardonunesp/kvzika/pkg/proto"
	"google.golang.org/grpc"
)

func NewServer(port int) (*Server, error) {
	kvStore, err := db.NewKVStore("/tmp/kvzika")
	if err != nil {
		return nil, err
	}

	if port == 0 {
		port = DefaultPort
	}

	return &Server{
		kvStore: kvStore,
		port:    port,
	}, nil
}

func (s *Server) Close() {
	s.kvStore.Close()
}

func (s *Server) SetKeyValue(ctx context.Context, req *proto.KeyValueRequest) (*proto.KeyValueResponse, error) {
	err := s.kvStore.Set(req.Key, req.Value)
	if err != nil {
		return &proto.KeyValueResponse{
			Status:  proto.Status_ERROR,
			Message: err.Error(),
		}, err
	}

	return &proto.KeyValueResponse{
		Status:  proto.Status_OK,
		Message: "Key set successfully",
	}, nil
}

func (s *Server) GetValue(ctx context.Context, req *proto.KeyRequest) (*proto.KeyValueResponse, error) {
	value, err := s.kvStore.Get(req.Key)
	if err != nil {
		return &proto.KeyValueResponse{
			Status:  proto.Status_ERROR,
			Message: err.Error(),
		}, err
	}

	return &proto.KeyValueResponse{
		Status:  proto.Status_OK,
		Value:   value,
		Message: "Key retrieved successfully",
	}, nil
}

func (s *Server) Run(ctx context.Context) error {
	defer s.Close()

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer listener.Close()

	var opts []grpc.ServerOption = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(maxRecSize),
	}

	grpcServer := grpc.NewServer(opts...)

	proto.RegisterKeyValueServiceServer(grpcServer, s)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}
