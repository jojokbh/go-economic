package economic_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/jojokbh/go-economic"
)

func TestQuotesDraftsPost(t *testing.T) {
	client := client()
	client.SetGrantToken("demo")
	client.SetSecretToken("demo")
	req := client.NewQuotesDraftsPostRequest()

	body := req.RequestBody()
	body.Date = economic.Date{time.Now()}
	body.Currency = "DKK"
	body.Customer.CustomerNumber = 1
	body.PaymentTerms.PaymentTermsNumber = 14
	// body.PaymentTerms.PaymentTermsType = "dueDate"
	body.Recipient = economic.QuoteDraftRecipient{
		Name: "Kees Zorge",
		VatZone: economic.QuoteDraftVatZone{
			VatZoneNumber: 1,
		},
	}
	body.Layout = economic.QuoteDraftLayout{
		LayoutNumber: 19,
	}
	line := economic.QuoteDraftLine{
		Description: "test",
	}
	body.Lines = append(body.Lines, line)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
