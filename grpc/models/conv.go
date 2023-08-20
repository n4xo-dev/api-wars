package models

import "github.com/iLopezosa/api-wars/grpc/pb"

func (m *Message) ToPbMessageDTO() *pb.MessageDTO {
	return &pb.MessageDTO{
		Id:        m.ID,
		Content:   m.Content,
		UserId:    m.UserID,
		ChatId:    m.ChatID,
		CreatedAt: m.CreatedAt.String(),
		UpdatedAt: m.UpdatedAt.String(),
	}
}

func (m *ReadMessageDTO) ToPbMessageDTO() *pb.MessageDTO {
	return &pb.MessageDTO{
		Id:        m.ID,
		Content:   m.Content,
		UserId:    m.UserID,
		ChatId:    m.ChatID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (c *Comment) ToPbCommentDTO() *pb.CommentDTO {
	return &pb.CommentDTO{
		Id:        c.ID,
		Content:   c.Content,
		UserId:    c.UserID,
		PostId:    c.PostID,
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}

func (c *ReadCommentDTO) ToPbCommentDTO() *pb.CommentDTO {
	return &pb.CommentDTO{
		Id:        c.ID,
		Content:   c.Content,
		UserId:    c.UserID,
		PostId:    c.PostID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func (p *Post) ToPbPostDTO() *pb.PostDTO {
	return &pb.PostDTO{
		Id:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserId:    p.UserID,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

func (p *ReadPostDTO) ToPbPostDTO() *pb.PostDTO {
	return &pb.PostDTO{
		Id:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserId:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (c *Chat) ToPbChat() *pb.Chat {
	pbMessages := make([]*pb.MessageDTO, len(c.Messages))
	for i, m := range c.Messages {
		pbMessages[i] = m.ToPbMessageDTO()
	}

	pbParticipants := make([]*pb.UserDTO, len(c.Participants))
	for i, u := range c.Participants {
		pbParticipants[i] = u.ToPbUserDTO()
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

func (u *User) ToPbUserDTO() *pb.UserDTO {
	return &pb.UserDTO{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}

func (u *ReadUserDTO) ToPbUserDTO() *pb.UserDTO {
	return &pb.UserDTO{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
