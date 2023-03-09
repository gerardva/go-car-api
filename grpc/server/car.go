package grpcserver

import (
	"context"
	"log"

	pb "github.com/gerardva/go-api/grpc/protos/carserver"
)

type carServer struct {
	pb.UnimplementedCarManagerServer
}

func (s *carServer) GetCarById(ctx context.Context, in *pb.Int32Value) (*pb.Car, error) {
	log.Printf("Received: %v", in.Value)
	return &pb.Car{Make: "Test"}, nil
}
