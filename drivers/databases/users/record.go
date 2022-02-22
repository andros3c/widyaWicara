package users

import (
	"time"
	"widyaWicaraBackend/businesses/users"

	"gorm.io/gorm"
)

type User struct {
	Id        uint
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time		`gorm:"<-:create"`
	UpdatedAt time.Time		`gorm:"<-:update"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	
}

func FromDomain(domain users.DomainUser)User{
	return User{
		Id: domain.Id,
		Username: domain.Username,
		Password: domain.Password,
		Email: domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (user User)ToDomain()(users.DomainUser){
	return users.DomainUser{
		Id : user.Id,
		Username: user.Username,
		Password: user.Password,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,

	}
}