package data

import (
	"github.com/hoysics/basic-im/app/ipconf/internal/biz"
)

//TODO 后续再补实际的满足Kratos接口定义的Etcd Discover/Register interface

type sourceRepo struct {
	ec chan *biz.Event
}

func NewSourceRepo() biz.SourceRepo {
	return &sourceRepo{
		ec: make(chan *biz.Event),
	}
}

func (s *sourceRepo) Watch() <-chan *biz.Event {
	return s.ec
}

func (s *sourceRepo) Publish() chan<- *biz.Event {
	//TODO implement me
	panic("implement me")
}
