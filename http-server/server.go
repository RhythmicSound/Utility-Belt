package httpserver

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"
)

//New returns a standard server with a single mux route of "/_healthcheck"
// which responds '200 OK'.
//
//If a mux is passed the above route is merged in and panics if already exists
func New(port string, mux *http.ServeMux) *http.Server {
	if mux == nil {
		mux = http.NewServeMux()
	}
	mux.HandleFunc("/_healthcheck", HealthcheckHandler)

	if port == "" {
		port = "8080"
	}
	return &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           mux,
		ReadTimeout:       2 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		IdleTimeout:       2 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion: tls.TLS_AES_128_GCM_SHA256,
		},
	}
}

//ValueFromReqQueryOrHeader finds a value from either the URL query string or the header
// using the given key.
//
//Useful when giving the client multiple methods to submit a value depending on
//their use case
//
//Checks URL query strings first and will return that value even if both are present.
//
//Both header and query string are checked as case insensitive.
//
//The first value found is returned.
func ValueFromReqQueryOrHeader(key string, req *http.Request) (string, error) {
	//If not found canonicalize casing and check again
	key = strings.ToLower(key)
	for k, v := range req.URL.Query() {
		if strings.ToLower(k) == key {
			return v[0], nil
		}
	}
	//otherwise check request header
	q := req.Header.Get(key)
	if q != "" {
		return q, nil
	}

	return q, fmt.Errorf("Key not found")
}
