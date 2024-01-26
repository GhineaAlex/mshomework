package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type CoinGeckoResponse struct {
	Bitcoin struct {
		Usd float64 `json:"usd"`
	} `json:"bitcoin"`
}

var (
	mu            sync.RWMutex
	currentPrice  float64
	prices        []float64
)

func fetchBitcoinPrice() (float64, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result CoinGeckoResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}

	return result.Bitcoin.Usd, nil
}

func updateBitcoinPrice() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		price, err := fetchBitcoinPrice()
		if err != nil {
			fmt.Println("Error fetching: ", err)
			continue
		}

		mu.Lock()
		currentPrice = price
		prices = append(prices, price)
		if len(prices) > 60 { 
			prices = prices[1:]
		}
		mu.Unlock()
	}
}

func averageBitcoinPrice() float64 {
	mu.RLock()
	defer mu.RUnlock()

	sum := 0.0
	for _, price := range prices {
		sum += price
	}
	if len(prices) == 0 {
		return 0
	}
	return sum / float64(len(prices))
}

func bitcoinPriceHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	current := currentPrice
	average := averageBitcoinPrice()
	mu.RUnlock()

	fmt.Fprintf(w, "Current Bitcoin Price: USD %f\nAverage (10 min): USD %f", current, average)
}

func main() {
	go updateBitcoinPrice()

	http.HandleFunc("/", bitcoinPriceHandler)

	fmt.Println("Starting server")
	err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
