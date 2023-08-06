package model

type WriteUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type WritePostDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint64 `json:"userId"`
}

type WriteCommentDTO struct {
	Content string `json:"content"`
	UserID  uint64 `json:"userId"`
	PostID  uint64 `json:"postId"`
}

type WriteMessageDTO struct {
	Content string `json:"content"`
	UserID  uint64 `json:"userId"`
	ChatID  uint64 `json:"chatId"`
}
