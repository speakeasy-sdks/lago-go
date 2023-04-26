// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// WalletObjectStatusEnum - Status
type WalletObjectStatusEnum string

const (
	WalletObjectStatusEnumActive     WalletObjectStatusEnum = "active"
	WalletObjectStatusEnumTerminated WalletObjectStatusEnum = "terminated"
)

func (e WalletObjectStatusEnum) ToPointer() *WalletObjectStatusEnum {
	return &e
}

func (e *WalletObjectStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "active":
		fallthrough
	case "terminated":
		*e = WalletObjectStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for WalletObjectStatusEnum: %s", s)
	}
}

type WalletObject struct {
	Balance              float64    `json:"balance"`
	ConsumedCredits      float64    `json:"consumed_credits"`
	CreatedAt            time.Time  `json:"created_at"`
	CreditsBalance       float64    `json:"credits_balance"`
	Currency             string     `json:"currency"`
	ExpirationAt         *time.Time `json:"expiration_at,omitempty"`
	ExternalCustomerID   string     `json:"external_customer_id"`
	LagoCustomerID       string     `json:"lago_customer_id"`
	LagoID               string     `json:"lago_id"`
	LastBalanceSyncAt    *time.Time `json:"last_balance_sync_at,omitempty"`
	LastConsumedCreditAt *time.Time `json:"last_consumed_credit_at,omitempty"`
	Name                 *string    `json:"name,omitempty"`
	RateAmount           float64    `json:"rate_amount"`
	// Status
	Status       WalletObjectStatusEnum `json:"status"`
	TerminatedAt *time.Time             `json:"terminated_at,omitempty"`
}
