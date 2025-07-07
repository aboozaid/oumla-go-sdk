package types

import "fmt"

type Network string

const (
	NetworkBTC  Network = "BTC"
	NetworkETH  Network = "ETH"
	NetworkTBTC Network = "tBTC"
	NetworkTETH Network = "tETH"
)

// String method for fmt.Stringer interface
func (n Network) String() string {
	return string(n)
}

// IsValid checks if the network is valid
func (n Network) IsValid() bool {
	switch n {
	case NetworkBTC, NetworkETH, NetworkTBTC, NetworkTETH:
		return true
	default:
		return false
	}
}

// IsTestnet returns true if this is a testnet network
func (n Network) IsTestnet() bool {
	return n == NetworkTBTC || n == NetworkTETH
}

// MarshalJSON implements json.Marshaler interface
func (n Network) MarshalJSON() ([]byte, error) {
	if !n.IsValid() {
		return nil, fmt.Errorf("invalid network value: %s", n)
	}
	return []byte(`"` + string(n) + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (n *Network) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("invalid JSON format for network: %s", string(data))
	}

	str := string(data[1 : len(data)-1])
	network := Network(str)

	if !network.IsValid() {
		return fmt.Errorf("invalid network value: %s. Valid values are: BTC, ETH, tBTC, tETH", str)
	}

	*n = network
	return nil
}