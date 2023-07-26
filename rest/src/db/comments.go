package db

import "github.com/iLopezosa/api-wars/rest/src/models"

// Updates or creates a comment if the id provided within the comment is found or not, respectively
func CommentUpsert(c *models.Comment) error {

	ctx := DBClient.Save(c)

	return ctx.Error
}

// Gets the data of the comment with the provided id
func CommentRead(id uint) (models.Comment, error) {

	var comment = models.Comment{
		ID: id,
	}
	ctx := DBClient.First(&comment)

	return comment, ctx.Error
}

// Deletes the comment with the provided id
func CommentDelete(id uint) error {

	var comment = models.Comment{
		ID: id,
	}
	ctx := DBClient.Delete(&comment)

	return ctx.Error
}

// Gets the data of all the comments
func CommentList() ([]models.Comment, error) {

	var comments []models.Comment
	ctx := DBClient.Find(&comments)

	return comments, ctx.Error
}

// Gets the data of the comments with the provided post id
func CommentListByPostID(postID uint) ([]models.Comment, error) {

	var comments []models.Comment
	ctx := DBClient.Where("post_id = ?", postID).Find(&comments)

	return comments, ctx.Error
}

// Gets the data of the comments with the provided user id
func CommentListByUserID(userID uint) ([]models.Comment, error) {

	var comments []models.Comment
	ctx := DBClient.Where("user_id = ?", userID).Find(&comments)

	return comments, ctx.Error
}