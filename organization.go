package oumlagosdk

import (
	"net/http"
	"oumla-go-sdk/types"
	"time"
)

type Organization struct {
	client *Client
}

type GetOrganizationResponse struct {
	ID string  `json:"id"`
	Name   string `json:"name"`
    Email string	`json:"email"`
}


func (s *Organization) Get() (types.DefaultResponse[GetOrganizationResponse], error) {
	req, err := s.client.NewRequest(http.MethodGet, "/organizations")
	var resp types.DefaultResponse[GetOrganizationResponse]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/* FIXME: Response is different than the one on official typescript SDK */
type GetVolumeResponse struct {
	TotalVolume int  `json:"totalVolume"`
	OrganizationId   string `json:"organizationId"`
    InsightsID *string	`json:"insightsId"`
	CreatedAt time.Time	`json:"createdAt"`
    Updated_at time.Time	`json:"updatedAt"`
}

func (s *Organization) GetVolume() (types.DefaultResponse[GetVolumeResponse], error) {
	req, err := s.client.NewRequest(http.MethodGet, "/statistics/organization/volume")
	var resp types.DefaultResponse[GetVolumeResponse]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/* FIXME: Response is different than the one on official typescript SDK */
type GetInsightsResponse struct {
	TotalVolume int  `json:"totalVolume"`
	VolumeByDay	[]struct {
		TotalVolume int  `json:"totalVolume"`
		CreatedAt time.Time	`json:"createdAt"`
	}
}

func (s *Organization) GetInsights() (types.DefaultResponse[*GetInsightsResponse], error) {
	req, err := s.client.NewRequest(http.MethodGet, "/statistics/organization/insights")
	var resp types.DefaultResponse[*GetInsightsResponse]
	if err != nil {
		return resp, err
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}