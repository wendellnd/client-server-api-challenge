package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wendellnd/client-server-api-challenge/server/db"
	"github.com/wendellnd/client-server-api-challenge/server/integration"
	serverutils "github.com/wendellnd/client-server-api-challenge/server/serverutils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", HandleGetExchange)

	port := ":8080"
	log.Println("running at port:", port)
	http.ListenAndServe(port, responseMiddleware(mux))
}

func responseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func HandleGetExchange(w http.ResponseWriter, r *http.Request) {
	exchange, err := integration.GetCurrentExchange()
	if err != nil {
		serverutils.HandleError(w, err)
		return
	}

	err = db.InsertExchange(r.Context(), exchange.USDBRL.Bid)
	if err != nil {
		serverutils.HandleError(w, err)
		return
	}

	result, err := json.Marshal(exchange)
	if err != nil {
		serverutils.HandleError(w, err)
		return
	}

	w.Write(result)
}
