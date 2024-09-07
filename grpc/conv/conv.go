package conv

import (
	"github.com/n4xo-dev/api-wars/grpc/pb"
	"github.com/n4xo-dev/api-wars/lib/models"
)

func MessageToPb(m models.Message) *pb.MessageDTO {
	return &pb.MessageDTO{
		Id:        m.ID,
		Content:   m.Content,
		UserId:    m.UserID,
		ChatId:    m.ChatID,
		CreatedAt: m.CreatedAt.String(),
		UpdatedAt: m.UpdatedAt.String(),
	}
}

func MessageDTOToPb(m models.ReadMessageDTO) *pb.MessageDTO {
	return &pb.MessageDTO{
		Id:        m.ID,
		Content:   m.Content,
		UserId:    m.UserID,
		ChatId:    m.ChatID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func CommentToPb(c models.Comment) *pb.CommentDTO {
	return &pb.CommentDTO{
		Id:        c.ID,
		Content:   c.Content,
		UserId:    c.UserID,
		PostId:    c.PostID,
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}

func CommentDTOToPb(c models.ReadCommentDTO) *pb.CommentDTO {
	return &pb.CommentDTO{
		Id:        c.ID,
		Content:   c.Content,
		UserId:    c.UserID,
		PostId:    c.PostID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func PostToPb(p models.Post) *pb.PostDTO {
	return &pb.PostDTO{
		Id:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserId:    p.UserID,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

func PostDTOToPb(p models.ReadPostDTO) *pb.PostDTO {
	return &pb.PostDTO{
		Id:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserId:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ChatToPb(c models.Chat) *pb.Chat {
	pbMessages := make([]*pb.MessageDTO, len(c.Messages))
	for i, m := range c.Messages {
		pbMessages[i] = MessageToPb(m)
	}

	pbParticipants := make([]*pb.UserDTO, len(c.Participants))
	for i, u := range c.Participants {
		pbParticipants[i] = UserToPb(*u)
	}

	return &pb.Chat{
		Id:           c.ID,
		Messages:     pbMessages,
		Participants: pbParticipants,
		CreatedAt:    c.CreatedAt.String(),
		UpdatedAt:    c.UpdatedAt.String(),
		DeletedAt:    c.DeletedAt.Time.String(),
	}
}

func UserToPb(u models.User) *pb.UserDTO {
	return &pb.UserDTO{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}

func UserDTOToPb(u models.ReadUserDTO) *pb.UserDTO {
	return &pb.UserDTO{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
