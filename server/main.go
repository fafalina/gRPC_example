package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/grpc/reflection"
	pb "server/proto"
	"server/db"
)

type timestampService struct {
	pb.UnimplementedTimestampServiceServer
}

func (s *timestampService) GetTimestampData(ctx context.Context, req *emptypb.Empty) (*pb.TimestampData, error) {
	db.WriteCurrentTime()
    timestampData := db.ReadTime()

    grpcData := &pb.TimestampData{
        Timestamp:    timestampData.Timestamp.String(),
    }

    return grpcData, nil
}

func main() {
	db.Init();
	lis, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTimestampServiceServer(s, &timestampService{})

	log.Println("Server is listening on port 8081...")

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
