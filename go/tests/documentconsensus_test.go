package tests

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/centrifuge/functional-testing/go/utils"
	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndUpdateInvoiceFromOrigin(t *testing.T) {
	// nodes
	e := utils.GetInsecureClient(t, utils.NODE1)
	e1 := utils.GetInsecureClient(t, utils.NODE2)

	// create invoice
	currency := "USD"
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_due":     "2018-09-26T23:12:37.902198664Z",
			"gross_amount": "40",
			"currency":     currency,
			"net_amount":   "40",
		},
		"write_access": map[string]interface{}{
			"collaborators": []string{utils.Nodes[utils.NODE2].ID},
		},
	}

	obj := CreateDocument(t, utils.INVOICE, e, utils.Nodes[utils.NODE1].ID, payload)

	docIdentifier := obj.Value("header").Path("$.document_id").String().NotEmpty().Raw()

	doc := GetDocument(t, utils.INVOICE, e, utils.Nodes[utils.NODE1].ID, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// Receiver has document
	doc = GetDocument(t, utils.INVOICE, e1, utils.Nodes[utils.NODE2].ID, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// update invoice
	payload = map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_due":     "2018-09-26T23:12:37.902198664Z",
			"gross_amount": "41",
			"currency":     currency,
			"net_amount":   "41",
		},
		"write_access": map[string]interface{}{
			"collaborators": []string{utils.Nodes[utils.NODE2].ID},
		},
	}

	obj = UpdateDocument(t, utils.INVOICE, e, utils.Nodes[utils.NODE1].ID, docIdentifier, payload)

	// check updated gross amount
	obj.Value("data").Path("$.gross_amount").String().Equal("41")
	GetDocument(t, utils.INVOICE, e, utils.Nodes[utils.NODE1].ID, docIdentifier)

	// Receiver has document
	doc = GetDocument(t, utils.INVOICE, e1, utils.Nodes[utils.NODE2].ID, docIdentifier)
	doc.Path("$.data.gross_amount").String().Equal("41")
}

func TestCreateAndUpdateInvoiceFromCollaborator(t *testing.T) {
	// nodes
	e := utils.GetInsecureClient(t, utils.NODE1)
	e1 := utils.GetInsecureClient(t, utils.NODE2)

	// create invoice
	currency := "USD"
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_due":     "2018-09-26T23:12:37.902198664Z",
			"gross_amount": "40",
			"currency":     currency,
			"net_amount":   "40",
		},
		"write_access": map[string]interface{}{
			"collaborators": []string{utils.Nodes[utils.NODE2].ID},
		},
	}

	obj := CreateDocument(t, utils.INVOICE, e, utils.Nodes[utils.NODE1].ID, payload)

	docIdentifier := obj.Value("header").Path("$.document_id").String().NotEmpty().Raw()

	doc := GetDocument(t, utils.INVOICE, e, utils.Nodes[utils.NODE1].ID, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// Receiver has document
	doc = GetDocument(t, utils.INVOICE, e1, utils.Nodes[utils.NODE2].ID, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// update invoice
	payload = map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_due":     "2018-09-26T23:12:37.902198664Z",
			"gross_amount": "41",
			"currency":     currency,
			"net_amount":   "41",
		},
		"write_access": map[string]interface{}{
			"collaborators": []string{utils.Nodes[utils.NODE1].ID},
		},
	}

	obj = UpdateDocument(t, utils.INVOICE, e1, utils.Nodes[utils.NODE2].ID, docIdentifier, payload)

	// check updated gross amount
	obj.Value("data").Path("$.gross_amount").String().Equal("41")
	GetDocument(t, utils.INVOICE, e1, utils.Nodes[utils.NODE2].ID, docIdentifier)

	// Receiver has document
	doc = GetDocument(t, utils.INVOICE, e, utils.Nodes[utils.NODE1].ID, docIdentifier)
	doc.Path("$.data.gross_amount").String().Equal("41")
}

func TestCreateAndUpdatePurchaseOrderFromOrigin(t *testing.T) {
	// nodes
	e := utils.GetInsecureClient(t, utils.NODE1)
	e1 := utils.GetInsecureClient(t, utils.NODE2)

	// create invoice
	currency := "USD"
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_created": "2018-09-26T23:12:37.902198664Z",
			"total_amount": "40",
			"currency":     currency,
		},
		"write_access": map[string]interface{}{
			"collaborators": []string{utils.Nodes[utils.NODE2].ID},
		},
	}

	obj := CreateDocument(t, utils.PURCHASEORDER, e, utils.Nodes[utils.NODE1].ID, payload)

	docIdentifier := obj.Value("header").Path("$.document_id").String().NotEmpty().Raw()

	doc := GetDocument(t, utils.PURCHASEORDER, e, utils.Nodes[utils.NODE1].ID, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// Receiver has document
	doc = GetDocument(t, utils.PURCHASEORDER, e1, utils.Nodes[utils.NODE2].ID, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// update invoice
	payload = map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_created": "2018-09-26T23:12:37.902198664Z",
			"total_amount": "41",
			"currency":     currency,
		},
		"write_access": map[string]interface{}{
			"collaborators": []string{utils.Nodes[utils.NODE2].ID},
		},
	}

	obj = UpdateDocument(t, utils.PURCHASEORDER, e, utils.Nodes[utils.NODE1].ID, docIdentifier, payload)

	// check updated gross amount
	obj.Value("data").Path("$.total_amount").String().Equal("41")
	GetDocument(t, utils.PURCHASEORDER, e, utils.Nodes[utils.NODE1].ID, docIdentifier)

	// Receiver has document
	doc = GetDocument(t, utils.PURCHASEORDER, e1, utils.Nodes[utils.NODE2].ID, docIdentifier)
	doc.Path("$.data.total_amount").String().Equal("41")
}

func TestCreateAndUpdatePurchaseOrderFromCollaborator(t *testing.T) {
	// nodes
	e := utils.GetInsecureClient(t, utils.NODE1)
	e1 := utils.GetInsecureClient(t, utils.NODE2)

	// create invoice
	currency := "USD"
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_created": "2018-09-26T23:12:37.902198664Z",
			"total_amount": "40",
			"currency":     currency,
		},
		"write_access": map[string]interface{}{
			"collaborators": []string{utils.Nodes[utils.NODE2].ID},
		},
	}

	obj := CreateDocument(t, utils.PURCHASEORDER, e, utils.Nodes[utils.NODE1].ID, payload)

	docIdentifier := obj.Value("header").Path("$.document_id").String().NotEmpty().Raw()

	doc := GetDocument(t, utils.PURCHASEORDER, e, utils.Nodes[utils.NODE1].ID, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// Receiver has document
	doc = GetDocument(t, utils.PURCHASEORDER, e1, utils.Nodes[utils.NODE2].ID, docIdentifier)
	doc.Path("$.data.currency").String().Equal(currency)

	// update invoice
	payload = map[string]interface{}{
		"data": map[string]interface{}{
			"number":       "12324",
			"date_created": "2018-09-26T23:12:37.902198664Z",
			"total_amount": "41",
			"currency":     currency,
		},
		"write_access": map[string]interface{}{
			"collaborators": []string{utils.Nodes[utils.NODE2].ID},
		},
	}

	obj = UpdateDocument(t, utils.PURCHASEORDER, e1, utils.Nodes[utils.NODE2].ID, docIdentifier, payload)

	// check updated gross amount
	obj.Value("data").Path("$.total_amount").String().Equal("41")
	GetDocument(t, utils.PURCHASEORDER, e1, utils.Nodes[utils.NODE2].ID, docIdentifier)

	// Receiver has document
	doc = GetDocument(t, utils.PURCHASEORDER, e, utils.Nodes[utils.NODE1].ID, docIdentifier)
	doc.Path("$.data.total_amount").String().Equal("41")
}

func GetDocument(t *testing.T, docType string, e *httpexpect.Expect, auth string, docIdentifier string) *httpexpect.Value {
	objGet := utils.AddCommonHeaders(e.GET(fmt.Sprintf("/%s/%s", docType, docIdentifier)), auth).
		Expect().Status(http.StatusOK)
	assertOkResponse(t, objGet)
	objGet.JSON().Path("$.header.document_id").String().Equal(docIdentifier)
	return objGet.JSON()
}

func CreateDocument(t *testing.T, docType string, e *httpexpect.Expect, auth string, payload map[string]interface{}) *httpexpect.Object {
	path := fmt.Sprintf("/%s", docType)
	method := "POST"
	resp := getResponse(method, path, e, auth, payload).Status(http.StatusOK)
	assertOkResponse(t, resp)
	obj := resp.JSON().Object()
	txID := getTransactionID(t, obj)
	waitTillSuccess(t, e, auth, txID)
	return obj
}

func UpdateDocument(t *testing.T, docType string, e *httpexpect.Expect, auth string, documentID string, payload map[string]interface{}) *httpexpect.Object {
	path := fmt.Sprintf("/%s/%s", docType, documentID)
	method := "PUT"
	resp := getResponse(method, path, e, auth, payload).Status(http.StatusOK)
	assertOkResponse(t, resp)
	obj := resp.JSON().Object()
	txID := getTransactionID(t, obj)
	waitTillSuccess(t, e, auth, txID)
	return obj
}

func getResponse(method, path string, e *httpexpect.Expect, auth string, payload map[string]interface{}) *httpexpect.Response {
	return utils.AddCommonHeaders(e.Request(method, path), auth).
		WithJSON(payload).
		Expect()
}

func assertOkResponse(t *testing.T, response *httpexpect.Response) {
	if response.Raw().StatusCode != http.StatusOK {
		assert.Fail(t, "Response Payload: ", response.Body().Raw())
	}
}

func getTransactionID(t *testing.T, resp *httpexpect.Object) string {
	txID := resp.Value("header").Path("$.job_id").String().Raw()
	if txID == "" {
		t.Error("transaction ID empty")
	}

	return txID
}

func waitTillSuccess(t *testing.T, e *httpexpect.Expect, auth string, txID string) {
	for {
		resp := utils.AddCommonHeaders(e.GET("/jobs/"+txID), auth).Expect().Status(200).JSON().Object()
		status := resp.Path("$.status").String().Raw()
		if status == "pending" {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if status == "failed" {
			t.Error(resp.Path("$.message").String().Raw())
		}

		break
	}
}
