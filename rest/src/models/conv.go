package models

func ParseUserDTO(u *UserDTO) *User {
	return &User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
