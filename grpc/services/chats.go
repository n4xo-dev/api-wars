package services

import (
	context "context"

	"github.com/iLopezosa/api-wars/grpc/pb"
	//grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ChatsServer struct {
	pb.UnimplementedChatsServer
}

func (c *ChatsServer) ListChats(ctx context.Context, listReq *pb.ListChatsRequest) (*pb.ListChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChats not implemented")
}
func (c *ChatsServer) GetChat(ctx context.Context, getReq *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (c *ChatsServer) CreateChat(ctx context.Context, createReq *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (c *ChatsServer) AddUsersToChat(ctx context.Context, addUsersReq *pb.AddUsersToChatRequest) (*pb.AddUsersToChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUsersToChat not implemented")
}
func (c *ChatsServer) DeleteChat(ctx context.Context, deleteReq *pb.DeleteChatRequest) (*pb.DeleteChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (c *ChatsServer) GetChatMessages(ctx context.Context, getMessagesReq *pb.GetChatMessagesRequest) (*pb.GetChatMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMessages not implemented")
}
func (c *ChatsServer) GetChatUserMessages(ctx context.Context, getUserMessagesReq *pb.GetChatUserMessagesRequest) (*pb.GetChatUserMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatUserMessages not implemented")
}
