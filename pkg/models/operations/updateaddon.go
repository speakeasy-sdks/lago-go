// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type UpdateAddOnRequest struct {
	// Update an existing add-on
	AddOnInput shared.AddOnInput `request:"mediaType=application/json"`
	// Code of the existing add-on
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type UpdateAddOnResponse struct {
	// Successful response
	AddOn *shared.AddOn
	// Bad Request error
	APIResponseBadRequest *shared.APIResponseBadRequest
	// Not Found error
	APIResponseNotFound *shared.APIResponseNotFound
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	// Unprocessable entity error
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
}
