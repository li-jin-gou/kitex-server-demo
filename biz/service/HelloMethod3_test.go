package service

import (
	"context"
	"testing"

	"github.com/li-jin-gou/kitex-server-demo/kitex_gen/hello/example"
)

func TestHelloMethod3_Run(t *testing.T) {
	ctx := context.Background()
	s := NewHelloMethod3Service(ctx)
	// init req and assert value

	request := &example.HelloReq{}

	resp, err := s.Run(request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	// todo: edit your unit test
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
}
