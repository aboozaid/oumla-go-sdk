package types

import "time"

type Wallet struct {
	OrganizationId   string        `json:"organizationId"`
	Reference        string        `json:"reference"`
	Type             WalletType    `json:"type"`
	Index            int           `json:"index"`
	CurrentDeepIndex int           `json:"currentDeepIndex"`
	Addresses        []Address         `json:"address"`
	Transactions     []Transaction `json:"transaction"`
	WalletCoinType   Network       `json:"network"`
	Date             time.Time     `json:"date"`
}