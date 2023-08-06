package db

import "github.com/iLopezosa/api-wars/graphql/graph/model"

// Updater or creates a post if the id provided within the post is found or not, respectively
func PostUpsert(post *model.Post) error {

	ctx := DBClient.Save(post).Take(post)

	return ctx.Error
}

// Gets the data of the post with the provided id
func PostRead(id uint64) (model.Post, error) {

	var post = model.Post{}
	ctx := DBClient.Model(&model.Post{}).First(&post, id)

	return post, ctx.Error
}

// Deletes the post with the provided id
func PostDelete(id uint64) error {

	var post = model.Post{
		ID: id,
	}
	ctx := DBClient.Model(&model.Post{}).Delete(&post)

	return ctx.Error
}

// Patch update the post with the provided id
func PostPatch(post *model.Post) error {

	ctx := DBClient.Updates(post).Take(post)

	return ctx.Error
}

// Gets the data of all the posts
func PostList() ([]*model.Post, error) {

	var posts []*model.Post
	ctx := DBClient.Model(&model.Post{}).Find(&posts)

	return posts, ctx.Error
}

// Gets the data of the posts with the provided user id
func PostListByUserID(userID uint64) ([]*model.Post, error) {

	var posts []*model.Post
	ctx := DBClient.Model(&model.Post{}).Where("user_id = ?", userID).Find(&posts)

	return posts, ctx.Error
}
