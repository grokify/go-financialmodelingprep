package financialmodelingprep

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/grokify/mogo/net/http/httpsimple"
	"github.com/grokify/mogo/net/urlutil"
)

type Client struct {
	*httpsimple.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		Client: &httpsimple.Client{
			BaseURL: ServerURL,
			Query: url.Values{
				QueryParamAPIKey: []string{apiKey},
			},
		},
	}
}

func (c *Client) GetIncomeStatement(symbol, period string, limit uint) (*http.Response, error) {
	period = strings.ToLower(strings.TrimSpace(period))
	if c.Client == nil {
		return nil, errors.New("simpleclient not set")
	} else if strings.TrimSpace(symbol) == "" {
		return nil, errors.New("symbol must be set")
	} else if period == "" || (period != QueryParamPeriodValueAnnual && period != QueryParamPeriodValueQuarter) {
		return nil, errors.New("period must be set to `annual` or `quarter`")
	} else if limit == 0 {
		limit = QueryParamLimitExample
	}
	req := httpsimple.Request{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(urlutil.JoinAbsolute(ServerURL, IncomeStatementURLFormat), symbol),
		Query: url.Values{
			QueryParamPeriod: []string{period},
			QueryParamLimit:  []string{strconv.Itoa(int(limit))},
		},
	}
	return c.Client.Do(req)
}
