package main

import (
	"context"

	"github.com/li-jin-gou/kitex-server-demo/biz/service"
	"github.com/li-jin-gou/kitex-server-demo/kitex_gen/hello/example"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// HelloMethod1 implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod1(ctx context.Context, request *example.HelloReq) (resp *example.HelloResp, err error) {
	resp, err = service.NewHelloMethod1Service(ctx).Run(request)

	return resp, err
}

// HelloMethod2 implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod2(ctx context.Context, request *example.HelloReq) (resp *example.HelloResp, err error) {
	resp, err = service.NewHelloMethod2Service(ctx).Run(request)

	return resp, err
}

// HelloMethod3 implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod3(ctx context.Context, request *example.HelloReq) (resp *example.HelloResp, err error) {
	resp, err = service.NewHelloMethod3Service(ctx).Run(request)

	return resp, err
}
