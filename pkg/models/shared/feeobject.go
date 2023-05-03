// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type FeeObjectItemItemTypeEnum string

const (
	FeeObjectItemItemTypeEnumAddOn             FeeObjectItemItemTypeEnum = "AddOn"
	FeeObjectItemItemTypeEnumBillableMetric    FeeObjectItemItemTypeEnum = "BillableMetric"
	FeeObjectItemItemTypeEnumSubscription      FeeObjectItemItemTypeEnum = "Subscription"
	FeeObjectItemItemTypeEnumWalletTransaction FeeObjectItemItemTypeEnum = "WalletTransaction"
)

func (e FeeObjectItemItemTypeEnum) ToPointer() *FeeObjectItemItemTypeEnum {
	return &e
}

func (e *FeeObjectItemItemTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AddOn":
		fallthrough
	case "BillableMetric":
		fallthrough
	case "Subscription":
		fallthrough
	case "WalletTransaction":
		*e = FeeObjectItemItemTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FeeObjectItemItemTypeEnum: %v", v)
	}
}

// FeeObjectItemTypeEnum - Billing time
type FeeObjectItemTypeEnum string

const (
	FeeObjectItemTypeEnumCharge       FeeObjectItemTypeEnum = "charge"
	FeeObjectItemTypeEnumAddOn        FeeObjectItemTypeEnum = "add_on"
	FeeObjectItemTypeEnumSubscription FeeObjectItemTypeEnum = "subscription"
	FeeObjectItemTypeEnumCredit       FeeObjectItemTypeEnum = "credit"
)

func (e FeeObjectItemTypeEnum) ToPointer() *FeeObjectItemTypeEnum {
	return &e
}

func (e *FeeObjectItemTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "charge":
		fallthrough
	case "add_on":
		fallthrough
	case "subscription":
		fallthrough
	case "credit":
		*e = FeeObjectItemTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FeeObjectItemTypeEnum: %v", v)
	}
}

type FeeObjectItem struct {
	Code       string                    `json:"code"`
	ItemType   FeeObjectItemItemTypeEnum `json:"item_type"`
	LagoItemID string                    `json:"lago_item_id"`
	Name       string                    `json:"name"`
	// Billing time
	Type FeeObjectItemTypeEnum `json:"type"`
}

type FeeObjectPaymentStatusEnum string

const (
	FeeObjectPaymentStatusEnumPending   FeeObjectPaymentStatusEnum = "pending"
	FeeObjectPaymentStatusEnumSucceeded FeeObjectPaymentStatusEnum = "succeeded"
	FeeObjectPaymentStatusEnumFailed    FeeObjectPaymentStatusEnum = "failed"
	FeeObjectPaymentStatusEnumRefunded  FeeObjectPaymentStatusEnum = "refunded"
)

func (e FeeObjectPaymentStatusEnum) ToPointer() *FeeObjectPaymentStatusEnum {
	return &e
}

func (e *FeeObjectPaymentStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "succeeded":
		fallthrough
	case "failed":
		fallthrough
	case "refunded":
		*e = FeeObjectPaymentStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FeeObjectPaymentStatusEnum: %v", v)
	}
}

// FeeObject - Successful response
type FeeObject struct {
	AmountCents            int64                      `json:"amount_cents"`
	AmountCurrency         string                     `json:"amount_currency"`
	CreatedAt              time.Time                  `json:"created_at"`
	EventsCount            *int64                     `json:"events_count,omitempty"`
	ExternalSubscriptionID *string                    `json:"external_subscription_id,omitempty"`
	FailedAt               *time.Time                 `json:"failed_at,omitempty"`
	FromDate               *time.Time                 `json:"from_date,omitempty"`
	Item                   FeeObjectItem              `json:"item"`
	LagoGroupID            *string                    `json:"lago_group_id,omitempty"`
	LagoID                 string                     `json:"lago_id"`
	LagoInvoiceID          *string                    `json:"lago_invoice_id,omitempty"`
	PaymentStatus          FeeObjectPaymentStatusEnum `json:"payment_status"`
	RefundedAt             *time.Time                 `json:"refunded_at,omitempty"`
	SucceededAt            *time.Time                 `json:"succeeded_at,omitempty"`
	ToDate                 *time.Time                 `json:"to_date,omitempty"`
	TotalAmountCents       *int64                     `json:"total_amount_cents,omitempty"`
	TotalAmountCurrency    *string                    `json:"total_amount_currency,omitempty"`
	Units                  float64                    `json:"units"`
	VatAmountCents         int64                      `json:"vat_amount_cents"`
	VatAmountCurrency      string                     `json:"vat_amount_currency"`
}
