package main

import (
	"context"
	"fmt"
	"io"

	"example.com/grpcapp/proto"
	"google.golang.org/grpc"
)

const (
	addx  = 3
	addy  = 4
	onesx = 5
)

func main() {

	sumx := []int64{2, 3, 4}

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	c := proto.NewAppClient(conn)
	res, err := c.Add(context.Background(), &proto.AddValues{X: addx, Y: addy})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Add: " + fmt.Sprint(res.GetX()))
	stream, err := c.Sum(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, it := range sumx {
		stream.Send(&proto.Value{X: it})
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sum: " + fmt.Sprint(reply.GetX()))
	stream2, err := c.Ones(context.Background(), &proto.Value{X: onesx})
	if err != nil {
		fmt.Println(err)
		return
	}
	first := true
	fmt.Print("Ones: " + fmt.Sprint(onesx) + " = ")
	for {
		one, err := stream2.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		if !first {
			fmt.Print(" + ")
		}
		first = false
		fmt.Print(fmt.Sprint(one.GetX())) //just for practise
	}
	fmt.Println()
}
