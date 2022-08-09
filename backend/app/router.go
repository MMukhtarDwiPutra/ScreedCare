package app

import(
	"github.com/gorilla/mux"
	"ScreedCare/backend/controller"
	"ScreedCare/backend/service"
	"ScreedCare/backend/repository"
	"ScreedCare/backend/handler"
	"path/filepath"
	"net/http"
)

func AddRouter(router *mux.Router){
	db := NewDB()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	templateHandler := handler.NewTemplateHandler(productService)

	router.HandleFunc("/api/product/", productController.GetAllProduct).Methods("GET")
	router.HandleFunc("/api/product/", productController.AddProduct).Methods("POST")
	router.HandleFunc("/api/product/{id}", productController.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/product/{id}", productController.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/product/{id}", productController.GetProduct).Methods("GET")

	router.HandleFunc("/", templateHandler.HomeHandler)

	router.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css/", http.FileServer(http.Dir(filepath.Join("assets", "css")))))
	router.PathPrefix("/static/img/").Handler(http.StripPrefix("/static/img/", http.FileServer(http.Dir(filepath.Join("assets", "img")))))
	router.PathPrefix("/static/js/").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir(filepath.Join("assets", "js")))))
}