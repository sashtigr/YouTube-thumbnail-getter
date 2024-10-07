package main

import (
	"context"
	"log"
	"time"

	pb "echelon-test-task/api"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewYouTubeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetThumbnail(ctx, &pb.ThumbnailRequest{
		VideoUrl: "https://www.youtube.com/watch?v=4ZBfmdfCfvQ",
	})
	if err != nil {
		log.Fatalf("could not get thumbnail: %v", err)
	}

	log.Printf("Thumbnail URL: %s", response.GetThumbnailUrl())
}
