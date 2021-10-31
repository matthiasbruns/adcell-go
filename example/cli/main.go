package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/matthiasbruns/adcell-go/adcell"
	"net/http"
	"os"
)

const cliUsage = "expected 'feed' subcommands"
const feedUsage = "./adcell-go feed -promoId PROMO_ID -slotId SLOT_ID"

func main() {
	adcellClient := adcell.NewAdcellClient(&http.Client{})

	feedCmd := flag.NewFlagSet("feed", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println(cliUsage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "feed":
		handleFeedCmd(feedCmd, adcellClient)
	default:
		fmt.Println(cliUsage)
		os.Exit(1)
	}
}

func handleFeedCmd(feedListCmd *flag.FlagSet, adcellClient *adcell.AdcellClient) {
	promoId := feedListCmd.String("promoId", "", "-promoId PROMO_ID")
	slotId := feedListCmd.String("slotId", "", "-slotId SLOT_ID")

	if err := feedListCmd.Parse(os.Args[2:]); err != nil {
		fmt.Print(feedUsage)
		os.Exit(1)
	}

	fmt.Println("loading datafeed from Adcell")

	results, err := adcellClient.FetchDataFeed(&adcell.DataFeedOptions{
		PromoId: *promoId,
		SlotId:  *slotId,
	})

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	if j, err := json.Marshal(results); err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		fmt.Print(string(j))
	}
}
