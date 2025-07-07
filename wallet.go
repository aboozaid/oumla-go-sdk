package oumlagosdk

import (
	"fmt"
	"net/http"
	"oumla-go-sdk/types"
	"time"
)

type Wallet struct {
	client *Client
}

type GetWalletResponse struct {
	Reference string  `json:"reference"`
	Type   types.WalletType `json:"type"`
	Network   types.Network `json:"network"`
	OrganizationId string	`json:"organizationId"`
	CurrentDeepIndex int           `json:"currentDeepIndex"`
	Index int           `json:"index"`
    Date time.Time	`json:"date"`
}

func (s *Wallet) Get(pagination ...types.Pagination) (types.PaginatedResponse[[]GetWalletResponse], error) {
	pgn := types.Pagination{Take: 10}
	if len(pagination) > 0 {
		if pagination[0].Take != 0 {
			pgn.Take = pagination[0].Take
		}
		if pagination[0].Skip != 0 {
			pgn.Skip = pagination[0].Skip
		}
	}
	path := fmt.Sprintf("/wallets/organization?skip=%d&take=%d", pgn.Skip, pgn.Take)
	req, err := s.client.NewRequest(http.MethodGet, path)
	var resp types.PaginatedResponse[[]GetWalletResponse]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type GetWalletByReferenceResponse struct {
	Reference string  `json:"reference"`
	Type   types.WalletType `json:"type"`
	// OrganizationId string	`json:"organizationId"`
    Label *string	`json:"label,omitempty"`
	IsActive bool	`json:"isActive"`
    IsVerified bool	`json:"isVerified"`
	CurrentIndex int           `json:"currentIndex"`
    MiniWallets []types.Wallet	`json:"miniWallets"`
	CreatedAt time.Time	`json:"createdAt"`
    Updated_at time.Time	`json:"updated_at"`
}

func (s *Wallet) GetByReference(reference string, pagination ...types.Pagination) (types.PaginatedResponse[[]GetWalletByReferenceResponse], error) {
	pgn := types.Pagination{Take: 10}
	if len(pagination) > 0 {
		if pagination[0].Take != 0 {
			pgn.Take = pagination[0].Take
		}
		if pagination[0].Skip != 0 {
			pgn.Skip = pagination[0].Skip
		}
	}
	path := fmt.Sprintf("/wallets/profile/%s?skip=%d&take=%d", reference, pgn.Skip, pgn.Take)
	req, err := s.client.NewRequest(http.MethodGet, path)
	var resp types.PaginatedResponse[[]GetWalletByReferenceResponse]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type GenerateWalletResponse struct {
	Reference string	`json:"reference"`
	Network	types.Network	`json:"network"`
	Type	string	`json:"type"`
	Date	time.Time	`json:"date"`
}

func (s *Wallet) Generate(network types.Network, reference string) (types.DefaultResponse[GenerateWalletResponse], error) {
	req, err := s.client.NewRequest(http.MethodPost, "/wallets/generate", map[string]interface{}{
		"network": network,
		"reference": reference,
	})
	var resp types.DefaultResponse[GenerateWalletResponse]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
