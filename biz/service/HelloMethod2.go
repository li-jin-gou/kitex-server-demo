package service

import (
	"context"

	"github.com/li-jin-gou/kitex-server-demo/kitex_gen/hello/example"
)

type HelloMethod2Service struct {
	ctx context.Context
}

// NewHelloMethod2Service new HelloMethod2Service
func NewHelloMethod2Service(ctx context.Context) *HelloMethod2Service {
	return &HelloMethod2Service{ctx: ctx}
}

// Run create note info
func (s *HelloMethod2Service) Run(request *example.HelloReq) (resp *example.HelloResp, err error) {
	// Finish your business logic.

	return
}
