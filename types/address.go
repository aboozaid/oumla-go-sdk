package types

import "time"

type Address struct {
	Address      string        `json:"address"`
	Index        int           `json:"index"`
	Label        *string       `json:"label,omitempty"`
	Transactions []Transaction `json:"Transaction"`
	Tags         []Tag        `json:"Tags"`
	Network      Network       `json:"network"`
	Balance      int           `json:"balance"`
	Reference    string        `json:"reference"`
	Date         time.Time     `json:"date"`
}

type Tag struct {
	Name string	`json:"name"`
    Addresses []Address	`json:"address"`
}