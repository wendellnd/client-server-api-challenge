package serverutils

import (
	"log"
	"net/http"
	"os"
)

func InternalServerError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"message": "Internal Server Error"}`))
}

func TimeoutError(w http.ResponseWriter) {
	log.Println("Timeout Error")
	w.WriteHeader(http.StatusRequestTimeout)
	w.Write([]byte(`{"message": "Request Timeout"}`))
}

func HandleError(w http.ResponseWriter, err error) {
	if os.IsTimeout(err) {
		TimeoutError(w)
		return
	}

	InternalServerError(w, err)
}
