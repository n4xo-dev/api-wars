package model

func (m *WriteMessageDTO) ToMessage() Message {
	return Message{
		Content: m.Content,
		UserID:  m.UserID,
		ChatID:  m.ChatID,
	}
}

func (c *WriteCommentDTO) ToComment() Comment {
	return Comment{
		Content: c.Content,
		UserID:  c.UserID,
		PostID:  c.PostID,
	}
}

func (p *WritePostDTO) ToPost() Post {
	return Post{
		Title:   p.Title,
		Content: p.Content,
		UserID:  p.UserID,
	}
}

func (u *WriteUserDTO) ToUser() User {
	return User{
		Name:  u.Name,
		Email: u.Email,
	}
}
