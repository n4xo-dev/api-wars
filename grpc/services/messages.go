package services

import (
	context "context"
	// grpc "google.golang.org/grpc"
	"github.com/iLopezosa/api-wars/grpc/pb"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type MessagesServiceServer struct {
	pb.UnimplementedMessagesServiceServer
}

func (m *MessagesServiceServer) ListMessages(ctx context.Context, listReq *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessages not implemented")
}
func (m *MessagesServiceServer) GetMessage(ctx context.Context, getReq *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (m *MessagesServiceServer) CreateMessage(ctx context.Context, createReq *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (m *MessagesServiceServer) UpdateMessage(ctx context.Context, updateReq *pb.UpdateMessageRequest) (*pb.UpdateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (m *MessagesServiceServer) DeleteMessage(ctx context.Context, deleteReq *pb.DeleteMessageRequest) (*pb.DeleteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
