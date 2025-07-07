package oumlagosdk

import (
	"fmt"
	"net/http"
	"oumla-go-sdk/types"
	"time"
)

type Profile struct {
	client *Client
}

type GetProfileResponse struct {
	Reference string  `json:"reference"`
	Type   types.WalletType `json:"type"`
    OrganizationId string	`json:"organizationId"`
    Label *string	`json:"label,omitempty"`
    MiniWallets []types.Wallet	`json:"miniWallets"`
    AddressesCount int	`json:"addressesCount"`
	CreatedAt time.Time	`json:"createdAt"`
    Updated_at time.Time	`json:"updated_at"`
}

func (s *Profile) Get(pagination ...types.Pagination) (types.PaginatedResponse[[]GetProfileResponse], error) {
	pgn := types.Pagination{Take: 10}
	if len(pagination) > 0 {
		if pagination[0].Take != 0 {
			pgn.Take = pagination[0].Take
		}
		if pagination[0].Skip != 0 {
			pgn.Skip = pagination[0].Skip
		}
	}
	path := fmt.Sprintf("/profiles?skip=%d&take=%d", pgn.Skip, pgn.Take)
	req, err := s.client.NewRequest(http.MethodGet, path)
	var resp types.PaginatedResponse[[]GetProfileResponse]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type CreateProfileResponse struct {
	Reference string  `json:"reference"`
	Type   types.WalletType `json:"type"`
}

func (s *Profile) Create(walletType types.WalletType, reference string) (types.DefaultResponse[CreateProfileResponse], error) {
	req, err := s.client.NewRequest(http.MethodPost, "/profiles", map[string]interface{}{
		"type":   walletType,
		"reference": reference,
	})
	var resp types.DefaultResponse[CreateProfileResponse]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}