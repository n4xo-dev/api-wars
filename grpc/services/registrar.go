package services

import (
	"github.com/iLopezosa/api-wars/grpc/pb"
	"google.golang.org/grpc"
)

func RegisterServices(s *grpc.Server) {
	pb.RegisterRedisServiceServer(s, &RedisServiceServer{})
}
