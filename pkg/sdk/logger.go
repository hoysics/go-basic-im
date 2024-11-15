package sdk

import (
	"github.com/go-kratos/kratos/v2/log"
	"os"
)

var _log *log.Helper

func InitLogger() *log.Helper {
	f, err := os.OpenFile("local.sdk.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf(`log init error:%v`, err)
	}
	logger := log.With(log.NewStdLogger(f),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	_log = log.NewHelper(logger)
	return _log
}
