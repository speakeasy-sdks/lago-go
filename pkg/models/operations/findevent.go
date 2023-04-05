// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type FindEventRequest struct {
	// Id of the existing transaction
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type FindEventResponse struct {
	// Not Found error
	APIResponseNotFound *shared.APIResponseNotFound
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	// Successful response
	Event       *shared.Event
	StatusCode  int
	RawResponse *http.Response
}
