package db

import "github.com/iLopezosa/api-wars/rest/src/models"

// Updates or creates a user if the id provided within the user is found or not, respectively
func UserUpsert(u *models.User) error {

	ctx := DBClient.Save(u)

	return ctx.Error
}

// Gets the data of the user with the provided id
func UserRead(id uint) (models.UserDTO, error) {

	user := models.UserDTO{}
	ctx := DBClient.Model(&models.User{}).Find(&models.UserDTO{}).First(&user, id)

	return user, ctx.Error
}

// Patch update the user with the provided id
func UserPatch(u *models.UserDTO) (models.UserDTO, error) {

	ctx := DBClient.Model(&models.User{}).Updates(u)

	return *u, ctx.Error
}

// Deletes the user with the provided id
func UserDelete(id uint) error {

	var user = models.User{
		ID: id,
	}
	ctx := DBClient.Delete(&user)

	return ctx.Error
}

// Gets the data of all the users
func UserList() ([]models.UserDTO, error) {

	var users []models.UserDTO
	ctx := DBClient.Model(&models.User{}).Find(&users)

	return users, ctx.Error
}

// Gets the data of the user with the provided email
func UserFindByEmail(email string) ([]models.UserDTO, error) {

	var users []models.UserDTO
	ctx := DBClient.Model(&models.User{}).Where("email = ?", email).Find(&users)

	return users, ctx.Error
}
