package helloworld

import (
    proto "go-practice/grpc/helloworld/proto"
    context "golang.org/x/net/context"
)

type HelloWorldServer struct {
}

func (s *HelloWorldServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
    name := req.GetName()
    return &proto.HelloReply{Message: "hello " + name}, nil
}
