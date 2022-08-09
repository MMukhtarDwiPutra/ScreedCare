package app

import(
	"github.com/gorilla/mux"
	"ScreedCare/backend/controller"
	"ScreedCare/backend/service"
	"ScreedCare/backend/repository"
)

func AddRouter(router *mux.Router){
	db := NewDB()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	router.HandleFunc("/api/product/", productController.GetAllProduct).Methods("GET")
	router.HandleFunc("/api/product/", productController.AddProduct).Methods("POST")
	router.HandleFunc("/api/product/{id}", productController.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/product/{id}", productController.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/product/{id}", productController.GetProduct).Methods("GET")
}