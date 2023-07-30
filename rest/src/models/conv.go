package models

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
