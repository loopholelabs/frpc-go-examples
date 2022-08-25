package main

import (
	"context"
	"net"
	"os"

	"github.com/loopholelabs/frpc-go-examples/grpc/echo"
	"google.golang.org/grpc"
)

type svc struct {
	echo.UnimplementedEchoServiceServer
}

func (s *svc) Echo(_ context.Context, req *echo.Request) (*echo.Response, error) {
	res := new(echo.Response)
	res.Message = req.Message
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	echo.RegisterEchoServiceServer(server, new(svc))

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
