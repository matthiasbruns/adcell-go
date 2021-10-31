package adcell_go

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/matthiasbruns/adcell-go/adcell"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type mockRoundTripper struct {
	response        *http.Response
	requestTestFunc func(r *http.Request) error
}

func (m mockRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if err := m.requestTestFunc(request); err != nil {
		return nil, err
	}
	return m.response, nil
}

func readCSVFileContents(filePath string) (string, error) {
	csvContent, err := ioutil.ReadFile(filePath) // just pass the file name
	if err != nil {
		return "", err
	}
	return string(csvContent), nil
}

func parseCSVToDataFeedEntry(csvContent string) (*[]adcell.DataFeedEntry, error) {
	reader := csv.NewReader(strings.NewReader(csvContent))
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1

	var entries []adcell.DataFeedEntry
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

		entry := adcell.DataFeedEntry{
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

func TestFetchDataFeed(t *testing.T) {
	// Read mock data from CSV
	csvContent, err := readCSVFileContents("testdata/data_feed.csv")
	if err != nil {
		t.Fatalf("coult not parse csv file '%v'", err)
	}

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(csvContent)); err != nil {
		t.Error(err)
	}
	if err := gz.Flush(); err != nil {
		t.Error(err)
	}
	if err := gz.Close(); err != nil {
		t.Error(err)
	}

	// Create mock response
	response := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(csvContent)),
	}

	// Create test client to run tests on
	adcellClient := adcell.NewAdcellClient(&http.Client{Transport: mockRoundTripper{response: response, requestTestFunc: func(r *http.Request) error {
		expectedUrl := "https://www.adcell.de/promotion/csv?promoId=promoId&slotId=slotId&stringMarker=%27&columnMarker=%3B"
		if r.URL.String() != expectedUrl {
			err := errors.New(fmt.Sprintf("invalid url found in test\nexpected '%s'\nfound '%s'", expectedUrl, r.URL.String()))
			t.Error(err)
			return err
		}

		expectedMethod := "GET"
		if r.Method != expectedMethod {
			err := errors.New(fmt.Sprintf("invalid request method in test\nexpected '%s'\nfound '%s'", expectedMethod, r.Method))
			t.Error(err)
			return err
		}

		return nil
	}}})

	result, err := adcellClient.FetchDataFeed(&adcell.DataFeedOptions{
		PromoId: "promoId",
		SlotId:  "slotId",
	})
	if err != nil {
		t.Fatalf("err is not null '%v'", err)
	}

	if len(*result) != 9 {
		t.Fatalf("Invalid amount of data rows received %d", len(*result))
	}

	// Check if received rows and expected rows match
	expectedRows, _ := parseCSVToDataFeedEntry(csvContent)
	for i, expectedRow := range *expectedRows {
		receivedRow := (*result)[i]
		if expectedRow != receivedRow {

			eJson, _ := json.Marshal(expectedRow)
			rJson, _ := json.Marshal(receivedRow)
			t.Fatalf("Invalid row parsed\nexpected '%v'\nreceived '%v'", string(eJson), string(rJson))
		}
	}
}

func TestFetchDataFeedFromUrl(t *testing.T) {
	// Read mock data from CSV
	csvContent, err := readCSVFileContents("testdata/data_feed.csv")
	if err != nil {
		t.Fatalf("coult not parse csv file '%v'", err)
	}

	// Create mock response
	response := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(csvContent)),
	}

	// Create test client to run tests on
	adcellClient := adcell.NewAdcellClient(&http.Client{Transport: mockRoundTripper{response: response, requestTestFunc: func(r *http.Request) error {
		fmt.Println(r.URL.String())
		expectedUrl := "https://www.adcell.de/promotion/csv?promoId=promoId&slotId=slotId&stringMarker=%27&columnMarker=%3B"
		if r.URL.String() != expectedUrl {
			err := errors.New(fmt.Sprintf("invalid url found in test\nexpected '%s'\nfound '%s'", expectedUrl, r.URL.String()))
			t.Error(err)
			return err
		}

		expectedMethod := "GET"
		if r.Method != expectedMethod {
			err := errors.New(fmt.Sprintf("invalid request method in test\nexpected '%s'\nfound '%s'", expectedMethod, r.Method))
			t.Error(err)
			return err
		}

		return nil
	}}})

	result, err := adcellClient.FetchDataFeedFromUrl("https://www.adcell.de/promotion/csv?promoId=promoId&slotId=slotId&stringMarker=%27&columnMarker=%3B")
	if err != nil {
		t.Fatalf("err is not null '%v'", err)
	}

	if len(*result) != 9 {
		t.Fatalf("Invalid amount of data rows received %d", len(*result))
	}

	// Check if received rows and expected rows match
	expectedRows, _ := parseCSVToDataFeedEntry(csvContent)
	for i, expectedRow := range *expectedRows {
		receivedRow := (*result)[i]
		if expectedRow != receivedRow {
			t.Fatalf("Invalid row parsed\nexpected '%v'\nreceived '%v'", expectedRow, receivedRow)
		}
	}
}
