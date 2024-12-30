// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.27.0
// source: api/ipconf/v1/ipconf.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationIpConfListIpInfo = "/api.ipconf.v1.IpConf/ListIpInfo"

type IpConfHTTPServer interface {
	ListIpInfo(context.Context, *ListIpInfoRequest) (*ListIpInfoReply, error)
}

func RegisterIpConfHTTPServer(s *http.Server, srv IpConfHTTPServer) {
	r := s.Route("/")
	r.GET("ip-conf/v1/list", _IpConf_ListIpInfo0_HTTP_Handler(srv))
}

func _IpConf_ListIpInfo0_HTTP_Handler(srv IpConfHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListIpInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationIpConfListIpInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListIpInfo(ctx, req.(*ListIpInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListIpInfoReply)
		return ctx.Result(200, reply)
	}
}

type IpConfHTTPClient interface {
	ListIpInfo(ctx context.Context, req *ListIpInfoRequest, opts ...http.CallOption) (rsp *ListIpInfoReply, err error)
}

type IpConfHTTPClientImpl struct {
	cc *http.Client
}

func NewIpConfHTTPClient(client *http.Client) IpConfHTTPClient {
	return &IpConfHTTPClientImpl{client}
}

func (c *IpConfHTTPClientImpl) ListIpInfo(ctx context.Context, in *ListIpInfoRequest, opts ...http.CallOption) (*ListIpInfoReply, error) {
	var out ListIpInfoReply
	pattern := "ip-conf/v1/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationIpConfListIpInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
