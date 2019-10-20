package server

import (
	"context"
	"fmt"
	pb "github.com/offer365/example/grpc_example/core/proto"
	"testing"
)

type HelloServiceImpl struct{}

// 实现HelloServiceServer接口
func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *pb.String,
) (*pb.String, error) {
	reply := &pb.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func TestNewRpcServer(t *testing.T) {
	//grpcserver,listener,err:=NewRpcServer()
	grpcserver,listener,err:=NewRpcServer(
		WithAddr(":7890"),
		)
	fmt.Println(err)
	helloServer:=new(HelloServiceImpl)
	pb.RegisterHelloServiceServer(grpcserver,helloServer)
	grpcserver.Serve(listener)
}