package godevill

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	httperror "github.com/eva-native/go-devill/http-error"
)

const baseUrl = "https://godville.net/gods/api/"

type Config struct {
	Token string
}

type ConfFunc func(*Config)

type API struct {
	client *http.Client
	req    *http.Request
}

func (a *API) GetWithContext(ctx context.Context) (*Superhero, error) {
	req := a.req.WithContext(ctx)
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, httperror.New(resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Superhero
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (a *API) Get() (*Superhero, error) {
	return a.GetWithContext(context.Background())
}

func WithToken(token string) ConfFunc {
	return func(cfg *Config) {
		cfg.Token = token
	}
}

func New(god string, confs ...ConfFunc) (*API, error) {
	cfg := Config{}
	for _, fn := range confs {
		fn(&cfg)
	}

	req, err := http.NewRequest(http.MethodGet, baseUrl, nil)
	if err != nil {
		return nil, err
	}
	req.URL = req.URL.JoinPath(god, cfg.Token)

	return &API{
		client: http.DefaultClient,
		req:    req,
	}, nil
}
