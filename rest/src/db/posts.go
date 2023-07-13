package db

import "github.com/iLopezosa/api-wars/rest/src/models"

// Updater or creates a post if the id provided within the post is found or not, respectively
func PostUpsert(post *models.Post) error {

	ctx := DBClient.Save(post)

	return ctx.Error
}

// Gets the data of the post with the provided id
func PostRead(id uint) (models.Post, error) {

	var post = models.Post{
		ID: id,
	}
	ctx := DBClient.First(&post)

	return post, ctx.Error
}

// Deletes the post with the provided id
func PostDelete(id uint) error {

	var post = models.Post{
		ID: id,
	}
	ctx := DBClient.Delete(&post)

	return ctx.Error
}

// Gets the data of all the posts
func PostList() ([]models.Post, error) {

	var posts []models.Post
	ctx := DBClient.Find(&posts)

	return posts, ctx.Error
}

// Gets the data of the posts with the provided user id
func PostListByUserID(userID uint) ([]models.Post, error) {

	var posts []models.Post
	ctx := DBClient.Where("user_id = ?", userID).Find(&posts)

	return posts, ctx.Error
}
