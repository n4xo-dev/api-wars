package db

import "github.com/iLopezosa/api-wars/rest/src/models"

// Updater or creates a post if the id provided within the post is found or not, respectively
func PostUpsert(post *models.Post) error {

	ctx := DBClient.Save(post)

	return ctx.Error
}

// Gets the data of the post with the provided id
func PostRead(id uint64) (models.PostDTO, error) {

	var post = models.PostDTO{}
	ctx := DBClient.Model(&models.Post{}).First(&post, id)

	return post, ctx.Error
}

// Deletes the post with the provided id
func PostDelete(id uint64) error {

	var post = models.Post{
		ID: id,
	}
	ctx := DBClient.Model(&models.Post{}).Delete(&post)

	return ctx.Error
}

// Gets the data of all the posts
func PostList() ([]models.PostDTO, error) {

	var posts []models.PostDTO
	ctx := DBClient.Model(&models.Post{}).Find(&posts)

	return posts, ctx.Error
}

// Gets the data of the posts with the provided user id
func PostListByUserID(userID uint64) ([]models.PostDTO, error) {

	var posts []models.PostDTO
	ctx := DBClient.Model(&models.Post{}).Where("user_id = ?", userID).Find(&posts)

	return posts, ctx.Error
}
