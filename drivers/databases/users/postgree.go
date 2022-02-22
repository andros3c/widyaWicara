package users

import (
	"context"
	"widyaWicaraBackend/businesses"
	"widyaWicaraBackend/businesses/users"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB)users.Repository{
	return &UserRepository{
		db: gormDb,
	}
}

func (repo *UserRepository)Login(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)

	err:= repo.db.Where("email = ?",userDb.Email).First(&userDb).Error

	if err != nil{
		return users.DomainUser{},err
	}
	return userDb.ToDomain(),nil

}
func(repo *UserRepository) FindUsername(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)
	user := User{}
	// err:= repo.db.Where("username = ?",userDb.Username).First(&userDb).Error
	err:= repo.db.First(&user,"username = ?",userDb.Username).Error
	
	if err == gorm.ErrRecordNotFound{
		return userDb.ToDomain(),gorm.ErrRecordNotFound
	}else if err == nil{
		return users.DomainUser{},businesses.ErrUsernameExisted
	}else{
		return users.DomainUser{},err
	}
}
func(repo *UserRepository) FindEmail(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	user := User{}
	userDb := FromDomain(domain)
	// db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	err:= repo.db.First(&user,"email = ?",userDb.Email).Error
	if err == gorm.ErrRecordNotFound{
		return userDb.ToDomain(),gorm.ErrRecordNotFound
	}else if err == nil{
		return users.DomainUser{},businesses.ErrEmailExisted
	}else{
		return users.DomainUser{},err
	}
}

func (repo *UserRepository) CreateNewUser(domain users.DomainUser,ctx context.Context)(users.DomainUser,error){
	userDb := FromDomain(domain)

	err := repo.db.Create(&userDb).Error

	if err != nil{
		return users.DomainUser{},err
	}
	return userDb.ToDomain(),nil
}

func (repo *UserRepository) FindById(id int,ctx context.Context)(users.DomainUser,error){
	userDb := User{}

	err := repo.db.First(&userDb, id).Error
	if err != nil{
		return users.DomainUser{},err
	}
	
	return userDb.ToDomain(),nil
}