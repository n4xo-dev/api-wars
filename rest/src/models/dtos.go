package models

type UserDTO struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type PostDTO struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    uint64 `json:"userId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CommentDTO struct {
	ID        uint64 `json:"id"`
	Content   string `json:"content"`
	UserID    uint64 `json:"userId"`
	PostID    uint64 `json:"postId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type MessageDTO struct {
	ID        uint64 `json:"id"`
	Content   string `json:"content"`
	UserID    uint64 `json:"userId"`
	ChatID    uint64 `json:"chatId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ChatDTO struct {
	ID        uint64 `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
