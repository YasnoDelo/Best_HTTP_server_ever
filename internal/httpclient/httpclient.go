package httpclient

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type HTTPClient struct {
	client  *http.Client
	baseURL string
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		client:  &http.Client{Timeout: 10 * time.Second},
		baseURL: baseURL,
	}
}

func (c *HTTPClient) SendMetric(metricType, metricName string, value interface{}) error {
	url := fmt.Sprintf("%s/update/%s/%s/%v", c.baseURL, metricType, metricName, value)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(""))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/plain")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status: %d", resp.StatusCode)
	}

	return nil
}
