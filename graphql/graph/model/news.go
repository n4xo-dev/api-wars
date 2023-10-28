package model

import "github.com/iLopezosa/api-wars/lib/models"

type NewUser struct {
	models.WriteUserDTO
}

type NewPost struct {
	models.WritePostDTO
}

type NewComment struct {
	models.WriteCommentDTO
}

type NewMessage struct {
	models.WriteMessageDTO
}

type NewChat struct {
	Participants []uint64
}

type NewRedisRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
