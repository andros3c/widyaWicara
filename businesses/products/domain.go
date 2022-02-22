package products

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type DomainProduct struct {
	Id        uint
	UserId	  int
	Name      string
	Qty       int
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Usecase interface{
	AddProduct(domain DomainProduct,ctx context.Context)(DomainProduct,error)
	ShowAll(id int,ctx context.Context)([]DomainProduct, error)
	FindById(id int,ctx context.Context)(DomainProduct,error)
	Update(id int,domain DomainProduct,ctx context.Context)(DomainProduct,error)
	Delete(id int,ctx context.Context)(DomainProduct,error)
}
type Repository interface{
	AddProduct(domain DomainProduct,ctx context.Context)(DomainProduct,error)
	ShowAll(id int,ctx context.Context)([]DomainProduct, error)
	FindById(id int,ctx context.Context)(DomainProduct,error)
	Update(id int,domain DomainProduct,ctx context.Context)(DomainProduct,error)
	Delete(id int,ctx context.Context)(DomainProduct,error)
}