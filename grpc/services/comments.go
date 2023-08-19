package services

import (
	context "context"
	// grpc "google.golang.org/grpc"
	"github.com/iLopezosa/api-wars/grpc/pb"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type CommentsServiceServer struct {
	pb.UnimplementedCommentsServiceServer
}

func (c *CommentsServiceServer) ListComments(ctx context.Context, listReq *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComments not implemented")
}
func (c *CommentsServiceServer) GetComment(ctx context.Context, getReq *pb.GetCommentRequest) (*pb.GetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (c *CommentsServiceServer) CreateComment(ctx context.Context, createReq *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (c *CommentsServiceServer) UpdateComment(ctx context.Context, updateReq *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (c *CommentsServiceServer) DeleteComment(ctx context.Context, deleteReq *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
