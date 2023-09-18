package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PauloRVF/desafio_client_server_api/server/dto"
	"github.com/PauloRVF/desafio_client_server_api/server/entity"
)

func main() {
	http.HandleFunc("/cotacao", handleCotacao)
	http.ListenAndServe(":8080", nil)
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	economiaApi, err := getEconomiaApi()
	if err != nil {
		log.Fatal(err)
	}

	exchange := entity.NewExchange(economiaApi)
	err = entity.PersistExchange(exchange)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(exchange)
}

func getEconomiaApi() (*dto.EconomiaApi, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var economiaApiResponse dto.EconomiaApi
	err = json.NewDecoder(res.Body).Decode(&economiaApiResponse)
	if err != nil {
		return nil, err
	}

	return &economiaApiResponse, nil
}
