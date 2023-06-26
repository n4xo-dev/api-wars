package db

import "github.com/iLopezosa/api-wars/rest/src/models"

// Updater or creates a user if the id provided within the user is found or not, respectively
func UserUpsert(u *models.User) error {

	ctx := DBClient.Save(u)

	return ctx.Error
}

// Gets the data of the user with the provided id
func UserRead(id uint) (models.User, error) {

	var user = models.User{
		ID: id,
	}
	ctx := DBClient.First(&user)

	return user, ctx.Error
}

func UserDelete(id uint) error {

	var user = models.User{
		ID: id,
	}
	ctx := DBClient.Delete(&user)

	return ctx.Error
}
