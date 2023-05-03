// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// AppliedCouponObjectFrequencyEnum - Frequency type
type AppliedCouponObjectFrequencyEnum string

const (
	AppliedCouponObjectFrequencyEnumOnce      AppliedCouponObjectFrequencyEnum = "once"
	AppliedCouponObjectFrequencyEnumRecurring AppliedCouponObjectFrequencyEnum = "recurring"
)

func (e AppliedCouponObjectFrequencyEnum) ToPointer() *AppliedCouponObjectFrequencyEnum {
	return &e
}

func (e *AppliedCouponObjectFrequencyEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "once":
		fallthrough
	case "recurring":
		*e = AppliedCouponObjectFrequencyEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppliedCouponObjectFrequencyEnum: %v", v)
	}
}

// AppliedCouponObjectStatusEnum - Status
type AppliedCouponObjectStatusEnum string

const (
	AppliedCouponObjectStatusEnumActive     AppliedCouponObjectStatusEnum = "active"
	AppliedCouponObjectStatusEnumTerminated AppliedCouponObjectStatusEnum = "terminated"
)

func (e AppliedCouponObjectStatusEnum) ToPointer() *AppliedCouponObjectStatusEnum {
	return &e
}

func (e *AppliedCouponObjectStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "terminated":
		*e = AppliedCouponObjectStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppliedCouponObjectStatusEnum: %v", v)
	}
}

type AppliedCouponObject struct {
	AmountCents          int64          `json:"amount_cents"`
	AmountCentsRemaining *int64         `json:"amount_cents_remaining,omitempty"`
	AmountCurrency       string         `json:"amount_currency"`
	CouponCode           string         `json:"coupon_code"`
	CreatedAt            time.Time      `json:"created_at"`
	Credits              []CreditObject `json:"credits"`
	ExpirationAt         *time.Time     `json:"expiration_at,omitempty"`
	ExternalCustomerID   string         `json:"external_customer_id"`
	// Frequency type
	Frequency                  AppliedCouponObjectFrequencyEnum `json:"frequency"`
	FrequencyDuration          *int64                           `json:"frequency_duration,omitempty"`
	FrequencyDurationRemaining *int64                           `json:"frequency_duration_remaining,omitempty"`
	LagoCouponID               string                           `json:"lago_coupon_id"`
	LagoCustomerID             string                           `json:"lago_customer_id"`
	LagoID                     string                           `json:"lago_id"`
	PercentageRate             *float64                         `json:"percentage_rate,omitempty"`
	// Status
	Status       AppliedCouponObjectStatusEnum `json:"status"`
	TerminatedAt *time.Time                    `json:"terminated_at,omitempty"`
}
