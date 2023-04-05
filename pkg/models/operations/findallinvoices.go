// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"github.com/speakeasy-sdks/lago-go/pkg/types"
	"net/http"
)

// FindAllInvoicesStatusEnum - Status
type FindAllInvoicesStatusEnum string

const (
	FindAllInvoicesStatusEnumDraft     FindAllInvoicesStatusEnum = "draft"
	FindAllInvoicesStatusEnumFinalized FindAllInvoicesStatusEnum = "finalized"
)

func (e *FindAllInvoicesStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "draft":
		fallthrough
	case "finalized":
		*e = FindAllInvoicesStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for FindAllInvoicesStatusEnum: %s", s)
	}
}

type FindAllInvoicesRequest struct {
	// External customer ID
	ExternalCustomerID *string `queryParam:"style=form,explode=true,name=external_customer_id"`
	// Date from
	IssuingDateFrom *types.Date `queryParam:"style=form,explode=true,name=issuing_date_from"`
	// Date to
	IssuingDateTo *types.Date `queryParam:"style=form,explode=true,name=issuing_date_to"`
	// Number of page
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
	// Status
	Status *FindAllInvoicesStatusEnum `queryParam:"style=form,explode=true,name=status"`
}

type FindAllInvoicesResponse struct {
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	// Successful response
	Invoices    *shared.Invoices
	StatusCode  int
	RawResponse *http.Response
}