// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type FindAllWalletsRequest struct {
	// External customer ID
	ExternalCustomerID string `queryParam:"style=form,explode=true,name=external_customer_id"`
	// Number of page
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllWalletsResponse struct {
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int
	RawResponse             *http.Response
	// Successful response
	Wallets *shared.Wallets
}