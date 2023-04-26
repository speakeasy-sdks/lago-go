// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// WalletTransactionObjectStatusEnum - Status
type WalletTransactionObjectStatusEnum string

const (
	WalletTransactionObjectStatusEnumPending WalletTransactionObjectStatusEnum = "pending"
	WalletTransactionObjectStatusEnumSettled WalletTransactionObjectStatusEnum = "settled"
)

func (e WalletTransactionObjectStatusEnum) ToPointer() *WalletTransactionObjectStatusEnum {
	return &e
}

func (e *WalletTransactionObjectStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "pending":
		fallthrough
	case "settled":
		*e = WalletTransactionObjectStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for WalletTransactionObjectStatusEnum: %s", s)
	}
}

// WalletTransactionObjectTransactionTypeEnum - Transaction type
type WalletTransactionObjectTransactionTypeEnum string

const (
	WalletTransactionObjectTransactionTypeEnumInbound  WalletTransactionObjectTransactionTypeEnum = "inbound"
	WalletTransactionObjectTransactionTypeEnumOutbound WalletTransactionObjectTransactionTypeEnum = "outbound"
)

func (e WalletTransactionObjectTransactionTypeEnum) ToPointer() *WalletTransactionObjectTransactionTypeEnum {
	return &e
}

func (e *WalletTransactionObjectTransactionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "inbound":
		fallthrough
	case "outbound":
		*e = WalletTransactionObjectTransactionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for WalletTransactionObjectTransactionTypeEnum: %s", s)
	}
}

type WalletTransactionObject struct {
	Amount       float64    `json:"amount"`
	CreatedAt    time.Time  `json:"created_at"`
	CreditAmount float64    `json:"credit_amount"`
	LagoID       string     `json:"lago_id"`
	LagoWalletID string     `json:"lago_wallet_id"`
	SettledAt    *time.Time `json:"settled_at,omitempty"`
	// Status
	Status WalletTransactionObjectStatusEnum `json:"status"`
	// Transaction type
	TransactionType WalletTransactionObjectTransactionTypeEnum `json:"transaction_type"`
}
