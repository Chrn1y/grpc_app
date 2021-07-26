package main

import (
	"context"
	"fmt"

	"example.com/grpcapp/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Went wrong 1")
		return
	}
	c := proto.NewAppClient(conn)
	res, err := c.Add(context.Background(), &proto.AddValues{X: 3, Y: 4})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.GetX())
}
