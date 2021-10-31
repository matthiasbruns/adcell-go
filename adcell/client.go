// Package adcell
// In the adcell package you will find all required functions and structs to communicate with the adcell.com services.
package adcell

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Constants for url building
const (
	baseUrl = "https://www.adcell.de"

	/// Example https://www.adcell.de/promotion/csv?promoId=12345&slotId=67890
	dataFeedUrl = "%s/promotion/csv?promoId=%s&slotId=%s&stringMarker=%s&columnMarker=%s"
)

// DataFeedOptions
/// PromoId You can get the download API key from a standard feed download as given by Create-a-Feed. You can also get the full download link including the relevant API key to access this file from the Create-a-Feed section in the interface (Adcell interface --> Toolbox --> Create-a-Feed).
/// SlotId The string of the feed id that should be requested
type DataFeedOptions struct {
	PromoId string
	SlotId  string
}

// AdcellClient
/// Client that takes over the communication with the Adcell endpoints as well as parsing the response csv data into structs.
type AdcellClient struct {
	client *http.Client
}

func (c AdcellClient) FetchDataFeed(options *DataFeedOptions) (*[]DataFeedEntry, error) {
	url := fmt.Sprintf(dataFeedUrl, baseUrl, options.PromoId, options.SlotId, "%27", "%3B")

	return c.FetchDataFeedFromUrl(url)
}

func (c AdcellClient) FetchDataFeedFromUrl(url string) (*[]DataFeedEntry, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	plainResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(plainResponse))
	} else {
		return parseCSVToDataFeedEntry(bytes.NewReader(plainResponse))
	}
}

func parseCSVToDataFeedEntry(r io.Reader) (*[]DataFeedEntry, error) {
	reader := csv.NewReader(r)
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	var entries []DataFeedEntry
	columnNamesSkipped := false
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Skip column names from csv
		if !columnNamesSkipped {
			columnNamesSkipped = true
			continue
		}

		entry := DataFeedEntry{
			Deeplink:                   record[0],
			ProductTitle:               record[1],
			ProductDescription:         record[2],
			ProductDescriptionLong:     record[3],
			PriceGross:                 record[4],
			PriceNet:                   record[5],
			Currency:                   record[6],
			EAN:                        record[7],
			AAN:                        record[8],
			Manufacturer:               record[9],
			HAN:                        record[10],
			ProductImageUrl:            record[11],
			ProductPreviewImageUrl:     record[12],
			ProductCategory:            record[13],
			Shipping:                   record[14],
			ShippingPaymentInAdvance:   record[15],
			ShippingCashOnDelivery:     record[16],
			ShippingCreditCard:         record[17],
			ShippingDebitCharge:        record[18],
			ShippingInvoice:            record[19],
			ShippingPayPal:             record[20],
			ShippingSofortueberweisung: record[21],
			DeliveryTime:               record[22],
			BasePrice:                  record[23],
			UnitPricingBaseMeasure:     record[24],
			UnitPricingMeasure:         record[25],
			ExtSalePrice:               record[26],
			ExtCurrency:                record[27],
			ExtUnitPricingMeasure:      record[28],
			ExtUnitPricingBaseMeasure:  record[29],
			ExtBasePrice:               record[30],
			ExtBasePriceFormatted:      record[31],
			Ext_:                       record[32],
		}

		entries = append(entries, entry)
	}

	return &entries, nil
}

// NewAdcellClient
/// Returns a new AdcellClient. Needs a http.Client passed from outside.
/// client Required to be passed from the caller
/// returns a new instance of AdcellClient
func NewAdcellClient(client *http.Client) *AdcellClient {
	return &AdcellClient{client: client}
}
