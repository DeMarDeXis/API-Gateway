package users

import (
	"ApiGateway/internal/models/auth"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Client struct {
	baseURL string
	client  *http.Client
	log     *slog.Logger
	methods map[string]string
}

func New(baseURL string, log *slog.Logger) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
		log:     log,
		methods: map[string]string{
			"sign-up": "/users/sign-up",
			"sign-in": "/users/sign-in",
		},
	}
}

func (c *Client) SignUp(input auth.InputSignUp) (int, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal input: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+c.methods["sign-up"], bytes.NewReader(body))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	var response struct {
		ID int `json:"id"`
	}

	if err := c.doRequest(req, &response); err != nil {
		return 0, err
	}

	return response.ID, nil
}

func (c *Client) SignIn(input auth.InputSignIn) (int, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal input: %w", err)
	}

	req, err := http.NewRequest("GET", c.baseURL+c.methods["sign-in"], bytes.NewReader(body))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	var response struct {
		ID int `json:"id"`
	}

	if err := c.doRequest(req, &response); err != nil {
		return 0, err
	}

	return response.ID, nil
}

func (c *Client) doRequest(req *http.Request, response interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
