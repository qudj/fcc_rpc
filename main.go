package main

import (
	"fmt"
	"github.com/qudj/fcc_rpc/config"
	"github.com/qudj/fcc_rpc/models/fcc_serv"
	"google.golang.org/grpc"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	config.InitConfig()

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	fccService := NewFccServiceServer()
	fcc_serv.RegisterFccServiceServer(s, fccService)

	fmt.Println(fmt.Sprintf("listen to %s:", Address))
	s.Serve(listen)
}