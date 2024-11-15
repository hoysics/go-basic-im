package service

import (
	"context"
	"github.com/hoysics/basic-im/app/gateway/internal/biz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"

	pb "github.com/hoysics/basic-im/api/gateway/v1"
)

type GatewayService struct {
	pb.UnimplementedGatewayServer

	uc *biz.ConnectUseCase
}

func NewGatewayService(uc *biz.ConnectUseCase) *GatewayService {
	return &GatewayService{
		uc: uc,
	}
}

func (s *GatewayService) Connect(conn pb.Gateway_ConnectServer) error {
	ctx, cancel := context.WithCancel(conn.Context())
	connId, err := s.uc.Connect(ctx, cancel, conn)
	if err != nil {
		return status.Errorf(codes.FailedPrecondition, "ServerStreamingEcho: 分配fd失败")
	}
	defer s.uc.Disconnect(connId)
	for {
		if ctx.Err() != nil {
			return nil
		}
		req, err := conn.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := s.uc.SendMsg(ctx, connId, req.Data); err != nil {
			return err
		}
	}
}
