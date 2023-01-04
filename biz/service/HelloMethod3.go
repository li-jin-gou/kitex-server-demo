package service

import (
	"context"

	"github.com/li-jin-gou/kitex-server-demo/kitex_gen/hello/example"
)

type HelloMethod3Service struct {
	ctx context.Context
}

// NewHelloMethod3Service new HelloMethod3Service
func NewHelloMethod3Service(ctx context.Context) *HelloMethod3Service {
	return &HelloMethod3Service{ctx: ctx}
}

// Run create note info
func (s *HelloMethod3Service) Run(request *example.HelloReq) (resp *example.HelloResp, err error) {
	// Finish your business logic.

	return
}
