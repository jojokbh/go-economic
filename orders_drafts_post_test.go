package economic_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/jojokbh/go-economic"
)

func TestOrdersDraftsPost(t *testing.T) {
	client := client()
	req := client.NewOrdersDraftsPostRequest()

	body := req.RequestBody()
	body.Date = economic.Date{time.Now()}
	body.Currency = "EUR"
	body.Customer.CustomerNumber = 1
	body.PaymentTerms.PaymentTermsNumber = 14
	// body.PaymentTerms.PaymentTermsType = "dueDate"
	body.Recipient = economic.OrderDraftRecipient{
		Name: "Kees Zorge",
		VatZone: economic.OrderDraftVatZone{
			VatZoneNumber: 1,
		},
	}
	body.Layout = economic.OrderDraftLayout{
		LayoutNumber: 19,
	}
	line := economic.OrderDraftLine{}
	body.Lines = append(body.Lines, line)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
