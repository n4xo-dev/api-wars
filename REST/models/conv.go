package models

func (m *WriteMessageDTO) ToMessage() Message {
	return Message{
		Content: m.Content,
		UserID:  m.UserID,
		ChatID:  m.ChatID,
	}
}

func (m *Message) ToReadMessageDTO() ReadMessageDTO {
	return ReadMessageDTO{
		ID:        m.ID,
		Content:   m.Content,
		UserID:    m.UserID,
		ChatID:    m.ChatID,
		CreatedAt: m.CreatedAt.String(),
		UpdatedAt: m.UpdatedAt.String(),
	}
}

func (c *WriteCommentDTO) ToComment() Comment {
	return Comment{
		Content: c.Content,
		UserID:  c.UserID,
		PostID:  c.PostID,
	}
}

func (c *Comment) ToReadCommentDTO() ReadCommentDTO {
	return ReadCommentDTO{
		ID:        c.ID,
		Content:   c.Content,
		UserID:    c.UserID,
		PostID:    c.PostID,
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}

func (p *WritePostDTO) ToPost() Post {
	return Post{
		Title:   p.Title,
		Content: p.Content,
		UserID:  p.UserID,
	}
}

func (p *Post) ToReadPostDTO() ReadPostDTO {
	return ReadPostDTO{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

func (u *WriteUserDTO) ToUser() User {
	return User{
		Name:  u.Name,
		Email: u.Email,
	}
}

func (u *User) ToReadUserDTO() ReadUserDTO {
	return ReadUserDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}
