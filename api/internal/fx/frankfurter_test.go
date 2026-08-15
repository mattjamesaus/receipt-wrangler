package fx

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestHistoricalRateReturnsProvenance(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/rate/USD/AUD" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("date") != "2025-06-14" || r.URL.Query().Get("providers") != "ECB" {
			t.Errorf("query = %q", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"date":"2025-06-13","base":"USD","quote":"AUD","rate":1.542}`))
	}))
	defer server.Close()

	now := time.Date(2025, 6, 15, 1, 2, 3, 0, time.UTC)
	provider := FrankfurterProvider{
		Client: server.Client(), BaseURL: server.URL, RateProvider: "ECB", Now: func() time.Time { return now },
	}
	quote, err := provider.HistoricalRate(context.Background(), "USD", "AUD", time.Date(2025, 6, 14, 0, 0, 0, 0, time.FixedZone("AWST", 8*60*60)))
	if err != nil {
		t.Fatal(err)
	}
	if !quote.Rate.Equal(decimal.RequireFromString("1.542")) {
		t.Errorf("rate = %s", quote.Rate)
	}
	if quote.EffectiveDate.Format("2006-01-02") != "2025-06-13" || quote.Provider != "Frankfurter:ECB" || !quote.RetrievedAt.Equal(now) {
		t.Errorf("unexpected provenance: %#v", quote)
	}
}

func TestHistoricalRateRejectsProviderFailures(t *testing.T) {
	tests := []struct {
		name   string
		status int
		body   string
	}{
		{name: "http error", status: http.StatusBadGateway, body: `{}`},
		{name: "invalid json", status: http.StatusOK, body: `{`},
		{name: "wrong pair", status: http.StatusOK, body: `{"date":"2025-06-13","base":"EUR","quote":"AUD","rate":1.5}`},
		{name: "zero rate", status: http.StatusOK, body: `{"date":"2025-06-13","base":"USD","quote":"AUD","rate":0}`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			}))
			defer server.Close()
			provider := FrankfurterProvider{Client: server.Client(), BaseURL: server.URL, RateProvider: "ECB"}
			if _, err := provider.HistoricalRate(context.Background(), "USD", "AUD", time.Now()); err == nil {
				t.Fatal("expected an error")
			}
		})
	}
}
