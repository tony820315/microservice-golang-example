package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"micro/protoes"
)

func main(){
	HelloRequest := protoes.HelloRequest{
		Name: *proto.String("zac"),
		UCount: *proto.Int32(98),
	}

	data, err := proto.Marshal(&HelloRequest)
	if err != nil {
		fmt.Println("Marshal err, ", err)
		return
	}
	fmt.Println("Marshal data: ", data)

	var list protoes.HelloRequest
	err = proto.Unmarshal(data, &list)
	if err != nil {
		fmt.Println("Unmarshal err, ", err)
		return
	}
	fmt.Println("Name: ", list.GetName())
	fmt.Println("Name: ", list.GetUCount())
}