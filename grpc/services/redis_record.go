package services

import (
	context "context"

	"github.com/iLopezosa/api-wars/grpc/db"
	"github.com/iLopezosa/api-wars/grpc/pb"
	"github.com/redis/go-redis/v9"

	// grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type RedisServiceServer struct {
	pb.UnimplementedRedisServiceServer
}

func (rr *RedisServiceServer) Ping(ctx context.Context, pingReq *pb.PingRequest) (*pb.PingResponse, error) {
	val, err := db.RedisPing()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error pinging redis: %v", err)
	}
	return &pb.PingResponse{Message: val}, nil
}

func (rr *RedisServiceServer) Get(ctx context.Context, getReq *pb.GetRedisRecordRequest) (*pb.GetRedisRecordResponse, error) {
	val, err := db.RedisGet(getReq.Key)
	if err == redis.Nil {
		return nil, status.Errorf(codes.NotFound, "key %s not found", getReq.Key)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting key %s: %v", getReq.Key, err)
	}

	return &pb.GetRedisRecordResponse{
		Record: &pb.RedisRecord{
			Key:   getReq.Key,
			Value: val,
		},
	}, nil
}

func (rr *RedisServiceServer) Set(ctx context.Context, setReq *pb.SetRedisRecordRequest) (*pb.SetRedisRecordResponse, error) {
	err := db.RedisSet(setReq.Record.Key, setReq.Record.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error setting key %s: %v", setReq.Record.Key, err)
	}
	return &pb.SetRedisRecordResponse{Message: "OK"}, nil
}
