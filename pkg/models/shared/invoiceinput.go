// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type InvoiceInputInvoiceMetadata struct {
	ID    *string `json:"id,omitempty"`
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// InvoiceInputInvoicePaymentStatusEnum - Status
type InvoiceInputInvoicePaymentStatusEnum string

const (
	InvoiceInputInvoicePaymentStatusEnumPending   InvoiceInputInvoicePaymentStatusEnum = "pending"
	InvoiceInputInvoicePaymentStatusEnumSucceeded InvoiceInputInvoicePaymentStatusEnum = "succeeded"
	InvoiceInputInvoicePaymentStatusEnumFailed    InvoiceInputInvoicePaymentStatusEnum = "failed"
)

func (e *InvoiceInputInvoicePaymentStatusEnum) UnmarshalJSON(data []byte) error {
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
		*e = InvoiceInputInvoicePaymentStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InvoiceInputInvoicePaymentStatusEnum: %s", s)
	}
}

type InvoiceInputInvoice struct {
	Metadata []InvoiceInputInvoiceMetadata `json:"metadata,omitempty"`
	// Status
	PaymentStatus InvoiceInputInvoicePaymentStatusEnum `json:"payment_status"`
}

// InvoiceInput - Update an existing invoice
type InvoiceInput struct {
	Invoice InvoiceInputInvoice `json:"invoice"`
}
