package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "example5/protoes/protoBuf/hello"
)

func main() {
	// connect the grpc server
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithInsecure())
	if err != nil {
		fmt.Println("connect fail: ", err)
	}
	defer conn.Close()
	c := pb.NewHelloServerClient(conn)

	// rpc rfc SayHi
	r1, err := c.SayHi(
		context.Background(),
		&pb.HelloRequest{
			Name: "Kitty",
			Msg: "you are so beautiful! ",
		},
	)
	if err != nil {
		fmt.Println("Can not get SayHil: ", err)
		return
	}
	fmt.Println("SayHi response: ", r1)

	// rpc rfc SayHi
	r2, err := c.GetMsg(
		context.Background(),
		&pb.HelloRequest{
			Name: "Zach",
			Msg: "where are you from?",
		},
	)
	if err != nil {
		fmt.Println("Can not get GetMsg: ", err)
		return
	}
	fmt.Println("GetMsg response: ", r2)
}
