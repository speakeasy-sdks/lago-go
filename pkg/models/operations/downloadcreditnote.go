// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type DownloadCreditNoteRequest struct {
	// ID of the existing Lago Credit note
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DownloadCreditNoteResponse struct {
	// Not Found error
	APIResponseNotFound *shared.APIResponseNotFound
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	// Successful response
	CreditNote  *shared.CreditNote
	StatusCode  int
	RawResponse *http.Response
}