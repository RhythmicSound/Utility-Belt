package httpclient

import (
	"net"
	"net/http"
	"time"
)

func New() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,

		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,

			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,

			ExpectContinueTimeout: 1 * time.Second,

			IdleConnTimeout:     30 * time.Second,
			MaxIdleConnsPerHost: 10,
			MaxIdleConns:        100,

			DisableKeepAlives: false,
			ForceAttemptHTTP2: true,
		},
	}
}
