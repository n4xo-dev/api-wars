package services

import (
	context "context"

	"github.com/iLopezosa/api-wars/grpc/pb"
	// grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (u *UserServiceServer) ListUsers(ctx context.Context, listReq *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (u *UserServiceServer) GetUser(ctx context.Context, getReq *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (u *UserServiceServer) CreateUser(ctx context.Context, createReq *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (u *UserServiceServer) DeleteUser(ctx context.Context, deleteReq *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (u *UserServiceServer) UpdateUser(ctx context.Context, updateReq *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (u *UserServiceServer) GetUsersPosts(ctx context.Context, getPostsReq *pb.GetUsersPostsRequest) (*pb.ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersPosts not implemented")
}
func (u *UserServiceServer) GetUsersComments(ctx context.Context, getCommentsReq *pb.GetUsersCommentsRequest) (*pb.ListCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersComments not implemented")
}
func (u *UserServiceServer) GetUsersMessages(ctx context.Context, getMessagesReq *pb.GetUsersMessagesRequest) (*pb.ListMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersMessages not implemented")
}
func (u *UserServiceServer) GetUsersMessagesFromChat(ctx context.Context, getChatMessagesReq *pb.GetUsersMessagesFromChatRequest) (*pb.ListMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersMessagesFromChat not implemented")
}
