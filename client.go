package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-demo/helloworld"

	"google.golang.org/grpc"
)

func main(){
	conn,err:=grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err!=nil{
		log.Fatalf("failed to connect : %v",err)
	}
	defer conn.Close()

	client:=pb.NewGreeterClient(conn)
	// name:="subasri"
	//var age int=21
	var age int32
	age =21

	var name string
	fmt.Scanln(&name)
	response,err:=client.SayHello(context.Background(),&pb.HelloRequest{Name: name,Age: age})
	if err!=nil{
		log.Fatalf("failed to call SayHello: %v",err)
	}

	fmt.Printf("Response : %s\n",response.Message)
}