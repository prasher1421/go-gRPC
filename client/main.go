package main

import (
	"log"

	pb "github.com/prasher1421/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	//to make req
	conn, err := grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil {
		log.Fatalf("Didn't connect client %v",err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Aryan", "Prasher","HAHA"},
	}

	// callSayHello(client)
	// log.Println("About a send a request")
	// callSayHelloServerStream(client,names)
	
	log.Println("About a send request streams")
	// callSayHelloClientStream(client,names)

	callHelloBidirectionalStream(client,names)
}