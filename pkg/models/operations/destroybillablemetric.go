// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type DestroyBillableMetricRequest struct {
	// Code of the existing billable metric
	Code string `pathParam:"style=simple,explode=false,name=code"`
}

type DestroyBillableMetricResponse struct {
	// Not Found error
	APIResponseNotFound *shared.APIResponseNotFound
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	// Successful response
	BillableMetric *shared.BillableMetric
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}