package handler

import(
	"net/http"
	"html/template"
	"path/filepath"
	"ScreedCare/backend/helper"
	"ScreedCare/backend/service"
)

type TemplateHandler interface{
	HomeHandler(w http.ResponseWriter, r *http.Request)	
}

type templateHandler struct{
	service service.ProductService
}

func NewTemplateHandler(service service.ProductService) *templateHandler{
	return &templateHandler{service}
}

func (h *templateHandler) HomeHandler(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles(filepath.Join("views","home.html"))
	helper.PanicIfError(err)

	product := h.service.GetAllProduct()

	
	data := map[string]interface{}{
		"Product" : product,
	}
	err = t.Execute(w, data)
	helper.PanicIfError(err)
}