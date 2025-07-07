package oumlagosdk

import (
	"net/http"
	"oumla-go-sdk/types"
)

type Address struct {
	client *Client
}

func (s *Address) Generate(network types.Network, reference string, clientShare string) (types.DefaultResponse[string], error) {
	req, err := s.client.NewRequest(http.MethodPost, "/address/generate", map[string]interface{}{
		"network":   network,
		"reference": reference,
		"clientShare": clientShare,
	})
	var resp types.DefaultResponse[string]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}