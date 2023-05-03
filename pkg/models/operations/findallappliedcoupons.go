// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

// FindAllAppliedCouponsStatusEnum - Status
type FindAllAppliedCouponsStatusEnum string

const (
	FindAllAppliedCouponsStatusEnumActive     FindAllAppliedCouponsStatusEnum = "active"
	FindAllAppliedCouponsStatusEnumTerminated FindAllAppliedCouponsStatusEnum = "terminated"
)

func (e FindAllAppliedCouponsStatusEnum) ToPointer() *FindAllAppliedCouponsStatusEnum {
	return &e
}

func (e *FindAllAppliedCouponsStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "terminated":
		*e = FindAllAppliedCouponsStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindAllAppliedCouponsStatusEnum: %v", v)
	}
}

type FindAllAppliedCouponsRequest struct {
	// External customer ID
	ExternalCustomerID *string `queryParam:"style=form,explode=true,name=external_customer_id"`
	// Number of page
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
	// Applied coupon status
	Status *FindAllAppliedCouponsStatusEnum `queryParam:"style=form,explode=true,name=status"`
}

type FindAllAppliedCouponsResponse struct {
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	// Successful response
	AppliedCoupons *shared.AppliedCoupons
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}
