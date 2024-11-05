package economic

import (
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/jojokbh/go-economic/omitempty"
)

func (c *Client) NewOrdersDraftsPostRequest() OrdersDraftsPostRequest {
	return OrdersDraftsPostRequest{
		client:      c,
		queryParams: c.NewOrdersDraftsPostQueryParams(),
		pathParams:  c.NewOrdersDraftsPostPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewOrdersDraftsPostRequestBody(),
	}
}

type OrdersDraftsPostRequest struct {
	client      *Client
	queryParams *OrdersDraftsPostQueryParams
	pathParams  *OrdersDraftsPostPathParams
	method      string
	headers     http.Header
	requestBody OrdersDraftsPostRequestBody
}

func (c *Client) NewOrdersDraftsPostQueryParams() *OrdersDraftsPostQueryParams {
	return &OrdersDraftsPostQueryParams{}
}

type OrdersDraftsPostQueryParams struct {
}

func (r OrdersDraftsPostRequest) RequiredProperties() []string {
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

func (r OrdersDraftsPostRequest) FilterableProperties() []string {
	return []string{}
}

func (r OrdersDraftsPostRequest) SortableProperties() []string {
	return []string{}
}

func (p OrdersDraftsPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *OrdersDraftsPostRequest) QueryParams() *OrdersDraftsPostQueryParams {
	return r.queryParams
}

func (c *Client) NewOrdersDraftsPostPathParams() *OrdersDraftsPostPathParams {
	return &OrdersDraftsPostPathParams{}
}

type OrdersDraftsPostPathParams struct {
}

func (p *OrdersDraftsPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OrdersDraftsPostRequest) PathParams() *OrdersDraftsPostPathParams {
	return r.pathParams
}

func (r *OrdersDraftsPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *OrdersDraftsPostRequest) Method() string {
	return r.method
}

func (s *Client) NewOrdersDraftsPostRequestBody() OrdersDraftsPostRequestBody {
	return OrdersDraftsPostRequestBody{}
}

type OrdersDraftsPostRequestBodyCustomer struct {
	CustomerNumber int `json:"customerNumber"`
}

type OrdersDraftsPostRequestBodyReference struct {
	Other string `json:"other"`
}

type OrdersDraftsPostRequestBody struct {
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
	PaymentTerms            OrderDraftPaymentTerms               `json:"paymentTerms,omitempty"`
	Customer                OrdersDraftsPostRequestBodyCustomer  `json:"customer"`
	Recipient               OrderDraftRecipient                  `json:"recipient,omitempty"`
	Delivery                OrderDraftDelivery                   `json:"delivery,omitempty"`
	References              OrdersDraftsPostRequestBodyReference `json:"references"`
	Layout                  OrderDraftLayout                     `json:"layout,omitempty"`
	Lines                   []OrderDraftLine                     `json:"lines,omitempty"`
}

func (r *OrdersDraftsPostRequest) RequestBody() *OrdersDraftsPostRequestBody {
	return &r.requestBody
}

func (r *OrdersDraftsPostRequest) SetRequestBody(body OrdersDraftsPostRequestBody) {
	r.requestBody = body
}

func (r *OrdersDraftsPostRequest) NewResponseBody() *OrdersDraftsPostResponseBody {
	return &OrdersDraftsPostResponseBody{}
}

type OrdersDraftsPostResponseBody struct {
	DraftOrderNumber int `json:"draftOrderNumber"`
	Soap             struct {
		CurrentOrderHandle struct {
			ID int `json:"id"`
		} `json:"currentOrderHandle"`
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

func (r *OrdersDraftsPostRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("orders/drafts", r.PathParams())
}

func (r *OrdersDraftsPostRequest) Do() (OrdersDraftsPostResponseBody, error) {
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

func (i OrdersDraftsPostRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

type OrderDraftLineUnit struct {
	UnitNumber int    `json:"unitNumber"`
	Name       string `json:"name"`
}

type OrderDraftLineProduct struct {
	ProductNumber string `json:"productNumber"`
}

type OrderDraftLine struct {
	LineNumber           int                   `json:"lineNumber"`
	SortKey              int                   `json:"sortKey"`
	Unit                 OrderDraftLineUnit    `json:"unit,omitempty"`
	Product              OrderDraftLineProduct `json:"product"`
	Quantity             float64               `json:"quantity"`
	UnitNetPrice         float64               `json:"unitNetPrice"`
	DiscountPercentage   float64               `json:"discountPercentage"`
	UnitCostPrice        float64               `json:"unitCostPrice"`
	TotalNetAmount       float64               `json:"totalNetAmount"`
	MarginInBaseCurrency float64               `json:"marginInBaseCurrency"`
	MarginPercentage     float64               `json:"marginPercentage"`
	Description          string                `json:"description"`
}

type OrderDraftDelivery struct {
	Address      string `json:"address"`
	Zip          string `json:"zip"`
	City         string `json:"city"`
	Country      string `json:"country"`
	DeliveryDate string `json:"deliveryDate"`
}

func (d OrderDraftDelivery) IsEmpty() bool {
	return zero.IsZero(d)
}

type OrderDraftRecipient struct {
	Name    string            `json:"name"`
	Address string            `json:"address"`
	Zip     string            `json:"zip"`
	City    string            `json:"city"`
	VatZone OrderDraftVatZone `json:"vatZone,omitempty"`
}

func (r OrderDraftRecipient) IsEmpty() bool {
	return zero.IsZero(r)
}

func (r OrderDraftRecipient) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type OrderDraftPaymentTerms struct {
	PaymentTermsNumber int    `json:"paymentTermsNumber"`
	DaysOfCredit       int    `json:"daysOfCredit,omitempty"`
	Name               string `json:"name,omitempty"`
	PaymentTermsType   string `json:"paymentTermsType,omitempty"`
}

func (pt OrderDraftPaymentTerms) IsEmpty() bool {
	return zero.IsZero(pt)
}

type OrderDraftVatZone struct {
	Name               string `json:"name"`
	VatZoneNumber      int    `json:"vatZoneNumber,omitempty"`
	EnabledForCustomer bool   `json:"enabledForCustomer"`
	EnabledForSupplier bool   `json:"enabledForSupplier"`
}

func (vz OrderDraftVatZone) IsEmpty() bool {
	return zero.IsZero(vz)
}

type OrderDraftLayout struct {
	LayoutNumber int `json:"layoutNumber,omitempty"`
}

func (l OrderDraftLayout) IsEmpty() bool {
	return zero.IsZero(l)
}
