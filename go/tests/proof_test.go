package tests

import (
	"testing"
	"github.com/centrifuge/functional-testing/go/utils"
	"net/http"
	"github.com/gavv/httpexpect"
)

func TestProofGenerationWithMultipleFields(t *testing.T) {
	e := utils.GetInsecureClient(t, utils.NODE1)

	currency := "USD"
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"invoice_number": "12324",
			"due_date": "2018-09-26T23:12:37.902198664Z",
			"gross_amount": "40",
			"currency": currency,
			"net_amount": "40",
		},
	}

	obj := CreateDocument(t, utils.INVOICE, e, payload)

	docIdentifier := obj.Value("header").Path("$.document_id").String().NotEmpty().Raw()

	proofPayload := map[string]interface{}{
		"type": "http://github.com/centrifuge/centrifuge-protobufs/invoice/#invoice.InvoiceData",
		"fields": []string{"net_amount", "currency"},
	}

	objProof := GetProof(t, e, docIdentifier, proofPayload)
	objProof.Path("$.header.document_id").String().Equal(docIdentifier)
	objProof.Path("$.field_proofs[0].property").String().Equal("net_amount")
}

func GetProof(t *testing.T, e *httpexpect.Expect, documentID string, payload map[string]interface{}) *httpexpect.Object {
	obj := e.POST("/document/" + documentID + "/proof").
		WithHeader("accept", "application/json").
		WithHeader("Content-Type", "application/json").
		WithJSON(payload).
		Expect().Status(http.StatusOK)
	assertOkResponse(t, obj)
	return obj.JSON().Object()
}
