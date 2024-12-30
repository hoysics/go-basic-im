package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hoysics/basic-im/app/ipconf/internal/biz"

	pb "github.com/hoysics/basic-im/api/ipconf/v1"
)

type IpConfService struct {
	pb.UnimplementedIpConfServer

	log *log.Helper

	dispatcher *biz.Dispatcher
}

func NewIpConfService(logger log.Logger, dispatcher *biz.Dispatcher) *IpConfService {
	return &IpConfService{
		log:        log.NewHelper(logger),
		dispatcher: dispatcher,
	}
}

func (s *IpConfService) ListIpInfo(ctx context.Context, req *pb.ListIpInfoRequest) (*pb.ListIpInfoReply, error) {
	return &pb.ListIpInfoReply{}, nil
}
