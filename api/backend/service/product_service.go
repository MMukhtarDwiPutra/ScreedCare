package service

import (
	"ScreedCare/model"
	"ScreedCare/backend/repository"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"ScreedCare/backend/exception"
	"fmt"
)

type ProductService interface{
	GetAllProduct() []model.ProductResponse
 	AddProduct(p model.ProductRequest)
 	DeleteProduct(id int) model.ProductResponse
 	FindProduct(id int) model.ProductResponse
 	UpdateProduct(p model.ProductRequest, id int)
 	AddManyProduct(p model.ProductData)
}

type productService struct{
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *productService{
	return &productService{repository}
}

func (s *productService) GetAllProduct() []model.ProductResponse{
	products := s.repository.GetAllProduct()
	p := message.NewPrinter(language.English)
	for i := 0; i < len(products); i++{
		products[i].HargaStr = p.Sprintf("%d",products[i].Harga)
	}
	return products
}

func (s *productService) AddManyProduct(p model.ProductData){
	for i := 0; i < len(p.Data); i++{
		s.repository.AddProduct(p.Data[i])
	}
}

func (s *productService) AddProduct(p model.ProductRequest) {
	s.repository.AddProduct(p)
}

func (s *productService) DeleteProduct(id int) model.ProductResponse{
	product := s.FindProduct(id)

	s.repository.DeleteProduct(id)
	return product
}

func (s *productService) FindProduct(id int) model.ProductResponse{
	product := s.repository.FindProduct(id)

	if (product == model.ProductResponse{}){
		panic(exception.NewNotFoundError(fmt.Sprintf("Data product with id %d not found",id)))
	}

	p := message.NewPrinter(language.English)
	product.HargaStr = p.Sprintf("%d",product.Harga)
	return product
}

func (s *productService) UpdateProduct(p model.ProductRequest, id int){
	s.FindProduct(id)
	s.repository.UpdateProduct(p, id)
}