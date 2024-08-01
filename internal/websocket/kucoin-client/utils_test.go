package wsclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
	kucoinentity "tradingbot/internal/entity"
)

func TestIn(t *testing.T) {
	testCases := []struct {
		name     string
		prefix   string
		topics   []string
		expected bool
	}{
		{
			name:     "test-1",
			prefix:   "awesome GO",
			topics:   []string{"b", "a", "c"},
			expected: true,
		},
		{
			name:     "test-2",
			prefix:   "aB",
			topics:   []string{"b", "a", "c"},
			expected: true,
		},
		{
			name:     "test-3",
			prefix:   "cccCCC",
			topics:   []string{"bbb", "dfj", "ccc"},
			expected: true,
		},
		{
			name:     "test-4",
			prefix:   "bigLong",
			topics:   []string{"bbb", "AAA", "ccc"},
			expected: false,
		},
		{
			name:     "test-5",
			prefix:   "bAbbb",
			topics:   []string{"bbb", "bca", "ccc"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := in(tc.prefix, tc.topics)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestAllTypesOfMessagesInGroup(t *testing.T) {
	testCases := []struct {
		name     string
		group    []string
		messages []subscribeMessage
		expected bool
	}{
		{
			name:  "test-1",
			group: allPublicTopics,
			messages: []subscribeMessage{
				NewTickerSubscribeMessages("BTC-USDT, ETH-USDT"),
				NewTickerUnSubscribeMessages("PEPE-USDT, SOL-USDT"),
			},
			expected: true,
		},

		{
			name:  "test-2",
			group: allPrivetTopics,
			messages: []subscribeMessage{
				NewTickerSubscribeMessages("BTC-USDT", "ETH-USDT"),
				NewTickerUnSubscribeMessages("PEPE-USDT", "SOL-USDT"),
			},
			expected: false,
		},

		{
			name:  "test-3",
			group: allPrivetTopics,
			messages: []subscribeMessage{
				NewCandleSubscribeMessages("BTC-USDT", kucoinentity.Min15),
				NewCandleSubscribeMessages("PEPE-USDT", kucoinentity.Min5),
			},
			expected: false,
		},

		{
			name:  "test-4",
			group: allPublicTopics,
			messages: []subscribeMessage{
				NewCandleSubscribeMessages("BTC-USDT", kucoinentity.Min15),
				NewCandleUnSubscribeMessages("PEPE-USDT", kucoinentity.Min5),
			},
			expected: true,
		},

		{
			name:  "test-5",
			group: allPublicTopics,
			messages: []subscribeMessage{
				NewOrderChangeMessageSubscribe(),
				NewOrderChangeMessageUnSubscribe(),
			},
			expected: false,
		},

		{
			name:  "test-6",
			group: allPrivetTopics,
			messages: []subscribeMessage{
				NewOrderChangeMessageSubscribe(),
				NewOrderChangeMessageUnSubscribe(),
			},
			expected: true,
		},

		{
			name:  "test-7",
			group: allPrivetTopics,
			messages: []subscribeMessage{
				NewOrderChangeV2MessageSubscribe(),
				NewOrderChangeV2MessageUnSubscribe(),
			},
			expected: true,
		},

		{
			name:  "test-8",
			group: allPublicTopics,
			messages: []subscribeMessage{
				NewOrderChangeV2MessageSubscribe(),
				NewOrderChangeV2MessageUnSubscribe(),
			},
			expected: false,
		},

		{
			name:  "test-9",
			group: allPrivetTopics,
			messages: []subscribeMessage{
				NewAccountBalanceChangeMessageUnSubscribe(),
				NewAccountBalanceChangeMessageSubscribe(),
			},
			expected: true,
		},

		{
			name:  "test-10",
			group: allPublicTopics,
			messages: []subscribeMessage{
				NewAccountBalanceChangeMessageUnSubscribe(),
				NewAccountBalanceChangeMessageSubscribe(),
			},
			expected: false,
		},

		{
			name:  "test-11",
			group: allPrivetTopics,
			messages: []subscribeMessage{
				NewStopOrderEventMessageSubscribe(),
				NewStopOrderEventMessageUnSubscribe(),
			},
			expected: true,
		},

		{
			name:  "test-12",
			group: allPublicTopics,
			messages: []subscribeMessage{
				NewStopOrderEventMessageSubscribe(),
				NewStopOrderEventMessageUnSubscribe(),
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := allTypesOfMessagesInGroup(tc.group, tc.messages...)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestFieldsSubscribeMessages(t *testing.T) {
	testCases := []struct {
		name                  string
		messageFunc           func() subscribeMessage
		expectedType          string
		expectedTopic         string
		expectedPrivetChannel bool
		expectedResponse      bool
	}{
		{
			name:                  "test-1",
			messageFunc:           NewOrderChangeMessageSubscribe,
			expectedType:          SubscribeMessage,
			expectedTopic:         privetOrderChange,
			expectedPrivetChannel: true,
			expectedResponse:      true,
		},

		{
			name:                  "test-2",
			messageFunc:           NewOrderChangeMessageUnSubscribe,
			expectedType:          UnSubscribeMessage,
			expectedTopic:         privetOrderChange,
			expectedPrivetChannel: true,
			expectedResponse:      true,
		},

		{
			name:                  "test-3",
			messageFunc:           NewOrderChangeV2MessageSubscribe,
			expectedType:          SubscribeMessage,
			expectedTopic:         privetOrderChangeV2,
			expectedPrivetChannel: true,
			expectedResponse:      true,
		},

		{
			name:                  "test-4",
			messageFunc:           NewOrderChangeV2MessageUnSubscribe,
			expectedType:          UnSubscribeMessage,
			expectedTopic:         privetOrderChangeV2,
			expectedPrivetChannel: true,
			expectedResponse:      true,
		},

		{
			name:                  "test-5",
			messageFunc:           NewAccountBalanceChangeMessageSubscribe,
			expectedType:          SubscribeMessage,
			expectedTopic:         accountBalanceChange,
			expectedPrivetChannel: true,
			expectedResponse:      true,
		},

		{
			name:                  "test-6",
			messageFunc:           NewAccountBalanceChangeMessageUnSubscribe,
			expectedType:          UnSubscribeMessage,
			expectedTopic:         accountBalanceChange,
			expectedPrivetChannel: true,
			expectedResponse:      true,
		},

		{
			name:                  "test-7",
			messageFunc:           NewStopOrderEventMessageSubscribe,
			expectedType:          SubscribeMessage,
			expectedTopic:         stopOrderEvent,
			expectedPrivetChannel: true,
			expectedResponse:      true,
		},

		{
			name:                  "test-8",
			messageFunc:           NewStopOrderEventMessageUnSubscribe,
			expectedType:          UnSubscribeMessage,
			expectedTopic:         stopOrderEvent,
			expectedPrivetChannel: true,
			expectedResponse:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			message := tc.messageFunc()

			assert.Equal(t, tc.expectedTopic, message.Topic)
			assert.Equal(t, tc.expectedType, message.Type)
			assert.Equal(t, tc.expectedPrivetChannel, message.PrivateChannel)
			assert.Equal(t, tc.expectedResponse, message.Response)
		})
	}
}

func TestTickersMessage(t *testing.T) {
	testCases := []struct {
		name                  string
		messageFunc           func(pairs ...string) subscribeMessage
		args                  []string
		expectedType          string
		expectedTopic         string
		expectedResponse      bool
		expectedPrivetChannel bool
	}{
		{
			name:                  "test-1",
			messageFunc:           NewTickerUnSubscribeMessages,
			args:                  []string{"BTC-USDT", "SOL-USDT", "ETH-USDT"},
			expectedTopic:         "/market/ticker:BTC-USDT,SOL-USDT,ETH-USDT",
			expectedType:          UnSubscribeMessage,
			expectedPrivetChannel: false,
			expectedResponse:      true,
		},

		{
			name:                  "test-2",
			messageFunc:           NewTickerSubscribeMessages,
			args:                  []string{"BTC-USDT", "SOL-USDT", "ETH-USDT"},
			expectedTopic:         "/market/ticker:BTC-USDT,SOL-USDT,ETH-USDT",
			expectedType:          SubscribeMessage,
			expectedPrivetChannel: false,
			expectedResponse:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			message := tc.messageFunc(tc.args...)

			assert.Equal(t, tc.expectedTopic, message.Topic)
			assert.Equal(t, tc.expectedPrivetChannel, message.PrivateChannel)
			assert.Equal(t, tc.expectedResponse, message.Response)
			assert.Equal(t, tc.expectedType, message.Type)

		})
	}
}

func TestCandlesMessages(t *testing.T) {
	testCases := []struct {
		name string

		messageFunc           func(pair string, candle kucoinentity.CandlePeriod) subscribeMessage
		firstArg              string
		secondArg             kucoinentity.CandlePeriod
		expectedTopic         string
		expectedType          string
		expectedResponse      bool
		expectedPrivetChannel bool
	}{
		{
			name:                  "test-1",
			messageFunc:           NewCandleSubscribeMessages,
			firstArg:              "BTC-USDT",
			secondArg:             kucoinentity.Min15,
			expectedType:          SubscribeMessage,
			expectedTopic:         "/market/candles:BTC-USDT_15min",
			expectedPrivetChannel: false,
			expectedResponse:      true,
		},

		{
			name:                  "test-2",
			messageFunc:           NewCandleUnSubscribeMessages,
			firstArg:              "ETH-USDT",
			secondArg:             kucoinentity.Day1,
			expectedType:          UnSubscribeMessage,
			expectedTopic:         "/market/candles:ETH-USDT_1day",
			expectedPrivetChannel: false,
			expectedResponse:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			message := tc.messageFunc(tc.firstArg, tc.secondArg)

			assert.Equal(t, tc.expectedType, message.Type)
			assert.Equal(t, tc.expectedTopic, message.Topic)
			assert.Equal(t, tc.expectedPrivetChannel, message.PrivateChannel)
			assert.Equal(t, tc.expectedResponse, message.Response)
		})
	}
}
