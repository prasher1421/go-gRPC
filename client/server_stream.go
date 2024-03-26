package main

import (
	"context"
	"io"
	"log"

	pb "github.com/prasher1421/go-grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient,names *pb.NamesList){
	log.Printf("Streaming Started")

	stream, err := client.SayHelloServerStreaming(context.Background(),names)
	if err!=nil {
		log.Fatalf("Couldnt send Names : %v",err)
	}

	for {
		message,err := stream.Recv()
		if err == io.EOF  {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v",err)
		}
		log.Println(message)
	}
	log.Println("Streaming Finished")
}