package main

import (
	"context"
	"fmt"
	"net"

	hw "grpc-demo/helloworld" //Import the generated Go code

	"google.golang.org/grpc"
)

type server struct {
	hw.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *hw.HelloRequest) (*hw.HelloResponse,error){
	return &hw.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!!, Your Age is %d", req.Name,req.Age),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to Listen: %v", err)
		return
	}

	s := grpc.NewServer()
	hw.RegisterGreeterServer(s, &server{})

	fmt.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
