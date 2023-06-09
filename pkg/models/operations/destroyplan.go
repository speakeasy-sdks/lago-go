// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type DestroyPlanRequest struct {
	// Code of the existing plan
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type DestroyPlanResponse struct {
	// Not Found error
	APIResponseNotFound *shared.APIResponseNotFound
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	// Successful response
	Plan        *shared.Plan
	StatusCode  int
	RawResponse *http.Response
}
