package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/hoysics/basic-im/api/state/v1"
)

var (
	ErrWhenCallGwDelConn = errors.New(1, `call gateway client failed`, ``)
	ErrGwDelConn         = errors.New(1, `this gateway client delete failed`, ``)
)

type GatewayRepo interface {
	PushPipe() chan *pb.ChanMsg
}

type GatewayClient struct {
	log  *log.Helper
	repo GatewayRepo
}

func NewGatewayClient(logger log.Logger, repo GatewayRepo) *GatewayClient {
	return &GatewayClient{
		log:  log.NewHelper(logger),
		repo: repo,
	}
}

func (c *GatewayClient) PushPipe() chan *pb.ChanMsg {
	return c.repo.PushPipe()
}
