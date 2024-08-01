package kucoinheader

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateSecretsHeaders(t *testing.T) {
	testCases := []struct {
		name       string
		method     string
		url        string
		body       string
		secret     string
		passPhrase string
		key        string
		version    string

		expected map[string]string
	}{
		{
			name:       "test-1",
			method:     http.MethodGet,
			url:        "www.some-test.com",
			body:       `{"currency": "BTC-ETH", "price": "200"}`,
			secret:     "83sfk34sHKS534D75jsdFsJSFds",
			passPhrase: "some-secret-phrase",
			key:        "438sd3fr4dFfk&Ds87kd04sfJH4SD",
			version:    "2",
			expected: map[string]string{
				"KC-API-KEY":         "438sd3fr4dFfk&Ds87kd04sfJH4SD",
				"KC-API-SIGN":        "MWLDlkt+QJ8keQIt9ouWbJFmp5Ghc8LWCMkCfjZEJHU=",
				"KC-API-TIMESTAMP":   "1601536844010",
				"KC-API-PASSPHRASE":  "NuIc5BUPYs8wsX1HElYvxf24s4jV8f+7X9IjLXJABsY=",
				"KC-API-KEY-VERSION": "2",
				"Content-Type":       "application/json",
			},
		},

		{
			name:       "test-2",
			method:     http.MethodPost,
			url:        "www.another-test.com",
			body:       `{"currency": "BTC-ETH", "price": "200"}`,
			secret:     "848s334D7sdF7JSF3sfk5jS534DsHKS54dFKds",
			passPhrase: "another-secret-phrase",
			key:        "438&Ds87kd3fr4dFfk&7&Ds87kdd3fr4dFfk&04sfJH4SD",
			version:    "1",
			expected: map[string]string{
				"KC-API-KEY":         "438&Ds87kd3fr4dFfk&7&Ds87kdd3fr4dFfk&04sfJH4SD",
				"KC-API-SIGN":        "pSfYi8z564eNW0kHeFzw7OMso3dVCQ0hWriPChdoebM=",
				"KC-API-TIMESTAMP":   "1601536844010",
				"KC-API-PASSPHRASE":  "36OputLn2o6bhZ9jBPcMjbhn3M/mlSgooj7OkB59iDM=",
				"KC-API-KEY-VERSION": "1",
				"Content-Type":       "application/json",
			},
		},

		{
			name:       "test-3",
			method:     http.MethodDelete,
			url:        "www.another-test/info-account.com",
			body:       `{"currency": ["BTC-USDT", "SOL-ETH"], "price": "300"}`,
			secret:     "8423jS5F7ksf-jJSFs45dF7-JSF36s8s33sf4D7sd4fk5dFKds",
			passPhrase: "very-secret-phrase",
			key:        "438&Dsk784jsdDsd74fk87kd3fr4dFf23sd&04sfJH4SD",
			version:    "3",
			expected: map[string]string{
				"KC-API-KEY":         "438&Dsk784jsdDsd74fk87kd3fr4dFf23sd&04sfJH4SD",
				"KC-API-SIGN":        "Mj5yXKJzWEf9XZKIvfb/Ol2PcGTEEnIkyZ4n/64Nltk=",
				"KC-API-TIMESTAMP":   "1601536844010",
				"KC-API-PASSPHRASE":  "tIeUnn9YCw6L/MtKneYj7e4/90+b3xhivrHV6NbffLE=",
				"KC-API-KEY-VERSION": "3",
				"Content-Type":       "application/json",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CreateSecretsHeaders(
				tc.method,
				tc.url,
				tc.body,
				tc.secret,
				tc.passPhrase,
				tc.key,
				tc.version,
				time.Date(2020, 10, 1, 7, 20, 43, 10000000, time.UTC),
			)

			assert.Equal(t, tc.expected, result)
		})
	}
}
