package main

import (
	"context"
	"log"
	"time"

	pb "github.com/prasher1421/go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("Client streaming started")

	stream,err := client.SayHelloClientStreaming(context.Background())
	if err!=nil {
		log.Fatalf("Couldn't send Names : %v",err)
	}

	for _,name := range names.Names {
		req := &pb.HelloRequest{
			Name:name,
		}
		if err := stream.Send(req); err!=nil {
			log.Fatalf("Error while sending %v",err)
		}
		log.Printf("Sent the request with name: %s",name)
		time.Sleep(2*time.Second)
	}

	//single response from the server will be recieved here
	res,err := stream.CloseAndRecv()
	
	log.Printf("Client streaming finished")
	
	if err!=nil {
		log.Fatalf("Error while recieving %v",err)
	}

	log.Printf("%v",res.Messages)
}