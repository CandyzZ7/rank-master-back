package net

import (
	"net/http/httptest"
	"testing"
)

func TestGetRemoteClientIp(t *testing.T) {
	tests := []struct {
		name       string
		headers    map[string]string
		remoteAddr string
		want       string
	}{
		{
			name: "X-Real-IP header set",
			headers: map[string]string{
				"X-Real-IP": "203.0.113.1",
			},
			remoteAddr: "198.51.100.1:1234",
			want:       "203.0.113.1",
		},
		{
			name: "X-Forwarded-For header set",
			headers: map[string]string{
				"X-Forwarded-For": "203.0.113.2",
			},
			remoteAddr: "198.51.100.2:5678",
			want:       "203.0.113.2",
		},
		{
			name:       "No headers set",
			headers:    map[string]string{},
			remoteAddr: "198.51.100.3:9012",
			want:       "198.51.100.3",
		},
		{
			name:       "Localhost",
			headers:    map[string]string{},
			remoteAddr: "::1:1234",
			want:       "127.0.0.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "http://example.com", nil)
			req.RemoteAddr = tt.remoteAddr
			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			got := GetRemoteClientIp(req)
			t.Logf("Test case: %s, Got IP: %s, Want IP: %s\n", tt.name, got, tt.want)
			if got != tt.want {
				t.Errorf("GetRemoteClientIp() = %v, want %v", got, tt.want)
			}
		})
	}
}
