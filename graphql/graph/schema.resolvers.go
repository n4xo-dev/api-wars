package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"strconv"

	"github.com/n4xo-dev/api-wars/graphql/graph/model"
	"github.com/n4xo-dev/api-wars/lib/db"
	"github.com/n4xo-dev/api-wars/lib/models"
)

// ID is the resolver for the id field.
func (r *chatResolver) ID(ctx context.Context, obj *models.Chat) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

// Messages is the resolver for the messages field.
func (r *chatResolver) Messages(ctx context.Context, obj *models.Chat) ([]*models.Message, error) {
	return db.ChatMessages(obj.ID)
}

// Participants is the resolver for the participants field.
func (r *chatResolver) Participants(ctx context.Context, obj *models.Chat) ([]*models.User, error) {
	return db.ChatParticipants(obj.ID)
}

// CreatedAt is the resolver for the createdAt field.
func (r *chatResolver) CreatedAt(ctx context.Context, obj *models.Chat) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *chatResolver) UpdatedAt(ctx context.Context, obj *models.Chat) (*string, error) {
	t := obj.UpdatedAt.String()
	return &t, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *chatResolver) DeletedAt(ctx context.Context, obj *models.Chat) (*string, error) {
	t := obj.DeletedAt.Time.String()
	return &t, nil
}

// ID is the resolver for the id field.
func (r *commentResolver) ID(ctx context.Context, obj *models.Comment) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

// UserID is the resolver for the userId field.
func (r *commentResolver) UserID(ctx context.Context, obj *models.Comment) (string, error) {
	return fmt.Sprintf("%d", obj.UserID), nil
}

// PostID is the resolver for the postId field.
func (r *commentResolver) PostID(ctx context.Context, obj *models.Comment) (string, error) {
	return fmt.Sprintf("%d", obj.PostID), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *commentResolver) CreatedAt(ctx context.Context, obj *models.Comment) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *commentResolver) UpdatedAt(ctx context.Context, obj *models.Comment) (*string, error) {
	t := obj.UpdatedAt.String()
	return &t, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *commentResolver) DeletedAt(ctx context.Context, obj *models.Comment) (*string, error) {
	t := obj.DeletedAt.Time.String()
	return &t, nil
}

// ID is the resolver for the id field.
func (r *messageResolver) ID(ctx context.Context, obj *models.Message) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

// UserID is the resolver for the userId field.
func (r *messageResolver) UserID(ctx context.Context, obj *models.Message) (string, error) {
	return fmt.Sprintf("%d", obj.UserID), nil
}

// ChatID is the resolver for the chatId field.
func (r *messageResolver) ChatID(ctx context.Context, obj *models.Message) (string, error) {
	return fmt.Sprintf("%d", obj.ChatID), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *messageResolver) CreatedAt(ctx context.Context, obj *models.Message) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *messageResolver) UpdatedAt(ctx context.Context, obj *models.Message) (*string, error) {
	t := obj.UpdatedAt.String()
	return &t, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *messageResolver) DeletedAt(ctx context.Context, obj *models.Message) (*string, error) {
	t := obj.DeletedAt.Time.String()
	return &t, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	u := input.WriteUserDTO.ToUser()

	if err := db.UserUpsert(&u); err != nil {
		return nil, err
	}

	return &u, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*models.Post, error) {
	p := input.WritePostDTO.ToPost()

	if err := db.PostUpsert(&p); err != nil {
		return nil, err
	}

	return &p, nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*models.Comment, error) {
	c := input.WriteCommentDTO.ToComment()

	if err := db.CommentUpsert(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// CreateMessage is the resolver for the createMessage field.
func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*models.Message, error) {
	m := input.WriteMessageDTO.ToMessage()

	if err := db.MessageUpsert(&m); err != nil {
		return nil, err
	}

	return &m, nil
}

// CreateChat is the resolver for the createChat field.
func (r *mutationResolver) CreateChat(ctx context.Context, input model.NewChat) (*models.Chat, error) {
	c := models.Chat{
		Participants: make([]*models.User, len(input.Participants)),
	}

	for i, id := range input.Participants {
		c.Participants[i] = &models.User{ID: id}
	}

	if err := db.ChatUpsert(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// CreateRedisRecord is the resolver for the createRedisRecord field.
func (r *mutationResolver) CreateRedisRecord(ctx context.Context, input model.NewRedisRecord) (*models.RedisRecord, error) {
	if err := db.RedisSet(input.Key, input.Value); err != nil {
		return nil, err
	}

	return &models.RedisRecord{
		Key:   input.Key,
		Value: input.Value,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUser) (*models.User, error) {
	uID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	if uID < 1 {
		return nil, fmt.Errorf("id must be greater than 0")
	}

	u := models.User{
		ID: uID,
	}
	if input.Name != nil {
		u.Name = *input.Name
	}
	if input.Email != nil {
		u.Email = *input.Email
	}

	if err := db.UserPatch(&u); err != nil {
		return nil, err
	}

	return &u, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input model.UpdatePost) (*models.Post, error) {
	pID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	if pID < 1 {
		return nil, fmt.Errorf("id must be greater than 0")
	}

	p := models.Post{
		ID: pID,
	}
	if input.Title != nil {
		p.Title = *input.Title
	}
	if input.Content != nil {
		p.Content = *input.Content
	}

	if err := db.PostPatch(&p); err != nil {
		return nil, err
	}

	return &p, nil
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input model.UpdateComment) (*models.Comment, error) {
	cID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	if cID < 1 {
		return nil, fmt.Errorf("id must be greater than 0")
	}

	c := models.Comment{
		ID: cID,
	}
	if input.Content != nil {
		c.Content = *input.Content
	}

	if err := db.CommentPatch(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// UpdateMessage is the resolver for the updateMessage field.
func (r *mutationResolver) UpdateMessage(ctx context.Context, id string, input model.UpdateMessage) (*models.Message, error) {
	mID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	if mID < 1 {
		return nil, fmt.Errorf("id must be greater than 0")
	}

	m := models.Message{
		ID: mID,
	}
	if input.Content != nil {
		m.Content = *input.Content
	}

	if err := db.MessagePatch(&m); err != nil {
		return nil, err
	}

	return &m, nil
}

// UpdateRedisRecord is the resolver for the updateRedisRecord field.
func (r *mutationResolver) UpdateRedisRecord(ctx context.Context, key string, value string) (*models.RedisRecord, error) {
	if err := db.RedisSet(key, value); err != nil {
		return nil, err
	}

	return &models.RedisRecord{
		Key:   key,
		Value: value,
	}, nil
}

// AddUsersToChat is the resolver for the addUsersToChat field.
func (r *mutationResolver) AddUsersToChat(ctx context.Context, chatID string, userIds []string) (*models.Chat, error) {
	cID, err := strconv.ParseUint(chatID, 10, 64)
	if err != nil {
		return nil, err
	}
	if cID < 1 {
		return nil, fmt.Errorf("id must be greater than 0")
	}

	uIDs := make([]uint64, len(userIds))
	for i, id := range userIds {
		uIDs[i], err = strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	currUs, err := db.ChatParticipants(cID)
	if err != nil {
		return nil, err
	}

	newUs := make([]*models.User, len(uIDs))
	for i, id := range uIDs {
		newUs[i] = &models.User{ID: id}
	}

	c := models.Chat{
		ID:           cID,
		Participants: append(currUs, newUs...),
	}

	if err := db.ChatPatch(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.Deletion, error) {
	uID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	if err = db.UserDelete(uID); err != nil {
		return nil, err
	}

	return &model.Deletion{ID: id, Msg: fmt.Sprintf("User #%d has been deleted succesfully.", uID)}, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*model.Deletion, error) {
	pID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	if err = db.PostDelete(pID); err != nil {
		return nil, err
	}

	return &model.Deletion{ID: id, Msg: fmt.Sprintf("Post #%d has been deleted succesfully.", pID)}, nil
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (*model.Deletion, error) {
	cID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	if err = db.CommentDelete(cID); err != nil {
		return nil, err
	}

	return &model.Deletion{ID: id, Msg: fmt.Sprintf("Comment #%d has been deleted succesfully.", cID)}, nil
}

// DeleteMessage is the resolver for the deleteMessage field.
func (r *mutationResolver) DeleteMessage(ctx context.Context, id string) (*model.Deletion, error) {
	mID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	if err = db.MessageDelete(mID); err != nil {
		return nil, err
	}

	return &model.Deletion{ID: id, Msg: fmt.Sprintf("Message #%d has been deleted succesfully.", mID)}, nil
}

// DeleteChat is the resolver for the deleteChat field.
func (r *mutationResolver) DeleteChat(ctx context.Context, id string) (*model.Deletion, error) {
	cID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	if err = db.ChatDelete(cID); err != nil {
		return nil, err
	}

	return &model.Deletion{ID: id, Msg: fmt.Sprintf("Chat #%d has been deleted succesfully.", cID)}, nil
}

// ID is the resolver for the id field.
func (r *postResolver) ID(ctx context.Context, obj *models.Post) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *models.Post) ([]*models.Comment, error) {
	return db.PostComments(obj.ID)
}

// UserID is the resolver for the userId field.
func (r *postResolver) UserID(ctx context.Context, obj *models.Post) (string, error) {
	return fmt.Sprintf("%d", obj.UserID), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *postResolver) CreatedAt(ctx context.Context, obj *models.Post) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *postResolver) UpdatedAt(ctx context.Context, obj *models.Post) (*string, error) {
	t := obj.UpdatedAt.String()
	return &t, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *postResolver) DeletedAt(ctx context.Context, obj *models.Post) (*string, error) {
	t := obj.DeletedAt.Time.String()
	return &t, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	users, err := db.FullUserList()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	uID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := db.FullUserRead(uID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserByEmail is the resolver for the userByEmail field.
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*models.User, error) {
	users, err := db.FullUserFindByEmail(email)
	if err != nil {
		return nil, err
	}

	return &users[0], nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	posts, err := db.FullPostList()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	pID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	post, err := db.FullPostRead(pID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// PostsByUser is the resolver for the postsByUser field.
func (r *queryResolver) PostsByUser(ctx context.Context, userID string) ([]*models.Post, error) {
	uID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, err
	}

	posts, err := db.FullPostListByUserID(uID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context) ([]*models.Comment, error) {
	comments, err := db.FullCommentList()
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, id string) (*models.Comment, error) {
	cID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	comment, err := db.FullCommentRead(cID)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// CommentsByUser is the resolver for the commentsByUser field.
func (r *queryResolver) CommentsByUser(ctx context.Context, userID string) ([]*models.Comment, error) {
	uID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, err
	}

	comments, err := db.FullCommentListByUserID(uID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// CommentsByPost is the resolver for the commentsByPost field.
func (r *queryResolver) CommentsByPost(ctx context.Context, postID string) ([]*models.Comment, error) {
	pID, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		return nil, err
	}

	comments, err := db.FullCommentListByPostID(pID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*models.Message, error) {
	messages, err := db.FullMessageList()
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// Message is the resolver for the message field.
func (r *queryResolver) Message(ctx context.Context, id string) (*models.Message, error) {
	mID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	message, err := db.FullMessageRead(mID)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// MessagesByUser is the resolver for the messagesByUser field.
func (r *queryResolver) MessagesByUser(ctx context.Context, userID string) ([]*models.Message, error) {
	uID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, err
	}

	messages, err := db.FullMessageListByUserID(uID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// MessagesByChat is the resolver for the messagesByChat field.
func (r *queryResolver) MessagesByChat(ctx context.Context, chatID string) ([]*models.Message, error) {
	cID, err := strconv.ParseUint(chatID, 10, 64)
	if err != nil {
		return nil, err
	}

	messages, err := db.FullMessageListByChatID(cID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// MessagesByChatAndUser is the resolver for the messagesByChatAndUser field.
func (r *queryResolver) MessagesByChatAndUser(ctx context.Context, chatID string, userID string) ([]*models.Message, error) {
	cID, err := strconv.ParseUint(chatID, 10, 64)
	if err != nil {
		return nil, err
	}

	uID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, err
	}

	messages, err := db.FullMessageListByChatIDAndUserID(cID, uID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// Chats is the resolver for the chats field.
func (r *queryResolver) Chats(ctx context.Context) ([]*models.Chat, error) {
	chats, err := db.ChatList(false)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

// Chat is the resolver for the chat field.
func (r *queryResolver) Chat(ctx context.Context, id string) (*models.Chat, error) {
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

// RedisRecord is the resolver for the redisRecord field.
func (r *queryResolver) RedisRecord(ctx context.Context, key string) (*models.RedisRecord, error) {
	val, err := db.RedisGet(key)
	if err != nil {
		return nil, err
	}

	return &models.RedisRecord{
		Key:   key,
		Value: val,
	}, nil
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return fmt.Sprintf("%d", obj.ID), nil
}

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *models.User) ([]*models.Post, error) {
	return db.FullPostListByUserID(obj.ID)
}

// Messages is the resolver for the messages field.
func (r *userResolver) Messages(ctx context.Context, obj *models.User) ([]*models.Message, error) {
	return db.FullMessageListByUserID(obj.ID)
}

// Comments is the resolver for the comments field.
func (r *userResolver) Comments(ctx context.Context, obj *models.User) ([]*models.Comment, error) {
	return db.FullCommentListByUserID(obj.ID)
}

// Chats is the resolver for the chats field.
func (r *userResolver) Chats(ctx context.Context, obj *models.User) ([]*models.Chat, error) {
	return db.ChatListByUserID(obj.ID)
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.CreatedAt.String(), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (*string, error) {
	t := obj.UpdatedAt.String()
	return &t, nil
}

// DeletedAt is the resolver for the deletedAt field.
func (r *userResolver) DeletedAt(ctx context.Context, obj *models.User) (*string, error) {
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
