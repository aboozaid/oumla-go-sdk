package oumlagosdk

import (
	"fmt"
	"net/http"
	"oumla-go-sdk/types"
)

type Transaction struct {
	client *Client
}

func (s *Transaction) GetByOrganization(pagination ...types.Pagination) (types.PaginatedResponse[[]Transaction], error) {
	resp, err := s.get("/transactions/organization", pagination...)
	return resp, err
}

func (s *Transaction) GetByAddress(address string, pagination ...types.Pagination) (types.PaginatedResponse[[]Transaction], error) {
	resp, err := s.get(fmt.Sprintf("/transactions/address/%s", address), pagination...)
	return resp, err
}

func (s *Transaction) GetByWallet(wallet string, pagination ...types.Pagination) (types.PaginatedResponse[[]Transaction], error) {
	resp, err := s.get(fmt.Sprintf("/transactions/wallet/%s", wallet), pagination...)
	return resp, err
}

func (s *Transaction) GetByProfile(profile string, pagination ...types.Pagination) (types.PaginatedResponse[[]Transaction], error) {
	resp, err := s.get(fmt.Sprintf("/transactions/profile/%s", profile), pagination...)
	return resp, err
}

func (s *Transaction) get(query string, pagination ...types.Pagination) (types.PaginatedResponse[[]Transaction], error) {
	pgn := types.Pagination{Take: 10}
	if len(pagination) > 0 {
		if pagination[0].Take != 0 {
			pgn.Take = pagination[0].Take
		}
		if pagination[0].Skip != 0 {
			pgn.Skip = pagination[0].Skip
		}
	}
	path := fmt.Sprintf("%s?skip=%d&take=%d", query, pgn.Skip, pgn.Take)
	req, err := s.client.NewRequest(http.MethodGet, path)
	var resp types.PaginatedResponse[[]Transaction]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type CreateTransaction struct {
	Network     types.Network
	ClientShare string
	To          string
	Amount      string
	From        []string
}

type CreateTransactionResponse struct {
	ID string	`json:"id"`
	Status	string	`json:"status"`
}

func (s *Transaction) Create(transaction CreateTransaction) (types.DefaultResponse[CreateTransactionResponse], error) {
	req, err := s.client.NewRequest(http.MethodPost, "/withdraw/address", map[string]interface{}{
		"network": transaction.Network,
		"clientShare": transaction.ClientShare,
		"to": transaction.To,
		"amount": transaction.Amount,
		"From": transaction.From,
	})
	var resp types.DefaultResponse[CreateTransactionResponse]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
