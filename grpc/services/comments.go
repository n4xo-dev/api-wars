package services

import (
	context "context"

	// grpc "google.golang.org/grpc"
	"github.com/iLopezosa/api-wars/grpc/db"
	"github.com/iLopezosa/api-wars/grpc/models"
	"github.com/iLopezosa/api-wars/grpc/pb"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type CommentsServiceServer struct {
	pb.UnimplementedCommentsServiceServer
}

func (c *CommentsServiceServer) ListComments(ctx context.Context, listReq *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	comments, err := db.CommentList()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list comments: %v", err)
	}

	commentsPB := make([]*pb.CommentDTO, len(comments))
	const dateLayout = "2006-01-02T15:04:05.999999-07:00"

	for i, comment := range comments {
		commentsPB[i] = &pb.CommentDTO{
			Id:        comment.ID,
			Content:   comment.Content,
			UserId:    comment.UserID,
			PostId:    comment.PostID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}
	}

	return &pb.ListCommentsResponse{Comments: commentsPB}, nil
}

func (c *CommentsServiceServer) GetComment(ctx context.Context, getReq *pb.GetCommentRequest) (*pb.GetCommentResponse, error) {
	if getReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	comment, err := db.CommentRead(getReq.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get comment: %v", err)
	}

	const dateLayout = "2006-01-02T15:04:05.999999-07:00"

	commentPB := &pb.CommentDTO{
		Id:        comment.ID,
		Content:   comment.Content,
		UserId:    comment.UserID,
		PostId:    comment.PostID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}

	return &pb.GetCommentResponse{Comment: commentPB}, nil
}

func (c *CommentsServiceServer) CreateComment(ctx context.Context, createReq *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	comment := &models.Comment{
		Content: createReq.Content,
		UserID:  createReq.UserId,
		PostID:  createReq.PostId,
	}
	if err := db.CommentUpsert(comment); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create comment: %v", err)
	}

	return &pb.CreateCommentResponse{Comment: &pb.CommentDTO{
		Id:        comment.ID,
		Content:   comment.Content,
		UserId:    comment.UserID,
		PostId:    comment.PostID,
		CreatedAt: comment.CreatedAt.String(),
		UpdatedAt: comment.UpdatedAt.String(),
	}}, nil
}

func (c *CommentsServiceServer) UpdateComment(ctx context.Context, updateReq *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	if updateReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	comment := &models.Comment{
		ID:      updateReq.Id,
		Content: updateReq.Content,
		UserID:  updateReq.UserId,
		PostID:  updateReq.PostId,
	}
	if err := db.CommentPatch(comment); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update comment: %v", err)
	}

	return &pb.UpdateCommentResponse{Comment: &pb.CommentDTO{
		Id:        comment.ID,
		Content:   comment.Content,
		UserId:    comment.UserID,
		PostId:    comment.PostID,
		CreatedAt: comment.CreatedAt.String(),
		UpdatedAt: comment.UpdatedAt.String(),
	}}, nil
}

func (c *CommentsServiceServer) DeleteComment(ctx context.Context, deleteReq *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	if deleteReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	if err := db.CommentDelete(deleteReq.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete comment: %v", err)
	}

	return &pb.DeleteCommentResponse{Deleted: true}, nil
}
