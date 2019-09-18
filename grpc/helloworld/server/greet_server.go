package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	pb "zhoushuai.com/practice/grpc/helloworld"
)

/*
@Time : 2018/11/3 18:18 
@Author : joshua
@File : greet_server
@Software: GoLand
*/

const port = ":50051"

type Server struct {
}

// SayHello ...
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello:" + in.Name}, nil
}
func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	reflection.Register(s)
	fmt.Println("server started...")

	if err := s.Serve(l); err != nil {
		fmt.Println("err:", err)
	}
}
