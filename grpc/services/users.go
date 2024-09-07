package services

import (
	context "context"
	"errors"

	"github.com/n4xo-dev/api-wars/grpc/conv"
	"github.com/n4xo-dev/api-wars/grpc/pb"
	"github.com/n4xo-dev/api-wars/lib/db"
	"github.com/n4xo-dev/api-wars/lib/models"
	"gorm.io/gorm"

	// grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (u *UserServiceServer) ListUsers(ctx context.Context, listReq *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := db.UserList()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error listing users: %v", err)
	}

	pbUsers := make([]*pb.UserDTO, len(users))
	for i, user := range users {
		pbUsers[i] = conv.UserDTOToPb(user)
	}

	return &pb.ListUsersResponse{Users: pbUsers}, nil
}

func (u *UserServiceServer) GetUser(ctx context.Context, getReq *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if getReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	user, err := db.UserRead(getReq.Id)
	if err == nil {
		return &pb.GetUserResponse{User: conv.UserDTOToPb(user)}, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return nil, status.Errorf(codes.Internal, "error getting user: %v", err)
}

func (u *UserServiceServer) CreateUser(ctx context.Context, createReq *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &models.User{
		Name:  createReq.Name,
		Email: createReq.Email,
	}
	if err := db.UserUpsert(user); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating user: %v", err)
	}

	return &pb.CreateUserResponse{User: conv.UserToPb(*user)}, nil
}

func (u *UserServiceServer) DeleteUser(ctx context.Context, deleteReq *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if deleteReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	if err := db.UserDelete(deleteReq.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting user: %v", err)
	}

	return &pb.DeleteUserResponse{Deleted: true}, nil
}

func (u *UserServiceServer) UpdateUser(ctx context.Context, updateReq *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if updateReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	user := &models.User{
		ID:    updateReq.Id,
		Name:  updateReq.Name,
		Email: updateReq.Email,
	}
	if err := db.UserPatch(user); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating user: %v", err)
	}

	return &pb.UpdateUserResponse{User: conv.UserToPb(*user)}, nil
}

func (u *UserServiceServer) GetUsersPosts(ctx context.Context, getPostsReq *pb.GetUsersPostsRequest) (*pb.ListPostsResponse, error) {
	if getPostsReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	posts, err := db.PostListByUserID(getPostsReq.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user posts: %v", err)
	}

	pbPosts := make([]*pb.PostDTO, len(posts))
	for i, post := range posts {
		pbPosts[i] = conv.PostDTOToPb(post)
	}

	return &pb.ListPostsResponse{Posts: pbPosts}, nil
}

func (u *UserServiceServer) GetUsersComments(ctx context.Context, getCommentsReq *pb.GetUsersCommentsRequest) (*pb.ListCommentsResponse, error) {
	if getCommentsReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	comments, err := db.CommentListByUserID(getCommentsReq.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user comments: %v", err)
	}

	pbComments := make([]*pb.CommentDTO, len(comments))
	for i, comment := range comments {
		pbComments[i] = conv.CommentDTOToPb(comment)
	}

	return &pb.ListCommentsResponse{Comments: pbComments}, nil
}

func (u *UserServiceServer) GetUsersMessages(ctx context.Context, getMessagesReq *pb.GetUsersMessagesRequest) (*pb.ListMessagesResponse, error) {
	if getMessagesReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	messages, err := db.MessageListByUserID(getMessagesReq.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user messages: %v", err)
	}

	pbMessages := make([]*pb.MessageDTO, len(messages))
	for i, message := range messages {
		pbMessages[i] = conv.MessageDTOToPb(message)
	}

	return &pb.ListMessagesResponse{Messages: pbMessages}, nil
}

func (u *UserServiceServer) GetUsersMessagesFromChat(ctx context.Context, getChatMessagesReq *pb.GetUsersMessagesFromChatRequest) (*pb.ListMessagesResponse, error) {
	if getChatMessagesReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id")
	}
	if getChatMessagesReq.ChatId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid chat id")
	}

	messages, err := db.MessageListByChatIDAndUserID(getChatMessagesReq.ChatId, getChatMessagesReq.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user messages: %v", err)
	}

	pbMessages := make([]*pb.MessageDTO, len(messages))
	for i, message := range messages {
		pbMessages[i] = conv.MessageDTOToPb(message)
	}

	return &pb.ListMessagesResponse{Messages: pbMessages}, nil
}
