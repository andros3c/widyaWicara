package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type DomainUser struct {
	Id        uint
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Token     string
}

type Usecase interface{
	Login(domain DomainUser,ctx context.Context)(DomainUser,error)
	CreateNewUser(domain DomainUser,ctx context.Context)(DomainUser,error)
	FindById(id int,ctx context.Context)(DomainUser,error)
}

type Repository interface{
	Login(domain DomainUser,ctx context.Context)(DomainUser,error)
	FindUsername(domain DomainUser,ctx context.Context)(DomainUser,error)
	FindEmail(domain DomainUser,ctx context.Context)(DomainUser,error)
	CreateNewUser(domain DomainUser,ctx context.Context)(DomainUser,error)
	FindById(id int,ctx context.Context)(DomainUser,error)	
}