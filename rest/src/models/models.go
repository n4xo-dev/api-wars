package models

import "time"

type User struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `gorm:"unique;not null;index" json:"email"`
	Posts    []Post    `json:"posts"`
	Messages []Message `json:"messages"`
	Comments []Comment `json:"comments"`
	Chats    []*Chat   `gorm:"many2many:participants;" json:"chats"`
}

type Post struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
	UserID   uint      `json:"userId"`
}

type Comment struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
	PostID  uint   `json:"postId"`
}

type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	UserID    uint      `json:"userId"`
	ChatID    uint      `json:"chatId"`
}

type Chat struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Messages     []Message `json:"messages"`
	Participants []*User   `gorm:"many2many:participants;" json:"participants"`
}
