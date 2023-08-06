package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"strconv"

	"github.com/iLopezosa/api-wars/graphql/db"
	"github.com/iLopezosa/api-wars/graphql/graph/model"
)

// ID is the resolver for the id field.
func (r *chatResolver) ID(ctx context.Context, obj *model.Chat) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *chatResolver) CreatedAt(ctx context.Context, obj *model.Chat) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *chatResolver) UpdatedAt(ctx context.Context, obj *model.Chat) (*string, error) {
	t := obj.UpdatedAt.String()
	return &t, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *chatResolver) DeletedAt(ctx context.Context, obj *model.Chat) (*string, error) {
	t := obj.DeletedAt.Time.String()
	return &t, nil
}

// ID is the resolver for the id field.
func (r *commentResolver) ID(ctx context.Context, obj *model.Comment) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

// UserID is the resolver for the userId field.
func (r *commentResolver) UserID(ctx context.Context, obj *model.Comment) (string, error) {
	return fmt.Sprintf("%d", obj.UserID), nil
}

// PostID is the resolver for the postId field.
func (r *commentResolver) PostID(ctx context.Context, obj *model.Comment) (string, error) {
	return fmt.Sprintf("%d", obj.PostID), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *commentResolver) CreatedAt(ctx context.Context, obj *model.Comment) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *commentResolver) UpdatedAt(ctx context.Context, obj *model.Comment) (*string, error) {
	t := obj.UpdatedAt.String()
	return &t, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *commentResolver) DeletedAt(ctx context.Context, obj *model.Comment) (*string, error) {
	t := obj.DeletedAt.Time.String()
	return &t, nil
}

// ID is the resolver for the id field.
func (r *messageResolver) ID(ctx context.Context, obj *model.Message) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

// UserID is the resolver for the userId field.
func (r *messageResolver) UserID(ctx context.Context, obj *model.Message) (string, error) {
	return fmt.Sprintf("%d", obj.UserID), nil
}

// ChatID is the resolver for the chatId field.
func (r *messageResolver) ChatID(ctx context.Context, obj *model.Message) (string, error) {
	return fmt.Sprintf("%d", obj.ChatID), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *messageResolver) CreatedAt(ctx context.Context, obj *model.Message) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *messageResolver) UpdatedAt(ctx context.Context, obj *model.Message) (*string, error) {
	t := obj.UpdatedAt.String()
	return &t, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *messageResolver) DeletedAt(ctx context.Context, obj *model.Message) (*string, error) {
	t := obj.DeletedAt.Time.String()
	return &t, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// CreateMessage is the resolver for the createMessage field.
func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	panic(fmt.Errorf("not implemented: CreateMessage - createMessage"))
}

// CreateChat is the resolver for the createChat field.
func (r *mutationResolver) CreateChat(ctx context.Context, input model.NewChat) (*model.Chat, error) {
	panic(fmt.Errorf("not implemented: CreateChat - createChat"))
}

// ID is the resolver for the id field.
func (r *postResolver) ID(ctx context.Context, obj *model.Post) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

// UserID is the resolver for the userId field.
func (r *postResolver) UserID(ctx context.Context, obj *model.Post) (string, error) {
	return fmt.Sprintf("%d", obj.UserID), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *postResolver) CreatedAt(ctx context.Context, obj *model.Post) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *postResolver) UpdatedAt(ctx context.Context, obj *model.Post) (*string, error) {
	t := obj.UpdatedAt.String()
	return &t, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *postResolver) DeletedAt(ctx context.Context, obj *model.Post) (*string, error) {
	t := obj.DeletedAt.Time.String()
	return &t, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := db.UserList()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	uID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := db.UserRead(uID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	posts, err := db.PostList()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	pID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	post, err := db.PostRead(pID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	comments, err := db.CommentList()
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, id string) (*model.Comment, error) {
	cID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	comment, err := db.CommentRead(cID)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	messages, err := db.MessageList()
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// Message is the resolver for the message field.
func (r *queryResolver) Message(ctx context.Context, id string) (*model.Message, error) {
	mID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	message, err := db.MessageRead(mID)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// Chats is the resolver for the chats field.
func (r *queryResolver) Chats(ctx context.Context) ([]*model.Chat, error) {
	chats, err := db.ChatList(false)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

// Chat is the resolver for the chat field.
func (r *queryResolver) Chat(ctx context.Context, id string) (*model.Chat, error) {
	cID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	chat, err := db.ChatRead(cID, false)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *model.User) (*string, error) {
	t := obj.UpdatedAt.String()
	return &t, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *userResolver) DeletedAt(ctx context.Context, obj *model.User) (*string, error) {
	t := obj.DeletedAt.Time.String()
	return &t, nil
}

// Participants is the resolver for the participants field.
func (r *newChatResolver) Participants(ctx context.Context, obj *model.NewChat, data []string) error {
	uIDs := make([]uint64, len(data))

	for i, id := range data {
		newId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return err
		}
		uIDs[i] = newId
	}

	obj.Participants = uIDs
	return nil
}

// UserID is the resolver for the userId field.
func (r *newCommentResolver) UserID(ctx context.Context, obj *model.NewComment, data string) error {
	var err error
	obj.UserID, err = strconv.ParseUint(data, 10, 64)
	return err
}

// PostID is the resolver for the postId field.
func (r *newCommentResolver) PostID(ctx context.Context, obj *model.NewComment, data string) error {
	var err error
	obj.PostID, err = strconv.ParseUint(data, 10, 64)
	return err
}

// UserID is the resolver for the userId field.
func (r *newMessageResolver) UserID(ctx context.Context, obj *model.NewMessage, data string) error {
	var err error
	obj.UserID, err = strconv.ParseUint(data, 10, 64)
	return err
}

// ChatID is the resolver for the chatId field.
func (r *newMessageResolver) ChatID(ctx context.Context, obj *model.NewMessage, data string) error {
	var err error
	obj.ChatID, err = strconv.ParseUint(data, 10, 64)
	return err
}

// UserID is the resolver for the userId field.
func (r *newPostResolver) UserID(ctx context.Context, obj *model.NewPost, data string) error {
	var err error
	obj.UserID, err = strconv.ParseUint(data, 10, 64)
	return err
}

// Chat returns ChatResolver implementation.
func (r *Resolver) Chat() ChatResolver { return &chatResolver{r} }

// Comment returns CommentResolver implementation.
func (r *Resolver) Comment() CommentResolver { return &commentResolver{r} }

// Message returns MessageResolver implementation.
func (r *Resolver) Message() MessageResolver { return &messageResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// NewChat returns NewChatResolver implementation.
func (r *Resolver) NewChat() NewChatResolver { return &newChatResolver{r} }

// NewComment returns NewCommentResolver implementation.
func (r *Resolver) NewComment() NewCommentResolver { return &newCommentResolver{r} }

// NewMessage returns NewMessageResolver implementation.
func (r *Resolver) NewMessage() NewMessageResolver { return &newMessageResolver{r} }

// NewPost returns NewPostResolver implementation.
func (r *Resolver) NewPost() NewPostResolver { return &newPostResolver{r} }

type chatResolver struct{ *Resolver }
type commentResolver struct{ *Resolver }
type messageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type newChatResolver struct{ *Resolver }
type newCommentResolver struct{ *Resolver }
type newMessageResolver struct{ *Resolver }
type newPostResolver struct{ *Resolver }
