// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type FindAllBillableMetricGroupsRequest struct {
	// Code of the existing billable metric
	Code string `pathParam:"style=simple,explode=false,name=code"`
	// Number of page
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type FindAllBillableMetricGroupsResponse struct {
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	// Successful response
	Groups      *shared.Groups
	StatusCode  int
	RawResponse *http.Response
}
