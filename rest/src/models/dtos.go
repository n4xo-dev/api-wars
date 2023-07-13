package models

type UserDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PostDTO struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
}

type CommentDTO struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
	PostID  uint   `json:"postId"`
}

type MessageDTO struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
	ChatID  uint   `json:"chatId"`
}

type ChatDTO struct {
	ID           uint   `json:"id"`
	Participants []User `json:"participants"`
}
