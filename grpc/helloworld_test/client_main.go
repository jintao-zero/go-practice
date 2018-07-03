package main

import (
    "log"
    pb "go-practice/grpc/helloworld/proto"
    "google.golang.org/grpc"
    "context"
)

func main() {
    conn, err := grpc.Dial("127.0.0.1:4000", grpc.WithInsecure())
    if err != nil {
        log.Fatal("eeee ", conn)
    }
    greeterClient := pb.NewGreeterClient(conn)
    reply, err := greeterClient.SayHello(context.Background(), &pb.HelloRequest{Name:"jintao"})
    log.Println(" werwr", reply.String())
}

