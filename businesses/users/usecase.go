package users

import (
	"context"
	"time"
	_middleware "widyaWicaraBackend/app/middleware"
	"widyaWicaraBackend/businesses"
	encrypt "widyaWicaraBackend/drivers/helpers"

	"gorm.io/gorm"
)

type userUseCase struct {
	repo Repository
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT
}

func NewUserUsecase(UserRepo Repository,contextTimeout time.Duration,configJWT *_middleware.ConfigJWT) Usecase{
	return &userUseCase{
		repo : UserRepo,
		ctx: contextTimeout,
		jwt: configJWT,
	}
}



func (usecase *userUseCase)Login(domain DomainUser,ctx context.Context)(DomainUser,error){
	if domain.Email == ""{
		return DomainUser{},businesses.ErrEmailEmpty
	}
	if domain.Password == ""{
		return DomainUser{},businesses.ErrPassEmpty
	}

	user,err := usecase.repo.Login(domain,ctx)

	if err != nil{
		return DomainUser{},err
	}
	if !encrypt.ValidateHash(domain.Password,user.Password){
		return  DomainUser{},businesses.ErrWrongPass
	}
	user.Token = usecase.jwt.GenererateToken(user.Id)
	return user,nil

}
func (usecase *userUseCase) CreateNewUser(domain DomainUser,ctx context.Context)(DomainUser,error){
	if domain.Email == "" {
		return DomainUser{},businesses.ErrEmailEmpty
	}
	if domain.Username == ""{
		return DomainUser{},businesses.ErrUsernameEmpty
	}
	if domain.Password==""{
		return DomainUser{},businesses.ErrPasswordEmpty
	}

	existedUsername,err := usecase.repo.FindUsername(domain,ctx)
	if err == gorm.ErrRecordNotFound{
		existedEmail,err := usecase.repo.FindEmail(domain,ctx)
		if err == gorm.ErrRecordNotFound{
			domain.Password,err = encrypt.Hash(domain.Password)
			if err != nil{
				return DomainUser{},businesses.ErrInternalServer
			}
			user,err := usecase.repo.CreateNewUser(domain,ctx)
			if err != nil{
				return DomainUser{},err
			}
			return user,nil
		}else{
			return existedEmail,businesses.ErrEmailExisted
		}

	}else{
		return existedUsername,businesses.ErrUsernameExisted
	}
	
}

func (usecase userUseCase) FindById(id int,ctx context.Context)(DomainUser,error){
	user ,err := usecase.repo.FindById(id,ctx)
	if err != nil {
		return DomainUser{}, err
	}
	return user,nil

}

