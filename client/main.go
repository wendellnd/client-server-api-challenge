package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Exchange struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	client := http.Client{
		Timeout: 300 * time.Millisecond,
	}

	request, err := http.NewRequest("GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		if os.IsTimeout(err) {
			log.Println("Request Timeout")
			return
		}

		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var exchange Exchange
	err = json.NewDecoder(response.Body).Decode(&exchange)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("cotação.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write([]byte(fmt.Sprintf("Dólar: %s", exchange.USDBRL.Bid)))
	if err != nil {
		panic(err)
	}
}
