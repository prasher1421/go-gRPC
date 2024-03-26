package main

import (
	"log"
	"time"

	pb "github.com/prasher1421/go-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names: %v",req.Names)

	for _,name := range req.Names{
		res := &pb.HelloResponse{
			Message: "Heyyy "+name,
		}
		//here a stream of messages will be sent
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2*time.Second)
	}

	return nil

}