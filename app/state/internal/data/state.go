package data

//import (
//	"context"
//	v1 "github.com/hoysics/basic-im/api/common/v1"
//	"github.com/hoysics/basic-im/app/state/internal/biz"
//	"google.golang.org/protobuf/proto"
//
//	"github.com/go-kratos/kratos/v2/log"
//)
//
//var _ biz.ConnStateRepo = &connStateRepo{}
//
//type connStateRepo struct {
//	data *Data
//	log  *log.Helper
//}
//
//func (s *connStateRepo) Login(ctx context.Context, connId uint64) error {
//	_stateCache.ConnLogin()
//	return nil
//}
//
//func (s *connStateRepo) Logout(ctx context.Context, connId uint64) error {
//	_, err := _stateCache.ConnLogout(ctx, connId)
//	return err
//}
//
//func (s *connStateRepo) Ack(ctx context.Context, ack *v1.AckMsg) {
//	_stateCache.AckLastMsg(ctx, ack.ConnId, ack.SessionId, ack.MsgId)
//}
//
//func (s *connStateRepo) Push(ctx context.Context, connId uint64, pushMsg *v1.PushMsg) error {
//	by, err := proto.Marshal(pushMsg)
//	if err != nil {
//		return err
//	}
//	//Step 1: Push by gw client
//	_gatewayRepo.Push(ctx, &biz.CmdCtx{ConnId: connId, Msg: &v1.MajorMsg{Cmd: v1.CmdType_PUSH, Payload: by}})
//	//1. State中针对ConnState仅保存一个msgTimer，针对当前连接上最新的Push消息，仅对这个消息进行超时重试
//	//2. 当有新消息Push过来时，取消之前的定时器，staet内仅保留最新push的msg
//	//3. 客户端仅维护对最新消息的Ack，其他消息不用回复
//	//4. 客户端若出现消息漏洞，则直接走兜底的pull模式拉取数据
//	return _stateCache.appendLastMsg(ctx, connId, pushMsg)
//}
//
//func (s *connStateRepo) CompareAndIncrLatestMsg(ctx context.Context, connId uint64, clientId uint64) bool {
//	//TODO Call业务层的RPC // 调用下游业务层rpc，只有当rpc回复成功后才能更新max_clientID
//	//TODO 此处先假设RPC调用成功
//	return _stateCache.CompareAndIncrClientId(ctx, connId, clientId)
//}
//
//func (s *connStateRepo) Delete(ctx context.Context, u uint64) error {
//	return nil
//}
//
//func (s *connStateRepo) ResetHeartbeat(ctx context.Context, u uint64) error {
//	_stateCache.ResetHeartTimer(u)
//	return nil
//}
//
//func (s *connStateRepo) ReConn(ctx context.Context, old, new uint64) (swapped error) {
//	_stateCache.ReConn()
//	return biz.ErrNilConn
//}
//
//// NewConnStateRepo .
//func NewConnStateRepo(data *Data, logger log.Logger) biz.ConnStateRepo {
//	return &connStateRepo{
//		data: data,
//		log:  log.NewHelper(logger),
//	}
//}
