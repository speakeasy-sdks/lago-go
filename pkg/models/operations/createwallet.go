// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type CreateWalletResponse struct {
	// Bad Request error
	APIResponseBadRequest *shared.APIResponseBadRequest
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	// Unprocessable entity error
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
	// Successful response
	Wallet *shared.Wallet
}
