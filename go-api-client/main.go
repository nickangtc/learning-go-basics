package api

import "sync"

func getCurrencyData(currencyCode string) {
	// do some http get request
}

func main() {
	currencies := []string{"BTC", "BCH", "ETH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}
