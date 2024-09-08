package services

import (
	context "context"
	"errors"

	"github.com/n4xo-dev/api-wars/grpc/conv"
	"github.com/n4xo-dev/api-wars/grpc/pb"
	"github.com/n4xo-dev/api-wars/lib/db"
	"github.com/n4xo-dev/api-wars/lib/models"
	"gorm.io/gorm"

	//grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ChatsServiceServer struct {
	pb.UnimplementedChatsServiceServer
}

func (c *ChatsServiceServer) ListChats(ctx context.Context, listReq *pb.ListChatsRequest) (*pb.ListChatsResponse, error) {
	chats, err := db.ChatList(listReq.Eager)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error listing chats: %v", err)
	}

	pbChats := make([]*pb.Chat, len(chats))
	for i, chat := range chats {
		pbChats[i] = conv.ChatToPb(*chat)
	}

	return &pb.ListChatsResponse{
		Chats: pbChats,
	}, nil
}

func (c *ChatsServiceServer) GetChat(ctx context.Context, getReq *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	if getReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	chat, err := db.ChatRead(getReq.Id, getReq.Eager)
	if err == nil {
		return &pb.GetChatResponse{Chat: conv.ChatToPb(chat)}, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "chat not found")
	}

	return nil, status.Errorf(codes.Internal, "error getting chat: %v", err)
}

func (c *ChatsServiceServer) CreateChat(ctx context.Context, createReq *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	chat := new(models.Chat)
	if err := db.ChatUpsert(chat); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating chat: %v", err)
	}

	return &pb.CreateChatResponse{
		Chat: conv.ChatToPb(*chat),
	}, nil
}

func (c *ChatsServiceServer) AddUsersToChat(ctx context.Context, addUsersReq *pb.AddUsersToChatRequest) (*pb.AddUsersToChatResponse, error) {
	if addUsersReq.ChatId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid chat id")
	}

	chat, err := db.ChatRead(addUsersReq.ChatId, true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting chat: %v", err)
	}

	users := make([]*models.User, len(addUsersReq.UserIds))
	for i, id := range addUsersReq.UserIds {
		users[i] = &models.User{
			ID: id,
		}
	}

	chat.Participants = append(chat.Participants, users...)

	if err := db.ChatPatch(&chat); err != nil {
		return nil, status.Errorf(codes.Internal, "error adding users to chat: %v", err)
	}

	return &pb.AddUsersToChatResponse{
		Chat: conv.ChatToPb(chat),
	}, nil
}

func (c *ChatsServiceServer) DeleteChat(ctx context.Context, deleteReq *pb.DeleteChatRequest) (*pb.DeleteChatResponse, error) {
	if deleteReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	if err := db.ChatDelete(deleteReq.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting chat: %v", err)
	}

	return &pb.DeleteChatResponse{Deleted: true}, nil
}

func (c *ChatsServiceServer) GetChatMessages(ctx context.Context, getMessagesReq *pb.GetChatMessagesRequest) (*pb.GetChatMessagesResponse, error) {
	if getMessagesReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid chat id")
	}

	messages, err := db.MessageListByChatID(getMessagesReq.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting messages: %v", err)
	}

	pbMessages := make([]*pb.MessageDTO, len(messages))
	for i, message := range messages {
		pbMessages[i] = conv.MessageDTOToPb(message)
	}

	return &pb.GetChatMessagesResponse{
		Messages: pbMessages,
	}, nil
}

func (c *ChatsServiceServer) GetChatUserMessages(ctx context.Context, getUserMessagesReq *pb.GetChatUserMessagesRequest) (*pb.GetChatUserMessagesResponse, error) {
	if getUserMessagesReq.ChatId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid chat id")
	}
	if getUserMessagesReq.UserId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id")
	}

	messages, err := db.MessageListByChatIDAndUserID(getUserMessagesReq.ChatId, getUserMessagesReq.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting messages: %v", err)
	}

	pbMessages := make([]*pb.MessageDTO, len(messages))
	for i, message := range messages {
		pbMessages[i] = conv.MessageDTOToPb(message)
	}

	return &pb.GetChatUserMessagesResponse{
		Messages: pbMessages,
	}, nil
}
