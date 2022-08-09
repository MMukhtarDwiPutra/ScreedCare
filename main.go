package main

import(
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"ScreedCare/backend/app"
)

func main(){
	router := mux.NewRouter()

	app.AddRouter(router)

	log.Println("Server serve at port : 8080")
	http.ListenAndServe(":8080", router)
}