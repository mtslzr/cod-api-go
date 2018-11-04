package cod

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// API holds configuration variables for accessing the API.
type API struct {
	BaseURL  *url.URL
	Game     string
	Platform string
	UserName string
}

// Validation returns a user ID and status for game/platform/username comvbination.
type Validation struct {
	ID       int
	Success  bool
	UserName string
}

// New creates a new API client.
func New(game string, platform string, username string) (*API, error) {
	base, err := url.Parse("https://cod-api.theapinetwork.com/api/")
	if err != nil {
		return &API{}, err
	}

	return &API{
		BaseURL:  base,
		Game:     game,
		Platform: platform,
		UserName: username,
	}, nil
}

// NewRequest creates the GET request to access the API.
func (a *API) NewRequest(endpoint string) (*http.Request, error) {
	end, err := url.Parse(endpoint)
	if err != nil {
		return &http.Request{}, err
	}
	urlStr := a.BaseURL.ResolveReference(end)

	req, err := http.NewRequest("GET", urlStr.String(), nil)
	if err != nil {
		return req, err
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}

// Do sends out a request to the API and unmarshals the data.
func (a *API) Do(req *http.Request, i interface{}) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &i)
}

// ValidateUser checks if game/user/platform combination exists.
func (a *API) ValidateUser() (*Validation, error) {
	endpoint := "validate/" + a.Game + "/" + url.QueryEscape(a.UserName) + "/" + a.Platform
	req, err := a.NewRequest(endpoint)

	if err != nil {
		return &Validation{}, err
	}

	var validUser Validation
	err = a.Do(req, &validUser)

	return &validUser, err
}
