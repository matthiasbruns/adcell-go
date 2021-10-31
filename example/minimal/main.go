package main

import (
	"fmt"
	"github.com/matthiasbruns/adcell-go/adcell"
	"net/http"
)

func main() {
	adcellClient := adcell.NewAdcellClient(&http.Client{})

	fetchDataFeed(adcellClient)
}

func fetchDataFeed(adcellClient *adcell.AdcellClient) {
	feed, err := adcellClient.FetchDataFeed(&adcell.DataFeedOptions{
		PromoId: "promoId",
		SlotId:  "slotId",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(feed)
}
