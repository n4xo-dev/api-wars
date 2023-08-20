package services

import (
	context "context"

	"github.com/iLopezosa/api-wars/grpc/db"
	"github.com/iLopezosa/api-wars/grpc/models"
	"github.com/iLopezosa/api-wars/grpc/pb"

	// grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type PostsServiceServer struct {
	pb.UnimplementedPostsServiceServer
}

func (p *PostsServiceServer) ListPosts(ctx context.Context, listReq *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := db.PostList()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while getting posts from database: %v", err)
	}

	pbPosts := make([]*pb.PostDTO, len(posts))
	for i, post := range posts {
		pbPosts[i] = post.ToPbPostDTO()
	}

	return &pb.ListPostsResponse{Posts: pbPosts}, nil
}

func (p *PostsServiceServer) GetPost(ctx context.Context, getReq *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	if getReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id")
	}

	post, err := db.PostRead(getReq.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while getting post from database: %v", err)
	}

	return &pb.GetPostResponse{Post: post.ToPbPostDTO()}, nil
}

func (p *PostsServiceServer) CreatePost(ctx context.Context, createReq *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	post := &models.Post{
		Title:   createReq.Title,
		Content: createReq.Content,
		UserID:  createReq.UserId,
	}

	if err := db.PostUpsert(post); err != nil {
		return nil, status.Errorf(codes.Internal, "Error while creating post: %v", err)
	}

	return &pb.CreatePostResponse{Post: post.ToPbPostDTO()}, nil
}

func (p *PostsServiceServer) UpdatePost(ctx context.Context, updateReq *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	post := &models.Post{
		ID:      updateReq.Id,
		Title:   updateReq.Title,
		Content: updateReq.Content,
		UserID:  updateReq.UserId,
	}

	if err := db.PostPatch(post); err != nil {
		return nil, status.Errorf(codes.Internal, "Error while updating post: %v", err)
	}

	return &pb.UpdatePostResponse{Post: post.ToPbPostDTO()}, nil
}

func (p *PostsServiceServer) DeletePost(ctx context.Context, deleteReq *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	if deleteReq.Id < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id")
	}

	if err := db.PostDelete(deleteReq.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "Error while deleting post: %v", err)
	}

	return &pb.DeletePostResponse{Deleted: true}, nil
}

func (p *PostsServiceServer) GetPostsComments(ctx context.Context, getCommentsReq *pb.GetPostsCommentsRequest) (*pb.GetPostsCommentsResponse, error) {
	if getCommentsReq.PostId < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id")
	}

	comments, err := db.CommentListByPostID(getCommentsReq.PostId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while getting comments from database: %v", err)
	}

	pbComments := make([]*pb.CommentDTO, len(comments))
	for i, comment := range comments {
		pbComments[i] = comment.ToPbCommentDTO()
	}

	return &pb.GetPostsCommentsResponse{Comments: pbComments}, nil
}
