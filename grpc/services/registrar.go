package services

import (
	"github.com/iLopezosa/api-wars/grpc/pb"
	"google.golang.org/grpc"
)

func RegisterServices(s *grpc.Server) {
	pb.RegisterCommentsServiceServer(s, &CommentsServiceServer{})
	pb.RegisterMessagesServiceServer(s, &MessagesServiceServer{})
	pb.RegisterRedisServiceServer(s, &RedisServiceServer{})
}
