package types

import (
	"fmt"
	"time"
)

type Transaction struct {
	Amount         int           `json:"amount"`
	Txid           string        `json:"txid"`
	AddressTo      *string       `json:"addressTo,omitempty"`
	AddressFrom    *string       `json:"addressFrom,omitempty"`
	Type           PaymentType   `json:"type"`
	IsSpent        bool          `json:"isSpent"`
	IsMempool      bool          `json:"isMempool"`
	Status         PaymentStatus `json:"status"`
	OrganizationId *string        `json:"organizationId,omitempty"`
	AddressId      string        `json:"addressId"`
	Date           time.Time     `json:"date"`
}

type PaymentType string

const (
	PaymentTypeWithdraw PaymentType = "Withdraw"
	PaymentTypeDeposit  PaymentType = "Deposit"
)

// String method for fmt.Stringer interface
func (p PaymentType) String() string {
	return string(p)
}

// IsValid checks if the payment type is valid
func (p PaymentType) IsValid() bool {
	switch p {
	case PaymentTypeWithdraw, PaymentTypeDeposit:
		return true
	default:
		return false
	}
}

// IsWithdraw returns true if this is a withdraw payment
func (p PaymentType) IsWithdraw() bool {
	return p == PaymentTypeWithdraw
}

// IsDeposit returns true if this is a deposit payment
func (p PaymentType) IsDeposit() bool {
	return p == PaymentTypeDeposit
}

// MarshalJSON implements json.Marshaler interface
func (p PaymentType) MarshalJSON() ([]byte, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid payment type: %s", p)
	}
	return []byte(`"` + string(p) + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (p *PaymentType) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("invalid JSON format for payment type: %s", string(data))
	}

	str := string(data[1 : len(data)-1])
	paymentType := PaymentType(str)

	if !paymentType.IsValid() {
		return fmt.Errorf("invalid payment type: %s. Valid values are: Withdraw, Deposit", str)
	}

	*p = paymentType
	return nil
}

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "Pending"
	PaymentStatusConfirmed PaymentStatus = "Confirmed"
)

// String method for fmt.Stringer interface
func (p PaymentStatus) String() string {
	return string(p)
}

// IsValid checks if the payment status is valid
func (p PaymentStatus) IsValid() bool {
	switch p {
	case PaymentStatusPending, PaymentStatusConfirmed:
		return true
	default:
		return false
	}
}

// IsPending returns true if the payment is pending
func (p PaymentStatus) IsPending() bool {
	return p == PaymentStatusPending
}

// IsConfirmed returns true if the payment is confirmed
func (p PaymentStatus) IsConfirmed() bool {
	return p == PaymentStatusConfirmed
}

// IsFinal returns true if the payment is in a final state
func (p PaymentStatus) IsFinal() bool {
	return p == PaymentStatusConfirmed
}

// MarshalJSON implements json.Marshaler interface
func (p PaymentStatus) MarshalJSON() ([]byte, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid payment status: %s", p)
	}
	return []byte(`"` + string(p) + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (p *PaymentStatus) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("invalid JSON format for payment status: %s", string(data))
	}

	str := string(data[1 : len(data)-1])
	paymentStatus := PaymentStatus(str)

	if !paymentStatus.IsValid() {
		return fmt.Errorf("invalid payment status: %s. Valid values are: Pending, Confirmed", str)
	}

	*p = paymentStatus
	return nil
}