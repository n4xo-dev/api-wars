package db

import "github.com/iLopezosa/api-wars/lib/models"

// Updates or creates a comment if the id provided within the comment is found or not, respectively
func CommentUpsert(c *models.Comment) error {

	ctx := DBClient.Save(c).Take(c)

	return ctx.Error
}

// Gets the data of the comment with the provided id
func CommentRead(id uint64) (models.ReadCommentDTO, error) {

	var comment = models.ReadCommentDTO{}
	ctx := DBClient.Model(&models.Comment{}).First(&comment, id)

	return comment, ctx.Error
}

// Patch update the comment with the provided id
func CommentPatch(c *models.Comment) error {

	ctx := DBClient.Updates(c).Take(c)

	return ctx.Error
}

// Deletes the comment with the provided id
func CommentDelete(id uint64) error {

	var comment = models.Comment{
		ID: id,
	}
	ctx := DBClient.Delete(&comment)

	return ctx.Error
}

// Gets the data of all the comments
func CommentList() ([]models.ReadCommentDTO, error) {

	var comments []models.ReadCommentDTO
	ctx := DBClient.Model(&models.Comment{}).Find(&comments)

	return comments, ctx.Error
}

// Gets the data of the comments with the provided post id
func CommentListByPostID(postID uint64) ([]models.ReadCommentDTO, error) {

	var comments []models.ReadCommentDTO
	ctx := DBClient.Model(&models.Comment{}).Where("post_id = ?", postID).Find(&comments)

	return comments, ctx.Error
}

// Gets the data of the comments with the provided user id
func CommentListByUserID(userID uint64) ([]models.ReadCommentDTO, error) {

	var comments []models.ReadCommentDTO
	ctx := DBClient.Model(&models.Comment{}).Where("user_id = ?", userID).Find(&comments)

	return comments, ctx.Error
}
