package controller

import (
	"encoding/json"
	"ScreedCare/backend/helper"
	"ScreedCare/backend/service"
	"ScreedCare/backend/middleware"
	"ScreedCare/model"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type productController struct{
	service service.ProductService
}

func NewProductController(service service.ProductService) *productController{
	return &productController{service}
}

func (c *productController) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products := c.service.GetAllProduct()

	webResponse := model.WebResponse{
		Code : 200,
		Status : "OK",
		Data : products,
	}
	byteData, _ := json.Marshal(&webResponse)
	w.Header().Add("Content-type", "application/json")
	w.Write(byteData)
}

func (c *productController) AddManyProduct(w http.ResponseWriter, r *http.Request){
	if(middleware.AuthMiddleware(w,r)){
		product := model.ProductData{}
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&product)
		helper.PanicIfError(err)

		c.service.AddManyProduct(product)

		webResponse := model.WebResponse{
			Code : 200,
			Status : "OK",
			Data : product,
		}
		byteData, _ := json.Marshal(&webResponse)
		w.Header().Add("Content-type", "application/json")
		w.Write(byteData)
	}
}

func (c *productController) AddProduct(w http.ResponseWriter, r *http.Request) {
	if(middleware.AuthMiddleware(w,r)){
		product := model.ProductRequest{}
		err := json.NewDecoder(r.Body).Decode(&product)
		helper.PanicIfError(err)

		c.service.AddProduct(product)
		
		webResponse := model.WebResponse{
			Code : 200,
			Status : "OK",
			Data : product,
		}
		byteData, _ := json.Marshal(&webResponse)
		w.Header().Add("Content-type", "application/json")
		w.Write(byteData)
	}
}

func (c *productController) DeleteProduct(w http.ResponseWriter, r *http.Request){
	if(middleware.AuthMiddleware(w,r)){
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		helper.PanicIfError(err)

		product := c.service.DeleteProduct(id)

		webResponse := model.WebResponse{
			Code : 200,
			Status : "OK",
			Data : product,
		}
		byteData, _ := json.Marshal(&webResponse)
		w.Header().Add("Content-type", "application/json")
		w.Write(byteData)
	}
}

func (c *productController) GetProduct(w http.ResponseWriter, r *http.Request){
	if(middleware.AuthMiddleware(w,r)){
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		helper.PanicIfError(err)

		product := c.service.FindProduct(id)

		webResponse := model.WebResponse{
			Code : 200,
			Status : "OK",
			Data : product,
		}
		byteData, _ := json.Marshal(&webResponse)
		w.Header().Add("Content-type", "application/json")
		w.Write(byteData)
	}
}

func (c *productController) UpdateProduct(w http.ResponseWriter, r *http.Request){
	if(middleware.AuthMiddleware(w,r)){
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		helper.PanicIfError(err)

		product := model.ProductRequest{}
		err = json.NewDecoder(r.Body).Decode(&product)
		helper.PanicIfError(err)

		c.service.UpdateProduct(product, id)

		productResponse := c.service.FindProduct(id)

		webResponse := model.WebResponse{
			Code : 200,
			Status : "OK",
			Data : productResponse,
		}
		byteData, _ := json.Marshal(&webResponse)
		w.Header().Add("Content-type", "application/json")
		w.Write(byteData)
	}
}