package response

import (
	"time"
	"widyaWicaraBackend/businesses/users"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id        uint				 `json:"id"`
	Username  string			 `json:"username"`
	Email     string 			 `json:"email"`
	CreatedAt time.Time			 `json:"created_at"`
	UpdatedAt time.Time 		 `json:"updated_at"`
	DeletedAt gorm.DeletedAt 	 `json:"deleted_at"` 
	Token     string 			 `json:"token"`
}

func FromDomain(domain users.DomainUser)UserResponse{
	return UserResponse{
		Id: domain.Id,
		Username: domain.Username,
		Email: domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Token: domain.Token,
	}
}