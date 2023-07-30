package models

type ReadUserDTO struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type WriteUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ReadPostDTO struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    uint64 `json:"userId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type WritePostDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint64 `json:"userId"`
}

type ReadCommentDTO struct {
	ID        uint64 `json:"id"`
	Content   string `json:"content"`
	UserID    uint64 `json:"userId"`
	PostID    uint64 `json:"postId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type WriteCommentDTO struct {
	Content string `json:"content"`
	UserID  uint64 `json:"userId"`
	PostID  uint64 `json:"postId"`
}

type ReadMessageDTO struct {
	ID        uint64 `json:"id"`
	Content   string `json:"content"`
	UserID    uint64 `json:"userId"`
	ChatID    uint64 `json:"chatId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type WriteMessageDTO struct {
	Content string `json:"content"`
	UserID  uint64 `json:"userId"`
	ChatID  uint64 `json:"chatId"`
}
