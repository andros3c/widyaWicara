package request

import "widyaWicaraBackend/businesses/products"

type Product struct {
	UserId int   `json:"user_id"`
	Name   string `json:"product_name"`
	Qty    int    `json:"product_qty"`
	Desc   string `json:"product_desc"`
}

func (product *Product) ToDomain() *products.DomainProduct{
	return &products.DomainProduct{
		UserId: product.UserId,
		Name: product.Name,
		Qty: product.Qty,
		Desc: product.Desc,
	}
}