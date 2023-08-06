package model

type NewUser struct {
	WriteUserDTO
}

type NewPost struct {
	WritePostDTO
}

type NewComment struct {
	WriteCommentDTO
}

type NewMessage struct {
	WriteMessageDTO
}

type NewChat struct {
	Participants []uint64
}
