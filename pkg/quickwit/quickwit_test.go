package quickwit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeResourceErrorBody(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		expected   string
	}{
		{
			name:       "wraps plain text upstream error",
			statusCode: http.StatusGatewayTimeout,
			body:       "upstream request timeout\n",
			expected:   `{"message":"upstream request timeout","status":504}`,
		},
		{
			name:       "preserves JSON error",
			statusCode: http.StatusBadGateway,
			body:       `{"message":"upstream unavailable","status":502}`,
			expected:   `{"message":"upstream unavailable","status":502}`,
		},
		{
			name:       "preserves successful response",
			statusCode: http.StatusOK,
			body:       `{"fields":{}}`,
			expected:   `{"fields":{}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, string(normalizeResourceErrorBody(test.statusCode, []byte(test.body))))
		})
	}
}