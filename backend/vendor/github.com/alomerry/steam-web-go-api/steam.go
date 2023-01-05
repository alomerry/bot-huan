package steam

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	STEAM_WEB_API_DOMAIN = "https://api.steampowered.com"
)

type SteamClient struct {
	client *http.Client
	key    string
}

type _interface struct {
	sc *SteamClient
}

func NewClient(token string) *SteamClient {
	return &SteamClient{
		key:    token,
		client: &http.Client{},
	}
}

func (s *SteamClient) GetIUserStats() *UserStatsInterface {
	return &UserStatsInterface{sc: s}
}

func (s *SteamClient) GetIUser() *UserInterface {
	return &UserInterface{sc: s}
}

func (i *_interface) execute(ctx context.Context, endpoint string, response interface{}) error {
	endpoint = fmt.Sprintf("%s/%s", STEAM_WEB_API_DOMAIN, endpoint)
	r, err := i.sc.client.Get(endpoint)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, response)
	if err != nil {
		return err
	}
	return nil
}
