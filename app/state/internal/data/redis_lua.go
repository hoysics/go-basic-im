package data

import (
	"context"
	"errors"
	"fmt"
)

const (
	LuaCompareAndIncrClientId = "LuaCompareAndIncrClientId"
)

type luaPart struct {
	LuaScript string
	Sha       string
}

var luaScriptTable map[string]*luaPart = map[string]*luaPart{
	LuaCompareAndIncrClientId: {
		//检查键KEYS[1]是否存在，如果不存在则创建它并设置值为0。
		//获取键KEYS[1]的值，如果它的值等于ARGV[1]，则增加该键的值并设置过期时间，返回1。
		//如果键KEYS[1]的值不等于ARGV[1]，则返回-1。
		LuaScript: "if redis.call('exists', KEYS[1]) == 0 then redis.call('set', KEYS[1], 0) end;if redis.call('get', KEYS[1]) == ARGV[1] then redis.call('incr', KEYS[1]);redis.call('expire',KEYS[1], ARGV[2]); return 1 else return -1 end",
	},
}

func InitLuaScript(ctx context.Context) error {
	for name, part := range luaScriptTable {
		cmd := _rdb.Client.ScriptLoad(ctx, part.LuaScript)
		if cmd == nil {
			return fmt.Errorf("lua script %s not found", name)
		}
		if cmd.Err() != nil {
			return cmd.Err()
		}
		part.Sha = cmd.Val()
	}
	return nil
}

func RunLuaInt(ctx context.Context, name string, keys []string, args ...interface{}) (int, error) {
	if part, ok := luaScriptTable[name]; !ok {
		return -1, errors.New("lua not registered")
	} else {
		cmd := _rdb.Client.EvalSha(ctx, part.Sha, keys, args)
		if cmd == nil {
			return -1, errors.New("lua EvalSha cmd is nil")
		}
		return cmd.Int()
	}
}
