package response

import (
	"time"
	"widyaWicaraBackend/businesses/products"

	"gorm.io/gorm"
)

type ProductResponse struct {
	Id 		  uint			`json:"id"`
	UserId    int           `json:"user_id"`
	Name      string         `json:"product_name"`
	Qty       int            `json:"product_qty"`
	Desc      string         `json:"product_desc"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain products.DomainProduct)ProductResponse{
	return ProductResponse{
		Id: domain.Id,
		UserId: domain.UserId,
		Name: domain.Name,
		Qty: domain.Qty,
		Desc: domain.Desc,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: gorm.DeletedAt{},
	}
}