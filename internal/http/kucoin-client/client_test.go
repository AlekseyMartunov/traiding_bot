package httpclient

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	kucoinentity "tradingbot/internal/entity"
)

func mockHeaderFunc(method, url, body, secret, passPhrase, key, version string, now time.Time) map[string]string {
	return map[string]string{"1": "1"}
}

type mockLogger struct{}

func (m mockLogger) Debug(msg string, args ...any) {}
func (m mockLogger) Info(msg string, args ...any)  {}
func (m mockLogger) Warn(msg string, args ...any)  {}
func (m mockLogger) Error(msg string, args ...any) {}

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

func TestAccountInfo(t *testing.T) {
	sever := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

	}))
	defer sever.Close()

	testCases := []struct {
		name           string
		data           []accountInfoJSON
		expectedResult []*kucoinentity.AccountInfo
		expectedErr    error
	}{
		{
			name: "test-1",
			data: []accountInfoJSON{{
				Id:        "123",
				Currency:  "BTC-USD",
				Type:      "Trade",
				Balance:   "200.3",
				Available: "200.3",
				Holds:     "0",
			}, {
				Id:        "124",
				Currency:  "BTC-SOL",
				Type:      "Trade",
				Balance:   "101.1",
				Available: "101.1",
				Holds:     "0.5",
			}},
			expectedResult: []*kucoinentity.AccountInfo{{
				ID:        "123",
				Currency:  "BTC-USDT",
				Type:      "Trade",
				Balance:   200.3,
				Available: 200.3,
				Holds:     0,
			}, {
				ID:        "124",
				Currency:  "BTC-SOL",
				Type:      "Trade",
				Balance:   101.1,
				Available: 101.1,
				Holds:     0.5,
			}},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				if request.URL.Path != accountEndpoint {
					writer.Write([]byte(""))
					return
				}

				if request.Method != http.MethodGet {
					writer.Write([]byte(""))
					return
				}

				b, err := json.Marshal(tc.data)
				assert.NoError(t, err)
				writer.Write(b)
			}))

			defer sever.Close()

			cfg := mockConfig{baseEndpoint: server.URL}
			httpClient := New(mockLogger{}, cfg, mockHeaderFunc)

			result, err := httpClient.AccountInfo()

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedErr, err)

		})
	}

}
