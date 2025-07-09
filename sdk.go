package oumlagosdk

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	BaseURL string
	ApiKey  string
	// Env     string
}

type OumlaSdk interface{}

type Sdk struct {
	client *Client

	Profiles	*Profile
	Wallets *Wallet
	Addresses	*Address
	Organizations	*Organization
}

const (
	SDK_VERSION           = "1.0.0"
	// CHECK_UPDATE_INTERVAL = 1 * time.Hour // every 1h
)

func NewOumla(config Config) Sdk {
	if config.BaseURL == "" {
		// NOTE: added api version once on the url
		config.BaseURL = "https://api.oumla.com/api/v1"
	}
	if config.ApiKey == "" {
		panic(errors.New("failed to initialize the sdk, apiKey is required"))
	}

	client, err := NewClient(config.BaseURL, config.ApiKey, map[string]string{
		"x-sdk-version": SDK_VERSION,
	})

	if err != nil {
		panic(fmt.Errorf("failed to create the client: %s", err.Error()))
	}

	s := Sdk{
		client: client,
		Profiles: &Profile{
			client: client,
		},
		Wallets: &Wallet{
			client: client,
		},
		Addresses: &Address{
			client: client,
		},
		Organizations: &Organization{
			client: client,
		},
	}
	latestVersion, err := s.checkForUpdates()
	if err != nil {
		panic(fmt.Errorf("failed to check for updated: %s", err.Error()))
	}
	if latestVersion != SDK_VERSION {
		log.Printf("A new version %s of the Oumla SDK is available. Plaese update.", latestVersion)
	}
	return s
}

func (s *Sdk) checkForUpdates() (string, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/sdk-version")
	if err != nil {
		return "", err
	}
	var resp struct {
		LatestVersion string	`json:"latestVersion"`
	}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return "", err
	}
	return resp.LatestVersion, nil
}