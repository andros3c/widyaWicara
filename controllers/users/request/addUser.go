package request

import "widyaWicaraBackend/businesses/users"

type CreateNewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (user *CreateNewUser) ToDomain() *users.DomainUser {
	return &users.DomainUser{
		Email:       user.Email,
		Password:    user.Password,
		Username:    user.Username,
	
	}
}