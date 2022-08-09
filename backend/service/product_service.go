package service

import (
	"ScreedCare/model"
	"ScreedCare/backend/repository"
)

type ProductService interface{
	GetAllProduct() []model.ProductResponse
 	AddProduct(p model.ProductRequest)
 	DeleteProduct(id int) model.ProductResponse
 	FindProduct(id int) model.ProductResponse
 	UpdateProduct(p model.ProductRequest, id int)
}

type productService struct{
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *productService{
	return &productService{repository}
}

func (s *productService) GetAllProduct() []model.ProductResponse{
	return s.repository.GetAllProduct()
}

func (s *productService) AddProduct(p model.ProductRequest) {
	s.repository.AddProduct(p)
}

func (s *productService) DeleteProduct(id int) model.ProductResponse{
	product := s.repository.FindProduct(id)

	s.repository.DeleteProduct(id)
	return product
}

func (s *productService) FindProduct(id int) model.ProductResponse{
	return s.repository.FindProduct(id)
}

func (s *productService) UpdateProduct(p model.ProductRequest, id int){
	s.repository.UpdateProduct(p, id)
}