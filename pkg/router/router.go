package router

import (
	"errors"
	"fmt"
	"github.com/hoysics/basic-im/pkg/cache"
	"golang.org/x/net/context"
	"strconv"
	"strings"
	"time"
)

const (
	gatewayRouterKey = "gateway_router_%d"
	ttl7D            = 7 * 24 * 60 * 60
)

var _rdb *cache.RedisClient

type Record struct {
	Endpoint string
	ConnId   uint64
}

func Init(ctx context.Context, client *cache.RedisClient) {
	_rdb = client
}

func AddRecord(ctx context.Context, did uint64, endpoint string, connId uint64) error {
	key := fmt.Sprintf(gatewayRouterKey, did)
	value := fmt.Sprintf("%s-%d", endpoint, connId)
	return _rdb.SetString(ctx, key, value, ttl7D*time.Second)
}

func DelRecord(ctx context.Context, did uint64) error {
	key := fmt.Sprintf(gatewayRouterKey, did)
	return _rdb.Del(ctx, key)
}

func QueryRecord(ctx context.Context, did uint64) (Record, error) {
	key := fmt.Sprintf(gatewayRouterKey, did)
	data, err := _rdb.GetString(ctx, key)
	if err != nil {
		return Record{}, err
	}
	ec := strings.Split(data, "-")
	if len(ec) != 2 {
		return Record{}, errors.New(`record format error`)
	}
	connId, err := strconv.ParseUint(ec[1], 10, 64)
	if err != nil {
		return Record{}, err
	}
	return Record{
		Endpoint: ec[0],
		ConnId:   connId,
	}, nil
}
