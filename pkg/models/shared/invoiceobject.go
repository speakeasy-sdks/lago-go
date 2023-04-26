// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type InvoiceObjectInvoiceTypeEnum string

const (
	InvoiceObjectInvoiceTypeEnumSubscription InvoiceObjectInvoiceTypeEnum = "subscription"
	InvoiceObjectInvoiceTypeEnumAddOn        InvoiceObjectInvoiceTypeEnum = "add_on"
	InvoiceObjectInvoiceTypeEnumCredit       InvoiceObjectInvoiceTypeEnum = "credit"
)

func (e InvoiceObjectInvoiceTypeEnum) ToPointer() *InvoiceObjectInvoiceTypeEnum {
	return &e
}

func (e *InvoiceObjectInvoiceTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "subscription":
		fallthrough
	case "add_on":
		fallthrough
	case "credit":
		*e = InvoiceObjectInvoiceTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InvoiceObjectInvoiceTypeEnum: %s", s)
	}
}

type InvoiceObjectPaymentStatusEnum string

const (
	InvoiceObjectPaymentStatusEnumPending   InvoiceObjectPaymentStatusEnum = "pending"
	InvoiceObjectPaymentStatusEnumSucceeded InvoiceObjectPaymentStatusEnum = "succeeded"
	InvoiceObjectPaymentStatusEnumFailed    InvoiceObjectPaymentStatusEnum = "failed"
)

func (e InvoiceObjectPaymentStatusEnum) ToPointer() *InvoiceObjectPaymentStatusEnum {
	return &e
}

func (e *InvoiceObjectPaymentStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = InvoiceObjectPaymentStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InvoiceObjectPaymentStatusEnum: %s", s)
	}
}

type InvoiceObjectStatusEnum string

const (
	InvoiceObjectStatusEnumDraft     InvoiceObjectStatusEnum = "draft"
	InvoiceObjectStatusEnumFinalized InvoiceObjectStatusEnum = "finalized"
)

func (e InvoiceObjectStatusEnum) ToPointer() *InvoiceObjectStatusEnum {
	return &e
}

func (e *InvoiceObjectStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "draft":
		fallthrough
	case "finalized":
		*e = InvoiceObjectStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InvoiceObjectStatusEnum: %s", s)
	}
}

type InvoiceObject struct {
	AmountCents          int64                          `json:"amount_cents"`
	AmountCurrency       string                         `json:"amount_currency"`
	ChargesFromDate      time.Time                      `json:"charges_from_date"`
	CreditAmountCents    int64                          `json:"credit_amount_cents"`
	CreditAmountCurrency string                         `json:"credit_amount_currency"`
	Credits              []CreditObject                 `json:"credits"`
	Customer             CustomerObject                 `json:"customer"`
	Fees                 []FeeObject                    `json:"fees"`
	FileURL              *string                        `json:"file_url,omitempty"`
	InvoiceType          *InvoiceObjectInvoiceTypeEnum  `json:"invoice_type,omitempty"`
	IssuingDate          time.Time                      `json:"issuing_date"`
	LagoID               string                         `json:"lago_id"`
	Legacy               bool                           `json:"legacy"`
	Metadata             []InvoiceMetadataObject        `json:"metadata,omitempty"`
	Number               string                         `json:"number"`
	PaymentStatus        InvoiceObjectPaymentStatusEnum `json:"payment_status"`
	SequentialID         int64                          `json:"sequential_id"`
	Status               InvoiceObjectStatusEnum        `json:"status"`
	Subscriptions        []SubscriptionObject           `json:"subscriptions"`
	TotalAmountCents     int64                          `json:"total_amount_cents"`
	TotalAmountCurrency  string                         `json:"total_amount_currency"`
	VatAmountCents       int64                          `json:"vat_amount_cents"`
	VatAmountCurrency    string                         `json:"vat_amount_currency"`
}
