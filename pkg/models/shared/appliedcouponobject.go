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

func (e *AppliedCouponObjectFrequencyEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "once":
		fallthrough
	case "recurring":
		*e = AppliedCouponObjectFrequencyEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for AppliedCouponObjectFrequencyEnum: %s", s)
	}
}

// AppliedCouponObjectStatusEnum - Status
type AppliedCouponObjectStatusEnum string

const (
	AppliedCouponObjectStatusEnumActive     AppliedCouponObjectStatusEnum = "active"
	AppliedCouponObjectStatusEnumTerminated AppliedCouponObjectStatusEnum = "terminated"
)

func (e *AppliedCouponObjectStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "active":
		fallthrough
	case "terminated":
		*e = AppliedCouponObjectStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for AppliedCouponObjectStatusEnum: %s", s)
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
