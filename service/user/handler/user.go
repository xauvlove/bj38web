package handler

import (
	pb "bj38web/service/user/proto"
	"context"
)

type User struct{}

// Return a new handler
func New() *User {
	return &User{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) SendSms(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	return nil
}
