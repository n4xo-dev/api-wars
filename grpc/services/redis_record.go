package services

import (
	context "context"

	"github.com/iLopezosa/api-wars/grpc/pb"
	// grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type RedisServiceServer struct {
	pb.UnimplementedRedisServiceServer
}

func (rr *RedisServiceServer) Ping(ctx context.Context, pingReq *pb.PingRequest) (*pb.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (rr *RedisServiceServer) Get(ctx context.Context, getReq *pb.GetRedisRecordRequest) (*pb.GetRedisRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (rr *RedisServiceServer) Set(ctx context.Context, setReq *pb.SetRedisRecordRequest) (*pb.SetRedisRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
