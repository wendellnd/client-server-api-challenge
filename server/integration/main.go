package integration

import (
	"encoding/json"
	"net/http"
	"time"
)

type Exchange struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

const timeout = 200 * time.Millisecond

func GetCurrentExchange() (exchange Exchange, err error) {
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return exchange, err
	}

	response, err := client.Do(request)
	if err != nil {
		return exchange, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&exchange)
	if err != nil {
		return exchange, err
	}

	return exchange, err
}
