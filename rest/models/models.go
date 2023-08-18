package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `gorm:"unique;not null;index" json:"email"`
	Posts     []Post         `json:"posts"`
	Messages  []Message      `json:"messages"`
	Comments  []Comment      `json:"comments"`
	Chats     []*Chat        `gorm:"many2many:participants;" json:"chats"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type Post struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Comments  []Comment      `json:"comments"`
	UserID    uint64         `gorm:"not null; index" json:"userId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type Comment struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	Content   string         `json:"content"`
	UserID    uint64         `gorm:"not null" json:"userId"`
	PostID    uint64         `gorm:"not null" json:"postId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type Message struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	Content   string         `json:"content"`
	UserID    uint64         `gorm:"not null" json:"userId"`
	ChatID    uint64         `gorm:"not null" json:"chatId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type Chat struct {
	ID           uint64         `gorm:"primaryKey" json:"id"`
	Messages     []Message      `json:"messages"`
	Participants []*User        `gorm:"many2many:participants;" json:"participants"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
}

type RedisRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
