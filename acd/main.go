package main

import (
	"acd/service"
	"net"
	"os"

	pb "acd/proto"

	"google.golang.org/grpc"
)

func main() {
	endpoint := os.Getenv("ACD_LISTEN")
	if len(endpoint) == 0 {
		endpoint = "0.0.0.0:10216"
	}

	lis, err := net.Listen("tcp", endpoint)
	if nil != err {
		panic("error tcp listen")
	}

	svc := service.AcdService{}
	srv := grpc.NewServer()
	pb.RegisterACDServiceServer(srv, svc)
	srv.Serve(lis)
}
