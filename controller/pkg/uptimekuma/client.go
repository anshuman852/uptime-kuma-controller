// Package uptimekuma provides a minimal client for interacting with the Uptime Kuma API.
package uptimekuma

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents a minimal Uptime Kuma API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	APIToken   string // Optional: for future authentication support
}

// NewClient creates a new Uptime Kuma API client.
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// GetMonitorStatus retrieves the status of a monitor from Uptime Kuma.
// For now, this uses a placeholder endpoint and returns the raw response body.
func (c *Client) GetMonitorStatus(monitorID string) (string, error) {
	url := fmt.Sprintf("%s/api/monitors/%s/status", c.BaseURL, monitorID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Optional: Add authentication header if needed in the future
	// req.Header.Set("Authorization", "Bearer "+c.APIToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return string(body), nil
}