// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
	"time"
)

// FindAllFeesFeeTypeEnum - Fee type
type FindAllFeesFeeTypeEnum string

const (
	FindAllFeesFeeTypeEnumCharge        FindAllFeesFeeTypeEnum = "charge"
	FindAllFeesFeeTypeEnumAddOn         FindAllFeesFeeTypeEnum = "add_on"
	FindAllFeesFeeTypeEnumSubscription  FindAllFeesFeeTypeEnum = "subscription"
	FindAllFeesFeeTypeEnumCredit        FindAllFeesFeeTypeEnum = "credit"
	FindAllFeesFeeTypeEnumInstantCharge FindAllFeesFeeTypeEnum = "instant_charge"
)

func (e *FindAllFeesFeeTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "charge":
		fallthrough
	case "add_on":
		fallthrough
	case "subscription":
		fallthrough
	case "credit":
		fallthrough
	case "instant_charge":
		*e = FindAllFeesFeeTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for FindAllFeesFeeTypeEnum: %s", s)
	}
}

// FindAllFeesPaymentStatusEnum - Payment status
type FindAllFeesPaymentStatusEnum string

const (
	FindAllFeesPaymentStatusEnumPending   FindAllFeesPaymentStatusEnum = "pending"
	FindAllFeesPaymentStatusEnumSucceeded FindAllFeesPaymentStatusEnum = "succeeded"
	FindAllFeesPaymentStatusEnumFailed    FindAllFeesPaymentStatusEnum = "failed"
	FindAllFeesPaymentStatusEnumRefunded  FindAllFeesPaymentStatusEnum = "refunded"
)

func (e *FindAllFeesPaymentStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "pending":
		fallthrough
	case "succeeded":
		fallthrough
	case "failed":
		fallthrough
	case "refunded":
		*e = FindAllFeesPaymentStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for FindAllFeesPaymentStatusEnum: %s", s)
	}
}

type FindAllFeesRequest struct {
	// Code of the source billable metric
	BillableMetricCode *string `queryParam:"style=form,explode=true,name=billable_metric_code"`
	// Creation datetime from
	CreatedAtFrom *time.Time `queryParam:"style=form,explode=true,name=created_at_from"`
	// Creation date to
	CreatedAtTo *time.Time `queryParam:"style=form,explode=true,name=created_at_to"`
	// Amount currency
	Currency *string `queryParam:"style=form,explode=true,name=currency"`
	// External customer ID
	ExternalCustomerID *string `queryParam:"style=form,explode=true,name=external_customer_id"`
	// External subscription ID
	ExternalSubscriptionID *string `queryParam:"style=form,explode=true,name=external_subscription_id"`
	// Payment failed date from
	FailedAtFrom *time.Time `queryParam:"style=form,explode=true,name=failed_at_from"`
	// Payment failed date to
	FailedAtTo *time.Time `queryParam:"style=form,explode=true,name=failed_at_to"`
	// Fee type
	FeeType *FindAllFeesFeeTypeEnum `queryParam:"style=form,explode=true,name=fee_type"`
	// Number of page
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Payment status
	PaymentStatus *FindAllFeesPaymentStatusEnum `queryParam:"style=form,explode=true,name=payment_status"`
	// Number of records per page
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
	// Payment refund date from
	RefundedAtFrom *time.Time `queryParam:"style=form,explode=true,name=refunded_at_from"`
	// Payment refund date to
	RefundedAtTo *time.Time `queryParam:"style=form,explode=true,name=refunded_at_to"`
	// Payment succees date from
	SucceededAtFrom *time.Time `queryParam:"style=form,explode=true,name=succeeded_at_from"`
	// Payment succees date to
	SucceededAtTo *time.Time `queryParam:"style=form,explode=true,name=succeeded_at_to"`
}

type FindAllFeesResponse struct {
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	// Unprocessable entity error
	APIResponseUnprocessableEntity *shared.APIResponseUnprocessableEntity
	ContentType                    string
	// Successful response
	FeesPaginated *shared.FeesPaginated
	StatusCode    int
	RawResponse   *http.Response
}
