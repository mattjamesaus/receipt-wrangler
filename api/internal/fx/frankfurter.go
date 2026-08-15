package fx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"receipt-wrangler/api/internal/constants"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

const (
	defaultBaseURL  = "https://api.frankfurter.dev"
	defaultProvider = "ECB"
)

type Quote struct {
	Rate          decimal.Decimal
	EffectiveDate time.Time
	Provider      string
	RetrievedAt   time.Time
}

// FrankfurterProvider retrieves a single historical reference rate. The
// concrete client and base URL are injectable so failure paths remain hermetic
// in tests.
type FrankfurterProvider struct {
	Client       *http.Client
	BaseURL      string
	RateProvider string
	Now          func() time.Time
}

func NewFrankfurterProvider() FrankfurterProvider {
	baseURL := strings.TrimRight(os.Getenv(string(constants.FxProviderBaseURL)), "/")
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	provider := strings.TrimSpace(os.Getenv(string(constants.FxRateProvider)))
	if provider == "" {
		provider = defaultProvider
	}
	return FrankfurterProvider{
		Client:       &http.Client{Timeout: 5 * time.Second},
		BaseURL:      baseURL,
		RateProvider: provider,
		Now:          time.Now,
	}
}

func (provider FrankfurterProvider) Name() string {
	return "Frankfurter:" + provider.RateProvider
}

func (provider FrankfurterProvider) HistoricalRate(
	ctx context.Context,
	baseCurrency string,
	quoteCurrency string,
	receiptDate time.Time,
) (Quote, error) {
	if provider.Client == nil {
		return Quote{}, fmt.Errorf("FX HTTP client is not configured")
	}

	endpoint, err := url.Parse(provider.BaseURL + "/v2/rate/" + url.PathEscape(baseCurrency) + "/" + url.PathEscape(quoteCurrency))
	if err != nil {
		return Quote{}, fmt.Errorf("build FX request: %w", err)
	}
	query := endpoint.Query()
	// A receipt date is a calendar date, not an instant. Preserve the date in
	// the offset supplied by the client so midnight in a positive UTC offset is
	// not accidentally looked up as the previous day.
	query.Set("date", receiptDate.Format("2006-01-02"))
	if provider.RateProvider != "" {
		query.Set("providers", provider.RateProvider)
	}
	endpoint.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return Quote{}, fmt.Errorf("build FX request: %w", err)
	}
	request.Header.Set("Accept", "application/json")

	response, err := provider.Client.Do(request)
	if err != nil {
		return Quote{}, fmt.Errorf("retrieve historical FX rate: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return Quote{}, fmt.Errorf("retrieve historical FX rate: provider returned HTTP %d", response.StatusCode)
	}

	var body struct {
		Date  string          `json:"date"`
		Base  string          `json:"base"`
		Quote string          `json:"quote"`
		Rate  decimal.Decimal `json:"rate"`
	}
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return Quote{}, fmt.Errorf("decode historical FX rate: %w", err)
	}
	if body.Base != baseCurrency || body.Quote != quoteCurrency || !body.Rate.IsPositive() {
		return Quote{}, fmt.Errorf("historical FX response did not contain a valid %s/%s rate", baseCurrency, quoteCurrency)
	}
	effectiveDate, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		return Quote{}, fmt.Errorf("decode historical FX date: %w", err)
	}

	now := time.Now
	if provider.Now != nil {
		now = provider.Now
	}
	return Quote{
		Rate:          body.Rate,
		EffectiveDate: effectiveDate,
		Provider:      provider.Name(),
		RetrievedAt:   now().UTC(),
	}, nil
}
