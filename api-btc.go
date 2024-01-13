package main

import (
	"fmt"
	"time"

	"github.com/JulianToledano/goingecko"
)

func main() {
	cgClient := goingecko.NewClient(nil)
	defer cgClient.Close()

	var prices []float64
	priceTicker := time.NewTicker(10 * time.Second)
	averageTicker := time.NewTicker(10 * time.Minute)

	for {
		select {
		case <-priceTicker.C:
			data, err := cgClient.CoinsId("bitcoin", true, true, true, false, false, false)
			if err != nil {
				fmt.Println("Error fetching Bitcoin price:", err)
				continue
			}
			prices = append(prices, data.MarketData.CurrentPrice.Usd)
			if len(prices) > 60 {
				prices = prices[1:] // Keep only the last 10 minutes of prices
			}

		case <-averageTicker.C:
			if len(prices) > 0 {
				averagePrice := calculateAverage(prices)
				fmt.Printf("Average Bitcoin price over the last 10 minutes: %f$\n", averagePrice)
				prices = []float64{} // Reset the prices slice for the next interval
			}
		}
	}
}

func calculateAverage(prices []float64) float64 {
	var total float64
	for _, price := range prices {
		total += price
	}
	return total / float64(len(prices))
}
