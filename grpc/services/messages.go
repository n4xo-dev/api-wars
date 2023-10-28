package services

import (
	context "context"
	"errors"
	// grpc "google.golang.org/grpc"
	"github.com/iLopezosa/api-wars/grpc/conv"
	"github.com/iLopezosa/api-wars/grpc/pb"
	"github.com/iLopezosa/api-wars/lib/db"
	"github.com/iLopezosa/api-wars/lib/models"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type MessagesServiceServer struct {
	pb.UnimplementedMessagesServiceServer
}

func (m *MessagesServiceServer) ListMessages(ctx context.Context, listReq *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	messages, err := db.MessageList()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not retrieve messages from database: %v", err)
	}

	pbMessages := make([]*pb.MessageDTO, len(messages))
	for i, message := range messages {
		pbMessages[i] = conv.MessageDTOToPb(message)
	}

	return &pb.ListMessagesResponse{Messages: pbMessages}, nil
}

func (m *MessagesServiceServer) GetMessage(ctx context.Context, getReq *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	if getReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	message, err := db.MessageRead(getReq.Id)
	if err == nil {
		return &pb.GetMessageResponse{Message: conv.MessageDTOToPb(message)}, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "message not found")
	}

	return nil, status.Errorf(codes.Internal, "could not retrieve message from database: %v", err)
}

func (m *MessagesServiceServer) CreateMessage(ctx context.Context, createReq *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {
	message := &models.Message{
		Content: createReq.Content,
		UserID:  createReq.UserId,
		ChatID:  createReq.ChatId,
	}

	if err := db.MessageUpsert(message); err != nil {
		return nil, status.Errorf(codes.Internal, "could not create message: %v", err)
	}

	return &pb.CreateMessageResponse{Message: conv.MessageToPb(*message)}, nil
}

func (m *MessagesServiceServer) UpdateMessage(ctx context.Context, updateReq *pb.UpdateMessageRequest) (*pb.UpdateMessageResponse, error) {
	message := &models.Message{
		ID:      updateReq.Id,
		Content: updateReq.Content,
		UserID:  updateReq.UserId,
		ChatID:  updateReq.ChatId,
	}

	if err := db.MessagePatch(message); err != nil {
		return nil, status.Errorf(codes.Internal, "could not update message: %v", err)
	}

	return &pb.UpdateMessageResponse{Message: conv.MessageToPb(*message)}, nil
}

func (m *MessagesServiceServer) DeleteMessage(ctx context.Context, deleteReq *pb.DeleteMessageRequest) (*pb.DeleteMessageResponse, error) {
	if deleteReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	if err := db.MessageDelete(deleteReq.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete message: %v", err)
	}

	return &pb.DeleteMessageResponse{Deleted: true}, nil
}
