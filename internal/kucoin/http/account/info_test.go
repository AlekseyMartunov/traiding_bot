package kucoinaccount

import (
	"net/http"
	"net/http/httptest"
	"testing"

	kucoinentity "tradingbot/internal/kucoin/entity"
	kucoinerrors "tradingbot/internal/kucoin/errors"

	"github.com/stretchr/testify/assert"
)

type mockLogger struct{}

func (m mockLogger) Trace(message string) {}
func (m mockLogger) Debug(message string) {}
func (m mockLogger) Info(message string)  {}
func (m mockLogger) Warn(message string)  {}
func (m mockLogger) Error(message string) {}

type mockConfig struct {
	key          string
	secret       string
	passPhrase   string
	version      string
	baseEndpoint string
}

func (c mockConfig) Key() string          { return c.key }
func (c mockConfig) Secret() string       { return c.secret }
func (c mockConfig) PassPhrase() string   { return c.passPhrase }
func (c mockConfig) Version() string      { return c.version }
func (c mockConfig) BaseEndpoint() string { return c.baseEndpoint }

func TestAccountManager_GetAccountInfo(t *testing.T) {
	cfg := mockConfig{
		key:        "123",
		secret:     "123",
		passPhrase: "123",
		version:    "2",
	}

	testCases := []struct {
		name        string
		description string
		result      []*kucoinentity.AccountInfo
		err         error
		handlerFunc func(http.ResponseWriter, *http.Request)
	}{
		{
			name:        "test-1",
			description: "successful test",
			result: []*kucoinentity.AccountInfo{{ID: "123", Currency: "BTC-USDT", Type: "main", Balance: 200.0, Available: 200.0, Holds: 0.0},
				{ID: "124", Currency: "ETH-USDT", Type: "trade", Balance: 300.0, Available: 300.0, Holds: 0.0}},
			err: nil,
			handlerFunc: func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte(
					`{"code":"200000","data": 
					 [{"id": "123", "currency": "BTC-USDT", "type": "main", "balance": "200.0","available":"200.0", "holds": "0.0"},
					  {"id": "124", "currency": "ETH-USDT", "type": "trade", "balance": "300.0","available":"300.0", "holds": "0.0"}
					]`))
			},
		},
		{
			name:        "test-2",
			description: "check  secret key in header",
			result: []*kucoinentity.AccountInfo{
				{ID: "124", Currency: "ETH-USDT", Type: "trade", Balance: 300.0, Available: 300.0, Holds: 0.0},
			},
			err: nil,
			handlerFunc: func(writer http.ResponseWriter, request *http.Request) {
				if request.Header.Get("KC-API-KEY") != "123" {
					writer.WriteHeader(http.StatusUnauthorized)
					return
				}
				writer.WriteHeader(http.StatusOK)
				writer.Write(
					[]byte(
						`{"code":"200000","data":
					[{"id": "124", "currency": "ETH-USDT", "type": "trade", "balance": "300.0","available":"300.0", "holds": "0.0"}]
					}`))
			},
		},
		{
			name:        "test-3",
			description: "parsing json error",
			result:      nil,
			err:         kucoinerrors.ErrUnmarshal,
			handlerFunc: func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte("ode\":\"2000"))
			},
		},
		{
			name:        "test-4",
			description: "recasting to dto error",
			result:      nil,
			err:         kucoinerrors.ErrRecastDTO,
			handlerFunc: func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusOK)
				writer.Write(
					[]byte(
						`{"code":"200000","data":
					[{"id": "124", "currency": "ETH-USDT", "type": "trade", "balance": "here should be a number","available":"300.0", "holds": "0.0"}]
					}`))
			},
		},
		{
			name:        "test-5",
			description: "check request method (should be GET)",
			result: []*kucoinentity.AccountInfo{
				{ID: "124", Currency: "ETH-USDT", Type: "trade", Balance: 300.0, Available: 300.0, Holds: 0.0},
			},
			err: nil,
			handlerFunc: func(writer http.ResponseWriter, request *http.Request) {
				if request.Method != http.MethodGet {
					writer.WriteHeader(http.StatusBadRequest)
					return
				}
				writer.WriteHeader(http.StatusOK)
				writer.Write(
					[]byte(
						`{"code":"200000","data":
					[{"id": "124", "currency": "ETH-USDT", "type": "trade", "balance": "300.0","available":"300.0", "holds": "0.0"}]
					}`))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tc.handlerFunc))
			defer server.Close()

			cfg.baseEndpoint = server.URL

			accountManager := New(mockLogger{}, cfg)
			result, err := accountManager.GetAccountInfo()

			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.result, result)
		})
	}

}
