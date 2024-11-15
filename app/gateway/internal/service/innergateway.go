package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hoysics/basic-im/app/gateway/internal/biz"

	pb "github.com/hoysics/basic-im/api/gateway/v1"
)

type InnerGatewayService struct {
	pb.UnimplementedInnerGatewayServer
	log *log.Helper
	uc  *biz.ConnectUseCase
}

func NewInnerGatewayService(logger log.Logger, uc *biz.ConnectUseCase) *InnerGatewayService {
	return &InnerGatewayService{
		log: log.NewHelper(logger),
		uc:  uc,
	}
}

func (s *InnerGatewayService) DelConn(ctx context.Context, req *pb.InnerGatewayRequest) (*pb.InnerGatewayResponse, error) {
	s.log.Infof(`DelConn.Recv: %v,`, req)
	s.uc.DeleteConn(req.ConnId)
	return &pb.InnerGatewayResponse{Code: 0}, nil
}
func (s *InnerGatewayService) PushMsg(ctx context.Context, req *pb.InnerGatewayRequest) (*pb.InnerGatewayResponse, error) {
	s.log.Infof(`PushMsg.Recv: %v,`, req)
	return &pb.InnerGatewayResponse{Code: 0}, nil
}
