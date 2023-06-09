// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type FetchPublicKeyResponse struct {
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int
	RawResponse             *http.Response
	// Successful response
	FetchPublicKey200TextPlainString *string
}
