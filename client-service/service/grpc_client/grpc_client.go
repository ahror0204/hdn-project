package grpcUser

import (
	"github.com/hdn-project/User-service/config"
)

//GrpcUserI ...
type GrpcUserI interface {
}

//GrpcUser ...
type GrpcUser struct {
	cfg         config.Config
	connections map[string]interface{}
}

//New ...
func New(cfg config.Config) (*GrpcUser, error) {
	return &GrpcUser{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}
