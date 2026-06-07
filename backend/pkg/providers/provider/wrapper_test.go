package provider

import (
	"errors"
	"testing"
	"time"
)

func TestIsTooManyRequestsError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "openrouter provider returned 429",
			err:  errors.New("API returned unexpected status code: 429: Provider returned error"),
			want: true,
		},
		{
			name: "statuscode compact format",
			err:  errors.New("statusCode: 429"),
			want: true,
		},
		{
			name: "rate limit text",
			err:  errors.New("provider rate limit exceeded"),
			want: true,
		},
		{
			name: "generic provider error is not retryable",
			err:  errors.New("provider returned error: invalid model"),
			want: false,
		},
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTooManyRequestsError(tt.err); got != tt.want {
				t.Fatalf("isTooManyRequestsError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTooManyRequestsRetryDelay(t *testing.T) {
	if got := tooManyRequestsRetryDelay(0); got != 10*time.Second {
		t.Fatalf("tooManyRequestsRetryDelay(0) = %s, want 10s", got)
	}

	if got := tooManyRequestsRetryDelay(100); got != MaxTooManyRequestsRetryDelay {
		t.Fatalf("tooManyRequestsRetryDelay(100) = %s, want %s", got, MaxTooManyRequestsRetryDelay)
	}
}
