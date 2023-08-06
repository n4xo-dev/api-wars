package db

import "github.com/iLopezosa/api-wars/graphql/graph/model"

// Updates or creates a comment if the id provided within the comment is found or not, respectively
func CommentUpsert(c *model.Comment) error {

	ctx := DBClient.Save(c).Take(c)

	return ctx.Error
}

// Gets the data of the comment with the provided id
func CommentRead(id uint64) (model.Comment, error) {

	var comment = model.Comment{}
	ctx := DBClient.Model(&model.Comment{}).First(&comment, id)

	return comment, ctx.Error
}

// Patch update the comment with the provided id
func CommentPatch(c *model.Comment) error {

	ctx := DBClient.Updates(c).Take(c)

	return ctx.Error
}

// Deletes the comment with the provided id
func CommentDelete(id uint64) error {

	var comment = model.Comment{
		ID: id,
	}
	ctx := DBClient.Delete(&comment)

	return ctx.Error
}

// Gets the data of all the comments
func CommentList() ([]model.Comment, error) {

	var comments []model.Comment
	ctx := DBClient.Model(&model.Comment{}).Find(&comments)

	return comments, ctx.Error
}

// Gets the data of the comments with the provided post id
func CommentListByPostID(postID uint64) ([]model.Comment, error) {

	var comments []model.Comment
	ctx := DBClient.Model(&model.Comment{}).Where("post_id = ?", postID).Find(&comments)

	return comments, ctx.Error
}

// Gets the data of the comments with the provided user id
func CommentListByUserID(userID uint64) ([]model.Comment, error) {

	var comments []model.Comment
	ctx := DBClient.Model(&model.Comment{}).Where("user_id = ?", userID).Find(&comments)

	return comments, ctx.Error
}
