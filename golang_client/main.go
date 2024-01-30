package main

import (
	"context"
	"time"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "golang_client/proto"
)

func main() {
	for {
		testGetTimestampData("grpc_server:8081")
		time.Sleep(1 * time.Second)
	}
}

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