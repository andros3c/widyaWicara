package products

import (
	"time"
	"widyaWicaraBackend/businesses/products"

	"gorm.io/gorm"
)

type Product struct {
	Id        uint
	UserId    int
	Name      string
	Qty       int
	Desc      string
	CreatedAt time.Time      `gorm:"<-:create"`
	UpdatedAt time.Time      `gorm:"<-:update"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain products.DomainProduct)Product{
	return Product{
		Id: domain.Id,
		UserId: domain.UserId,
		Name: domain.Name,
		Qty: domain.Qty,
		Desc: domain.Desc,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,

	}
}

func (product Product)ToDomain()(products.DomainProduct){
	return products.DomainProduct{
		Id : product.Id,
		UserId: product.UserId,
		Name: product.Name,
		Qty: product.Qty,
		Desc: product.Desc,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		DeletedAt: product.DeletedAt,
	}
}

func ToDomainArray(product []Product) []products.DomainProduct {
	var response []products.DomainProduct

	for _, val := range product{
		response = append(response, val.ToDomain())
	}
	return response
}