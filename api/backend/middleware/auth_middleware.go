package middleware

import(
	"net/http"
	"ScreedCare/model"
	"encoding/json"
)

func AuthMiddleware(w http.ResponseWriter, r *http.Request) bool{
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		// ok
		return true
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		byteData, _ := json.Marshal(&webResponse)
		w.Header().Add("Content-type", "application/json")
		w.Write(byteData)
		return false
	}
}