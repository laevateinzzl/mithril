package handlers

import (
	"context"

	pb "mithril/videosvc"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.VideosServer {
	return videosService{}
}

type videosService struct{}

func (s videosService) PostVideo(ctx context.Context, in *pb.PostRequest) (*pb.PostResponse, error) {
	var resp pb.PostResponse
	return &resp, nil
}

func (s videosService) ListVideo(ctx context.Context, in *pb.ListRequest) (*pb.ListResponse, error) {
	var resp pb.ListResponse
	return &resp, nil
}

func (s videosService) GetVideo(ctx context.Context, in *pb.GetVideoRequest) (*pb.GetResponse, error) {
	var resp pb.GetResponse
	return &resp, nil
}
