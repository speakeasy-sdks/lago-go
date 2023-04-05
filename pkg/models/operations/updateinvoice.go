// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type UpdateInvoiceRequest struct {
	// Update an existing invoice
	InvoiceInput shared.InvoiceInput `request:"mediaType=application/json"`
	// ID of the existing Lago Invoice
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateInvoiceResponse struct {
	// Bad Request error
	APIResponseBadRequest *shared.APIResponseBadRequest
	// Not Found error
	APIResponseNotFound *shared.APIResponseNotFound
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	// Unprocessable entity error
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	// Successful response
	Invoice     *shared.Invoice
	StatusCode  int
	RawResponse *http.Response
}
