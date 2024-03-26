package main

import (
	"context"
	
	
	pb "github.com/prasher1421/go-grpc/proto"
)

//returns a response message
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse,error){
	return &pb.HelloResponse{
		Message: "This is a Unary Message",
	},nil
}