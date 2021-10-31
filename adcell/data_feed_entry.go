package adcell

type DataFeedEntry struct {
	Deeplink                   string `json:"deeplink,omitempty"`
	ProductTitle               string `json:"product_title,omitempty"`
	ProductDescription         string `json:"product_description,omitempty"`
	ProductDescriptionLong     string `json:"product_description_long,omitempty"`
	PriceGross                 string `json:"price_gross,omitempty"`
	PriceNet                   string `json:"price_net,omitempty"`
	Currency                   string `json:"currency,omitempty"`
	EAN                        string `json:"ean,omitempty"`
	AAN                        string `json:"aan,omitempty"`
	Manufacturer               string `json:"manufacturer,omitempty"`
	HAN                        string `json:"han,omitempty"`
	ProductImageUrl            string `json:"product_image_url,omitempty"`
	ProductPreviewImageUrl     string `json:"product_preview_image_url,omitempty"`
	ProductCategory            string `json:"product_category,omitempty"`
	Shipping                   string `json:"shipping,omitempty"`
	ShippingPaymentInAdvance   string `json:"shipping_payment_in_advance,omitempty"`
	ShippingCashOnDelivery     string `json:"shipping_cash_on_delivery,omitempty" csv:"Versandkosten Nachnahme"`
	ShippingCreditCard         string `json:"shipping_credit_card,omitempty"`
	ShippingDebitCharge        string `json:"shipping_debit_charge,omitempty"`
	ShippingInvoice            string `json:"shipping_invoice,omitempty"`
	ShippingPayPal             string `json:"shipping_paypal,omitempty"`
	ShippingSofortueberweisung string `json:"shipping_sofortueberweisung,omitempty"`
	DeliveryTime               string `json:"delivery_time,omitempty"`
	BasePrice                  string `json:"base_price,omitempty"`
	UnitPricingBaseMeasure     string `json:"unit_pricing_base_measure,omitempty"`
	UnitPricingMeasure         string `json:"unit_pricing_measure,omitempty"`
	ExtSalePrice               string `json:"ext_sale_price,omitempty"`
	ExtCurrency                string `json:"ext_currency,omitempty"`
	ExtUnitPricingMeasure      string `json:"ext_unit_pricing_measure,omitempty"`
	ExtUnitPricingBaseMeasure  string `json:"ext_unit_pricing_base_measure,omitempty"`
	ExtBasePrice               string `json:"ext_base_price,omitempty"`
	ExtBasePriceFormatted      string `json:"ext_base_price_complete,omitempty"`
	Ext_                       string `json:"ext_,omitempty"`
}
