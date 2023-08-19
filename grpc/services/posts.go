package services

import (
	context "context"

	"github.com/iLopezosa/api-wars/grpc/pb"
	// grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type PostsServiceServer struct {
	pb.UnimplementedPostsServiceServer
}

func (p *PostsServiceServer) ListPosts(ctx context.Context, listReq *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (p *PostsServiceServer) GetPost(ctx context.Context, getReq *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (p *PostsServiceServer) CreatePost(ctx context.Context, createReq *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (p *PostsServiceServer) UpdatePost(ctx context.Context, updateReq *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (p *PostsServiceServer) DeletePost(ctx context.Context, deleteReq *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (p *PostsServiceServer) GetPostsComments(ctx context.Context, getCommentsReq *pb.GetPostsCommentsRequest) (*pb.GetPostsCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsComments not implemented")
}
