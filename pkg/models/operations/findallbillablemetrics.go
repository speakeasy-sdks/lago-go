// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type FindAllBillableMetricsRequest struct {
	// Number of page
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllBillableMetricsResponse struct {
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	// Successful response
	BillableMetrics *shared.BillableMetrics
	ContentType     string
	StatusCode      int
	RawResponse     *http.Response
}
