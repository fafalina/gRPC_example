package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
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
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTimestampServiceServer(s, &timestampService{})

	log.Println("Server is listening on port 8081...")

	// go func() {
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// }()
	// defer s.GracefulStop()

	// testGetTimestampData("localhost:8081")
}

/*
 * @brief This function is for testing
 */
func testGetTimestampData(serverAddr string) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewTimestampServiceClient(conn)

	resp, err := client.GetTimestampData(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("GetTimestampData call failed: %v", err)
	}

	log.Printf("Timestamp: %s", resp.Timestamp)
}