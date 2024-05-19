package podcastindex

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Config holds the configuration for the API client
type Config struct {
	BaseURL   string
	UserAgent string
}

// DefaultConfig is used when NewClient is used to create an API client
var DefaultConfig *Config = &Config{
	BaseURL:   BaseURL,
	UserAgent: UserAgent,
}

// Client connects to the podcastindex API
type Client struct {
	config *Config
	client *http.Client
	key    string
	secret string
}

// NewClient creates an API client with the default configuration
func NewClient(apiKey, apiSecret string) *Client {
	return NewClientWithConfig(apiKey, apiSecret, *DefaultConfig, http.DefaultClient)
}

// NewClientWithConfig creates an API client with an custom configuration
func NewClientWithConfig(apiKey, apiSecret string, config Config, client *http.Client) *Client {
	return &Client{
		key:    apiKey,
		secret: apiSecret,
		config: &config,
		client: client,
	}
}

func (c *Client) request(url string, result interface{}) error {
	u := fmt.Sprintf("%s%s", c.config.BaseURL, url)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	now := time.Now()
	auth := generateAuthorizationHeader(c.key, c.secret, now)
	req.Header.Set("User-Agent", c.config.UserAgent)
	req.Header.Set("X-Auth-Date", fmt.Sprintf("%d", now.Unix()))
	req.Header.Set("X-Auth-Key", c.key)
	req.Header.Set("Authorization", auth)

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	if res.Body == nil {
		return errors.New("API didn't returned a response")
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return decode(resBody, result)
}

func decode(in []byte, out interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(in))
	return decoder.Decode(out)
}
