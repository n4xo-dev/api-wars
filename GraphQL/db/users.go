package db

import "github.com/iLopezosa/api-wars/graphql/graph/model"

// Updates or creates a user if the id provided within the user is found or not, respectively
func UserUpsert(u *model.User) error {

	ctx := DBClient.Save(u).Take(u)

	return ctx.Error
}

// Gets the data of the user with the provided id
func UserRead(id uint64) (model.User, error) {

	user := model.User{}
	ctx := DBClient.Model(&model.User{}).First(&user, id)

	return user, ctx.Error
}

// Patch update the user with the provided id
func UserPatch(u *model.User) error {

	ctx := DBClient.Updates(u).Take(u)

	return ctx.Error
}

// Deletes the user with the provided id
func UserDelete(id uint64) error {

	var user = model.User{
		ID: id,
	}
	ctx := DBClient.Delete(&user)

	return ctx.Error
}

// Gets the data of all the users
func UserList() ([]*model.User, error) {

	var users []*model.User
	ctx := DBClient.Model(&model.User{}).Find(&users)

	return users, ctx.Error
}

// Gets the data of the user with the provided email
func UserFindByEmail(email string) ([]model.User, error) {

	var users []model.User
	ctx := DBClient.Model(&model.User{}).Where("email = ?", email).Find(&users)

	return users, ctx.Error
}
