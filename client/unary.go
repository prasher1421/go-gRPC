package main

import (
	"context"
	"log"
	"time"

	pb "github.com/prasher1421/go-grpc/proto"
)

//supposed to expect a response
func callSayHello(client pb.GreetServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(),time.Second)

	defer cancel()

	//sending a request here (no param req)
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err!=nil {
		log.Fatalf("Couldn't greet : %v",err)
	}

	log.Printf("%s",res.Message)
}