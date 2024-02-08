package server

import (
	"github.com/eduardonunesp/kvzika/pkg/db"
	"github.com/eduardonunesp/kvzika/pkg/proto"
)

const (
	maxRecSize  = 1024 * 1024 * 5 // 5MB
	DefaultPort = 5566
)

type Server struct {
	proto.UnimplementedKeyValueServiceServer

	kvStore *db.KVStore
	port    int
}
