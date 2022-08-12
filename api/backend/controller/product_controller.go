package controller

import (
	"encoding/json"
	"ScreedCare/backend/helper"
	"ScreedCare/backend/service"
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

	byteData, _ := json.Marshal(&products)
	w.Header().Add("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(byteData)
}

func (c *productController) AddManyProduct(w http.ResponseWriter, r *http.Request){
	product := model.ProductData{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&product)
	helper.PanicIfError(err)

	c.service.AddManyProduct(product)

	byteData, _ := json.Marshal(&product)
	w.Header().Add("Content-type", "application/json")
	w.Write(byteData)
}

func (c *productController) AddProduct(w http.ResponseWriter, r *http.Request) {
	product := model.ProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&product)
	helper.PanicIfError(err)

	c.service.AddProduct(product)
	

	byteData, _ := json.Marshal(&product)
	w.Header().Add("Content-type", "application/json")
	w.Write(byteData)
}

func (c *productController) DeleteProduct(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	helper.PanicIfError(err)

	product := c.service.DeleteProduct(id)

	byteData, _ := json.Marshal(&product)
	w.Header().Add("Content-type", "application/json")
	w.Write(byteData)
}

func (c *productController) GetProduct(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	helper.PanicIfError(err)

	product := c.service.FindProduct(id)

	byteData, _ := json.Marshal(&product)
	w.Header().Add("Content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Write(byteData)
}

func (c *productController) UpdateProduct(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	helper.PanicIfError(err)

	product := model.ProductRequest{}
	err = json.NewDecoder(r.Body).Decode(&product)
	helper.PanicIfError(err)

	c.service.UpdateProduct(product, id)

	productResponse := c.service.FindProduct(id)
	byteData, _ := json.Marshal(&productResponse)
	w.Header().Add("Content-type", "application/json")
	w.Write(byteData)
}