package main

import (
    "go-practice/grpc/helloworld"
    pb "go-practice/grpc/helloworld/proto"
    "google.golang.org/grpc"
    "net"
    "log"
)

func main()  {
    s := helloworld.HelloWorldServer{}
    rpcServer := grpc.NewServer()
    pb.RegisterGreeterServer(rpcServer, &s)
    l, err := net.Listen("tcp", ":4000")
    if err != nil {
        log.Fatal("listen err ", err)
    }
    log.Println(rpcServer.Serve(l))
}