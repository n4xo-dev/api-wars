package db

import "github.com/n4xo-dev/api-wars/lib/models"

// Updater or creates a post if the id provided within the post is found or not, respectively
func PostUpsert(post *models.Post) error {

	ctx := DBClient.Save(post).Take(post)

	return ctx.Error
}

// Gets the data of the post with the provided id
func PostRead(id uint64) (models.ReadPostDTO, error) {

	var post = models.ReadPostDTO{}
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

// Patch update the post with the provided id
func PostPatch(post *models.Post) error {

	ctx := DBClient.Updates(post).Take(post)

	return ctx.Error
}

// Gets the data of all the posts
func PostList() ([]models.ReadPostDTO, error) {

	var posts []models.ReadPostDTO
	ctx := DBClient.Model(&models.Post{}).Find(&posts)

	return posts, ctx.Error
}

// Gets the data of the posts with the provided user id
func PostListByUserID(userID uint64) ([]models.ReadPostDTO, error) {

	var posts []models.ReadPostDTO
	ctx := DBClient.Model(&models.Post{}).Where("user_id = ?", userID).Find(&posts)

	return posts, ctx.Error
}
