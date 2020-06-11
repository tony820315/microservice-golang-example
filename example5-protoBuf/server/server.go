package main

import (
	"fmt"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"net"
	pb "example5/protoes/protoBuf/hello"
)

type server struct {

}

// implement rpc SayHi interface
// SayHi(context.Context, *HelloRequest) (*HelloReplay, error)
func (s *server)SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReplay, error) {
	return &pb.HelloReplay{
		Message: "Hi " + in.Name,
	}, nil
}

// GetMsg(context.Context, *HelloRequest) (*HelloMessage, error)

func (s *server)GetMsg(ctx context.Context, in *pb.HelloRequest) (*pb.HelloMessage, error) {
	return &pb.HelloMessage{
		Msg: " message is: " + in.Msg,
	}, nil
}

func main() {
	// monitor network
	ln, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("Network error", err)
		return
	}

	// create grpc server
	svc := grpc.NewServer()

	// Register server
	pb.RegisterHelloServerServer(svc, &server{})

	err = svc.Serve(ln)
	if err != nil {
		fmt.Println("listen err: ", err.Error())
		return
	}
}