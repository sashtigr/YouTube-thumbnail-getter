package main

import (
	"context"
	pb "echelon-test-task/api" // Обновите путь к вашему сгенерированному коду
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strings"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedYouTubeServiceServer
}

func (s *server) GetThumbnail(ctx context.Context, req *pb.ThumbnailRequest) (*pb.ThumbnailResponse, error) {
	videoURL := req.GetVideoUrl()
	videoID := extractVideoID(videoURL)

	thumbnailURL := fmt.Sprintf("https://img.youtube.com/vi/%s/maxresdefault.jpg", videoID)

	if !checkThumbnailExists(thumbnailURL) {
		return nil, fmt.Errorf("thumbnail not found for video ID: %s", videoID)
	}

	return &pb.ThumbnailResponse{ThumbnailUrl: thumbnailURL}, nil
}

func extractVideoID(videoURL string) string {
	videoArr := strings.Split(videoURL, "v=")
	return videoArr[1]
}

func checkThumbnailExists(thumbnailURL string) bool {
	resp, err := http.Head(thumbnailURL)
	return err == nil && resp.StatusCode == http.StatusOK
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterYouTubeServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
