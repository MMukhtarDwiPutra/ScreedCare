package exception

import(
	"net/http"
	"ScreedCare/model"
	"encoding/json"
	// "github.com/go-playground/validator/v10"
)

func ErrorHandler(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func(){
			if err := recover(); err != nil{
				if(notFoundError(w, r, err)){
					return
				}
				if(validationError(w, r, err)){
					return
				}
				internalServerError(w, r, err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool{
	data, exist := err.(NotFoundError)
	if exist{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := model.WebResponse{
			Code : http.StatusNotFound,
			Status : "DATA NOT FOUND",
			Data : data.Error,
		}

		byteData, _ := json.Marshal(&webResponse)
		w.Header().Add("Content-type", "application/json")
		w.Write(byteData)

		return true
	}else{
		return false
	}
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool{
	return true
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := model.WebResponse{
		Code : http.StatusInternalServerError,
		Status : "INTERNAL SERVER ERROR",
		Data : err,
	}

	byteData, _ := json.Marshal(&webResponse)
	w.Header().Add("Content-type", "application/json")
	w.Write(byteData)
}