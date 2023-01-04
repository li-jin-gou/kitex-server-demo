package service

import (
	"context"

	"github.com/li-jin-gou/kitex-server-demo/kitex_gen/hello/example"
)

type HelloMethod1Service struct {
	ctx context.Context
}

// NewHelloMethod1Service new HelloMethod1Service
func NewHelloMethod1Service(ctx context.Context) *HelloMethod1Service {
	return &HelloMethod1Service{ctx: ctx}
}

// Run create note info
func (s *HelloMethod1Service) Run(request *example.HelloReq) (resp *example.HelloResp, err error) {
	// Finish your business logic.

	return
}
