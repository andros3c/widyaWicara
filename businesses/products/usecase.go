package products

import (
	"context"
	"time"
	"widyaWicaraBackend/businesses"

)

type productUseCase struct {
	repo Repository
	ctx time.Duration
}

func NewProductUsecase(ProductRepo Repository, contextTimeout time.Duration) Usecase{
	return &productUseCase{
		repo : ProductRepo,
		ctx : contextTimeout,
	}

}

func (usecase *productUseCase) AddProduct(domain DomainProduct,ctx context.Context)(DomainProduct,error){
	if domain.Name == ""{
		return DomainProduct{},businesses.ErrProdNameEmpty
	}
	if domain.Qty <= 0{
		return DomainProduct{},businesses.ErrQtyProduct
	}

	product,err := usecase.repo.AddProduct(domain,ctx)
	if err != nil{
		return DomainProduct{},err
	}
	return product,nil
}

func (usecase *productUseCase) ShowAll(id int,ctx context.Context)([]DomainProduct, error){
	product, err := usecase.repo.ShowAll(id ,ctx)
	if err != nil {
		return []DomainProduct{}, err
	}
	return product,nil
}

func (usecase *productUseCase) FindById(id int,ctx context.Context)(DomainProduct,error){
	product,err := usecase.repo.FindById(id,ctx)
	if err != nil {
		return DomainProduct{}, err
	}
	return product,nil
}

func (usecase *productUseCase) Update(id int,domain DomainProduct,ctx context.Context)(DomainProduct,error){
	product,err := usecase.repo.Update(id,domain,ctx)
	if err != nil{
		return DomainProduct{},err
	}
	return product,nil
}

func (usecase *productUseCase) Delete(id int,ctx context.Context)(DomainProduct,error){
	product,err := usecase.repo.Delete(id,ctx)
	if err != nil {
		return DomainProduct{}, err
	}
	return product,nil
}