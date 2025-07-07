package types

import "fmt"

type WalletType string

const (
	UserType       WalletType = "User"
	MerchantType   WalletType = "Merchant"
	DepartmentType WalletType = "Department"
)

// String method for fmt.Stringer interface
func (n WalletType) String() string {
	return string(n)
}

// IsValid checks if the WalletType is valid
func (n WalletType) IsValid() bool {
	switch n {
	case UserType, MerchantType, DepartmentType:
		return true
	default:
		return false
	}
}

// MarshalJSON implements json.Marshaler interface
func (n WalletType) MarshalJSON() ([]byte, error) {
	if !n.IsValid() {
		return nil, fmt.Errorf("invalid WalletType value: %s", n)
	}
	return []byte(`"` + string(n) + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (n *WalletType) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("invalid JSON format for WalletType: %s", string(data))
	}

	str := string(data[1 : len(data)-1])
	walletType := WalletType(str)

	if !walletType.IsValid() {
		return fmt.Errorf("invalid WalletType value: %s. Valid values are: User, Merchant, Department", str)
	}

	*n = walletType
	return nil
}