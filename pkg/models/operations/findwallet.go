// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type FindWalletRequest struct {
	// Lago ID of the existing wallet
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type FindWalletResponse struct {
	// Not Found error
	APIResponseNotFound *shared.APIResponseNotFound
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int
	RawResponse             *http.Response
	// Successful response
	Wallet *shared.Wallet
}