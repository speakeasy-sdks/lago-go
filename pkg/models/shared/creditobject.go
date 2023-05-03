// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreditObjectInvoicePaymentStatusEnum string

const (
	CreditObjectInvoicePaymentStatusEnumPending   CreditObjectInvoicePaymentStatusEnum = "pending"
	CreditObjectInvoicePaymentStatusEnumSucceeded CreditObjectInvoicePaymentStatusEnum = "succeeded"
	CreditObjectInvoicePaymentStatusEnumFailed    CreditObjectInvoicePaymentStatusEnum = "failed"
)

func (e CreditObjectInvoicePaymentStatusEnum) ToPointer() *CreditObjectInvoicePaymentStatusEnum {
	return &e
}

func (e *CreditObjectInvoicePaymentStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = CreditObjectInvoicePaymentStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreditObjectInvoicePaymentStatusEnum: %v", v)
	}
}

type CreditObjectInvoice struct {
	LagoID        string                               `json:"lago_id"`
	PaymentStatus CreditObjectInvoicePaymentStatusEnum `json:"payment_status"`
}

type CreditObjectItem struct {
	Code   string `json:"code"`
	LagoID string `json:"lago_id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

type CreditObject struct {
	AmountCents    int64               `json:"amount_cents"`
	AmountCurrency string              `json:"amount_currency"`
	Invoice        CreditObjectInvoice `json:"invoice"`
	Item           CreditObjectItem    `json:"item"`
	LagoID         string              `json:"lago_id"`
}
