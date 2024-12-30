package server

import (
	v1 "github.com/hoysics/basic-im/api/gateway/v1"
	"github.com/hoysics/basic-im/app/gateway/internal/conf"
	"github.com/hoysics/basic-im/app/gateway/internal/service"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, gw *service.GatewayService, inner *service.InnerGatewayService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGatewayServer(srv, gw)
	v1.RegisterInnerGatewayServer(srv, inner)
	return srv
}
