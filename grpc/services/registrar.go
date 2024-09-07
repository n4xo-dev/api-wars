package services

import (
	"github.com/n4xo-dev/api-wars/grpc/pb"
	"google.golang.org/grpc"
)

func RegisterServices(s *grpc.Server) {
	pb.RegisterChatsServiceServer(s, &ChatsServiceServer{})
	pb.RegisterCommentsServiceServer(s, &CommentsServiceServer{})
	pb.RegisterMessagesServiceServer(s, &MessagesServiceServer{})
	pb.RegisterPostsServiceServer(s, &PostsServiceServer{})
	pb.RegisterRedisServiceServer(s, &RedisServiceServer{})
	pb.RegisterUserServiceServer(s, &UserServiceServer{})
}
