package db

import "github.com/n4xo-dev/api-wars/lib/models"

// Updates or creates a user if the id provided within the user is found or not, respectively
func UserUpsert(u *models.User) error {

	ctx := DBClient.Save(u).Take(u)

	return ctx.Error
}

// Gets the data of the user with the provided id
func UserRead(id uint64) (models.User, error) {

	user := models.User{}
	ctx := DBClient.Model(&models.User{}).First(&user, id)

	return user, ctx.Error
}

// Patch update the user with the provided id
func UserPatch(u *models.User) error {

	ctx := DBClient.Updates(u).Take(u)

	return ctx.Error
}

// Deletes the user with the provided id
func UserDelete(id uint64) error {

	var user = models.User{
		ID: id,
	}
	ctx := DBClient.Delete(&user)

	return ctx.Error
}

// Gets the data of all the users
func UserList() ([]*models.User, error) {

	var users []*models.User
	ctx := DBClient.Model(&models.User{}).Find(&users)

	return users, ctx.Error
}

// Gets the data of the user with the provided email
func UserFindByEmail(email string) ([]models.User, error) {

	var users []models.User
	ctx := DBClient.Model(&models.User{}).Where("email = ?", email).Find(&users)

	return users, ctx.Error
}
