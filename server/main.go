package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wendellnd/client-server-api-challenge/server/db"
	"github.com/wendellnd/client-server-api-challenge/server/integration"
	serverutils "github.com/wendellnd/client-server-api-challenge/server/serverutils"
)

/*
Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.
*/

func main() {

	// DONE
	// O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL
	// e em seguida deverá retornar no formato JSON o resultado para o cliente.
	// O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.
	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", HandleGetExchange)

	port := ":8080"
	log.Println("running at port:", port)
	http.ListenAndServe(port, responseMiddleware(mux))

	// Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida,
	// sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms
	// e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.

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
