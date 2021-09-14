package main

import (
	"context"
	"fmt"

	pb "resource/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:10216", grpc.WithInsecure())
	if nil != err {
		fmt.Printf("error %v\n", err)
		panic("grpc dial error")
	}
	defer conn.Close()

	cli := pb.NewACDServiceClient(conn)
	req := pb.RequestAgentState{
		AppId:   "hahaha",
		AgentId: "jahaja",
	}
	resp, err := cli.AgentState(context.Background(), &req)
	if nil != err {
		panic("grcp call error")
	}

	fmt.Printf("resp state is %d\n", resp.State)
}
