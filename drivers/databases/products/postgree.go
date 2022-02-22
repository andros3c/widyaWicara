package products

import (
	"context"
	
	"widyaWicaraBackend/businesses/products"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(gormDb *gorm.DB)products.Repository{
	return &ProductRepository{
		db:gormDb,
	}
}

func (repo *ProductRepository)AddProduct(domain products.DomainProduct,ctx context.Context)(products.DomainProduct,error){
	prodDb := FromDomain(domain)

	err := repo.db.Create(&prodDb).Error

	if err != nil{
		return products.DomainProduct{},err
	}
	return prodDb.ToDomain(),nil

}

func (repo *ProductRepository)ShowAll(id int,ctx context.Context)([]products.DomainProduct, error){
	prodDb := []Product{}

	err := repo.db.Find(&prodDb,Product{UserId: id}).Error
	if err != nil{
		return []products.DomainProduct{},err
	}
	return ToDomainArray(prodDb),nil
	
}
func(repo *ProductRepository)FindById(id int,ctx context.Context)(products.DomainProduct,error){

	prodb := Product{}

err := repo.db.First(&prodb, id).Error
if err != nil{
	return products.DomainProduct{},err
}

return prodb.ToDomain(),nil
}

func(repo *ProductRepository) Update(id int,domain products.DomainProduct,ctx context.Context)(products.DomainProduct,error){
	product := Product{}
	prodDb := FromDomain(domain)
	
	err := repo.db.Model(&product).Where("id = ?",id).Updates(prodDb).Error

if err != nil{
	return products.DomainProduct{},err
}

return prodDb.ToDomain(),err

}

func (repo *ProductRepository) Delete(id int,ctx context.Context)(products.DomainProduct,error){
	product := Product{}

	err := repo.db.Delete(&product, id).Error

	if err != nil{
		return products.DomainProduct{},err
	}
	
	return product.ToDomain(),err
	
}