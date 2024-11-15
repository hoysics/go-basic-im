package biz

import (
	"fmt"
	"github.com/hoysics/basic-im/pkg/registry"
)

type EventType string

const (
	AddNodeEvent EventType = "addNode"
	DelNodeEvent EventType = "delNode"
)

type Event struct {
	Type         EventType
	IP           string
	Port         string
	ConnectNum   float64
	MessageBytes float64
}

func NewEvent(endpoint *registry.EndpointInfo) *Event {
	if endpoint == nil || endpoint.MetaData == nil {
		return nil
	}
	var connNum float64
	var messageBytes float64
	if data, ok := endpoint.MetaData["conn_num"]; ok {
		connNum = data.(float64)
	}
	if data, ok := endpoint.MetaData["message_bytes"]; ok {
		messageBytes = data.(float64)
	}
	return &Event{
		Type:         AddNodeEvent,
		IP:           endpoint.Ip,
		Port:         endpoint.Port,
		ConnectNum:   connNum,
		MessageBytes: messageBytes,
	}
}

func (e *Event) Key() string {
	return fmt.Sprintf(`%s:%s`, e.IP, e.Port)
}
