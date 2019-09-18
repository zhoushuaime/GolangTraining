package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"time"
	pb "zhoushuai.com/practice/grpc/helloworld"
)

/*
@Time : 2018/11/3 18:17 
@Author : joshua
@File : greet_client
@Software: GoLand
*/

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		 fmt.Println("err:",err)
		return
	}

	fmt.Println("greeting...",r.Message)
}
