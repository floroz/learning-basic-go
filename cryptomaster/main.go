package main

import (
	"fmt"
	"sync"

	"danieletortora.com/cryptomaster/api"
)

func main() {
	currencies := [2]string{"USD", "EUR"}
	var prices [2]float64
	var wg sync.WaitGroup

	for i, c := range currencies {
		wg.Add(1)

		go func() {
			defer wg.Done()

			data, err := api.GetRate("BTC", c)

			if err != nil {
				fmt.Printf("Error while getting USD \n")
			}

			prices[i] = data.Price
		}()
	}

	wg.Wait()

	for i, c := range currencies {
		fmt.Printf("Price %s: %v\n", c, prices[i])
	}

}
