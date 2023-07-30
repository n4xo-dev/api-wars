package models

func (m *WriteMessageDTO) ToMessage() Message {
	return Message{
		Content: m.Content,
		UserID:  m.UserID,
		ChatID:  m.ChatID,
	}
}
