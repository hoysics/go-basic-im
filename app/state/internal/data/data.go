package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hoysics/basic-im/app/state/internal/conf"
	"github.com/hoysics/basic-im/pkg/cache"
	"github.com/hoysics/basic-im/pkg/net"
	"github.com/hoysics/basic-im/pkg/router"
	clientV3 "go.etcd.io/etcd/client/v3"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRegistrar, NewConnStateRepo, NewGatewayRepo)

var (
	_stateServerAddr string
	_connSlotList    []int
)

// Data .
type Data struct {
	// TODO wrapped database client
	ctx context.Context
}

// NewData .
func NewData(c *conf.Data, c2 *conf.State, c3 *conf.Server, logger log.Logger) (*Data, func(), error) {
	//1. 加载服务器IP
	var err error
	_stateServerAddr, err = net.GetLocalServerHost(c3.Grpc.GetAddr())
	if err != nil {
		return nil, nil, err
	}
	//2. 初始化时间轮
	InitTimer()
	//3. 初始化全局Cache
	ctx, cancel := context.WithCancel(context.Background())
	_rdb, err = cache.NewRedisClient(c.GetRedis().GetEndpoint(), c.GetRedis().GetPassword())
	if err != nil {
		cancel()
		return nil, nil, err
	}
	if err = InitLuaScript(ctx); err != nil {
		cancel()
		return nil, nil, err
	}
	//4. 初始化路由SDK
	router.Init(ctx, _rdb)
	//5. 初始化缓存槽
	FormatLoginSlotRange(c2.ConnStateSlotRange)
	// 清理函数
	cleanup := func() {
		cancel()
		helper := log.NewHelper(logger)
		helper.Info("closing the data resources")
		CloseTimer()
		helper.Info("finished closing the data resources")
	}
	return &Data{
		ctx: ctx,
	}, cleanup, nil
}

func FormatLoginSlotRange(slotRangeStr string) {
	slotRange := strings.Split(slotRangeStr, ",")
	left, err := strconv.Atoi(slotRange[0])
	if err != nil {
		log.Fatal(err)
	}
	right, err := strconv.Atoi(slotRange[1])
	if err != nil {
		log.Fatal(err)
	}
	res := make([]int, right-left+1)
	for i := left; i < right; i++ {
		res[i-left] = i
	}
	_connSlotList = res
	return
}

func NewRegistrar(reg *conf.Registry) registry.Registrar {
	addr := strings.Split(reg.GetEtcd().GetAddr(), `,`)
	cli, err := clientV3.New(clientV3.Config{
		DialTimeout: reg.GetEtcd().GetTimeout().AsDuration(),
		Endpoints:   addr,
		Username:    reg.GetEtcd().GetUsername(),
		Password:    reg.GetEtcd().GetPassword(),
	})
	if err != nil {
		log.Fatalf("registry etcd connect error: %v", err)
	}
	r := etcd.New(cli)
	return r
}
