package httpclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YasnoDelo/Best_HTTP_server_ever/internal/httpclient"
)

func TestSendMetric(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := httpclient.NewHTTPClient(server.URL)

	err := client.SendMetric("gauge", "Alloc", 123.45)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
