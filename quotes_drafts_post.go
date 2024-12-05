package economic

import (
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/jojokbh/go-economic/omitempty"
)

func (c *Client) NewQuotesDraftsPostRequest() QuotesDraftsPostRequest {
	return QuotesDraftsPostRequest{
		client:      c,
		queryParams: c.NewQuotesDraftsPostQueryParams(),
		pathParams:  c.NewQuotesDraftsPostPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewQuotesDraftsPostRequestBody(),
	}
}

type QuotesDraftsPostRequest struct {
	client      *Client
	queryParams *QuotesDraftsPostQueryParams
	pathParams  *QuotesDraftsPostPathParams
	method      string
	headers     http.Header
	requestBody QuotesDraftsPostRequestBody
}

func (c *Client) NewQuotesDraftsPostQueryParams() *QuotesDraftsPostQueryParams {
	return &QuotesDraftsPostQueryParams{}
}

type QuotesDraftsPostQueryParams struct {
}

func (r QuotesDraftsPostRequest) RequiredProperties() []string {
	return []string{
		"currency",
		"customer",
		"date",
		"layout",
		"paymentTerms",
		"recipient",
		"recipient.name",
		"recipient.vatZone",
	}
}

func (r QuotesDraftsPostRequest) FilterableProperties() []string {
	return []string{}
}

func (r QuotesDraftsPostRequest) SortableProperties() []string {
	return []string{}
}

func (p QuotesDraftsPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *QuotesDraftsPostRequest) QueryParams() *QuotesDraftsPostQueryParams {
	return r.queryParams
}

func (c *Client) NewQuotesDraftsPostPathParams() *QuotesDraftsPostPathParams {
	return &QuotesDraftsPostPathParams{}
}

type QuotesDraftsPostPathParams struct {
}

func (p *QuotesDraftsPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *QuotesDraftsPostRequest) PathParams() *QuotesDraftsPostPathParams {
	return r.pathParams
}

func (r *QuotesDraftsPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *QuotesDraftsPostRequest) Method() string {
	return r.method
}

func (s *Client) NewQuotesDraftsPostRequestBody() QuotesDraftsPostRequestBody {
	return QuotesDraftsPostRequestBody{}
}

type QuotesDraftsPostRequestBodyCustomer struct {
	CustomerNumber int `json:"customerNumber"`
}

type QuotesDraftsPostRequestBodyReference struct {
	Other string `json:"other"`
}

type QuotesDraftsPostRequestBody struct {
	Date                    Date                                 `json:"date"`
	Currency                string                               `json:"currency"`
	ExchangeRate            float64                              `json:"exchangeRate,omitempty"`
	NetAmount               float64                              `json:"netAmount"`
	NetAmountInBaseCurrency float64                              `json:"netAmountInBaseCurrency"`
	GrossAmount             float64                              `json:"grossAmount"`
	MarginInBaseCurrency    float64                              `json:"marginInBaseCurrency"`
	MarginPercentage        float64                              `json:"marginPercentage"`
	VatAmount               float64                              `json:"vatAmount"`
	RoundingAmount          float64                              `json:"roundingAmount"`
	CostPriceInBaseCurrency float64                              `json:"costPriceInBaseCurrency"`
	Notes                   OrderNotes                           `json:"notes,omitempty"`
	PaymentTerms            QuoteDraftPaymentTerms               `json:"paymentTerms,omitempty"`
	Customer                QuotesDraftsPostRequestBodyCustomer  `json:"customer"`
	Recipient               QuoteDraftRecipient                  `json:"recipient,omitempty"`
	Delivery                QuoteDraftDelivery                   `json:"delivery,omitempty"`
	References              QuotesDraftsPostRequestBodyReference `json:"references"`
	Layout                  QuoteDraftLayout                     `json:"layout,omitempty"`
	Lines                   []QuoteDraftLine                     `json:"lines,omitempty"`
}

func (r *QuotesDraftsPostRequest) RequestBody() *QuotesDraftsPostRequestBody {
	return &r.requestBody
}

func (r *QuotesDraftsPostRequest) SetRequestBody(body QuotesDraftsPostRequestBody) {
	r.requestBody = body
}

func (r *QuotesDraftsPostRequest) NewResponseBody() *QuotesDraftsPostResponseBody {
	return &QuotesDraftsPostResponseBody{}
}

type QuotesDraftsPostResponseBody struct {
	DraftQuoteNumber int `json:"draftQuoteNumber"`
	Soap             struct {
		CurrentQuoteHandle struct {
			ID int `json:"id"`
		} `json:"currentQuoteHandle"`
	} `json:"soap"`
	Templates struct {
		BookingInstructions string `json:"bookingInstructions"`
		Self                string `json:"self"`
	} `json:"templates"`
	Date                    string  `json:"date"`
	Currency                string  `json:"currency"`
	ExchangeRate            float64 `json:"exchangeRate"`
	NetAmount               float64 `json:"netAmount"`
	NetAmountInBaseCurrency float64 `json:"netAmountInBaseCurrency"`
	GrossAmount             float64 `json:"grossAmount"`
	MarginInBaseCurrency    float64 `json:"marginInBaseCurrency"`
	MarginPercentage        float64 `json:"marginPercentage"`
	VatAmount               float64 `json:"vatAmount"`
	RoundingAmount          float64 `json:"roundingAmount"`
	CostPriceInBaseCurrency float64 `json:"costPriceInBaseCurrency"`
	DueDate                 string  `json:"dueDate"`
	PaymentTerms            struct {
		PaymentTermsNumber int    `json:"paymentTermsNumber"`
		Name               string `json:"name"`
		PaymentTermsType   string `json:"paymentTermsType"`
		Self               string `json:"self"`
	} `json:"paymentTerms"`
	Customer struct {
		CustomerNumber int    `json:"customerNumber"`
		Self           string `json:"self"`
	} `json:"customer"`
	Recipient struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		VatZone struct {
			Name               string `json:"name"`
			VatZoneNumber      int    `json:"vatZoneNumber"`
			EnabledForCustomer bool   `json:"enabledForCustomer"`
			EnabledForSupplier bool   `json:"enabledForSupplier"`
			Self               string `json:"self"`
		} `json:"vatZone"`
	} `json:"recipient"`
	Layout struct {
		LayoutNumber int    `json:"layoutNumber"`
		Self         string `json:"self"`
	} `json:"layout"`
	Pdf struct {
		Download string `json:"download"`
	} `json:"pdf"`
	Self string `json:"self"`
}

func (r *QuotesDraftsPostRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("quotes/drafts", r.PathParams())
}

func (r *QuotesDraftsPostRequest) Do() (QuotesDraftsPostResponseBody, error) {
	u, err := r.URL()
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), u, r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (i QuotesDraftsPostRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

type QuoteDraftLineUnit struct {
	UnitNumber int    `json:"unitNumber"`
	Name       string `json:"name"`
}

type QuoteDraftLineProduct struct {
	ProductNumber string `json:"productNumber"`
}

type QuoteDraftLine struct {
	LineNumber           int                   `json:"lineNumber"`
	SortKey              int                   `json:"sortKey"`
	Description          string                `json:"description"`
	Unit                 QuoteDraftLineUnit    `json:"unit,omitempty"`
	Product              QuoteDraftLineProduct `json:"product"`
	Quantity             float64               `json:"quantity"`
	UnitNetPrice         float64               `json:"unitNetPrice"`
	DiscountPercentage   float64               `json:"discountPercentage"`
	UnitCostPrice        float64               `json:"unitCostPrice"`
	TotalNetAmount       float64               `json:"totalNetAmount"`
	MarginInBaseCurrency float64               `json:"marginInBaseCurrency"`
	MarginPercentage     float64               `json:"marginPercentage"`
}

type QuoteDraftDelivery struct {
	Address      string `json:"address"`
	Zip          string `json:"zip"`
	City         string `json:"city"`
	Country      string `json:"country"`
	DeliveryDate string `json:"deliveryDate"`
}

func (d QuoteDraftDelivery) IsEmpty() bool {
	return zero.IsZero(d)
}

type QuoteDraftRecipient struct {
	Name    string            `json:"name"`
	Address string            `json:"address"`
	Zip     string            `json:"zip"`
	City    string            `json:"city"`
	VatZone QuoteDraftVatZone `json:"vatZone,omitempty"`
}

func (r QuoteDraftRecipient) IsEmpty() bool {
	return zero.IsZero(r)
}

func (r QuoteDraftRecipient) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type QuoteDraftPaymentTerms struct {
	PaymentTermsNumber int    `json:"paymentTermsNumber"`
	DaysOfCredit       int    `json:"daysOfCredit,omitempty"`
	Name               string `json:"name,omitempty"`
	PaymentTermsType   string `json:"paymentTermsType,omitempty"`
}

func (pt QuoteDraftPaymentTerms) IsEmpty() bool {
	return zero.IsZero(pt)
}

type QuoteDraftVatZone struct {
	Name               string `json:"name"`
	VatZoneNumber      int    `json:"vatZoneNumber,omitempty"`
	EnabledForCustomer bool   `json:"enabledForCustomer"`
	EnabledForSupplier bool   `json:"enabledForSupplier"`
}

func (vz QuoteDraftVatZone) IsEmpty() bool {
	return zero.IsZero(vz)
}

type QuoteDraftLayout struct {
	LayoutNumber int `json:"layoutNumber,omitempty"`
}

func (l QuoteDraftLayout) IsEmpty() bool {
	return zero.IsZero(l)
}
