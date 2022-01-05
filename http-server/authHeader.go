package httpserver

import (
	"fmt"
	"strings"
)

//GetTokenFromBearerString take a 'Bearer ...' string from a http.request header and returns the token without the Bearer prefix.
//
//Returns error if auth string is empty
func GetTokenFromBearerString(authorizationString string) (string, error) {
	bearerSchema := "Bearer "
	if !strings.HasPrefix(authorizationString, bearerSchema) || authorizationString == "" || len(bearerSchema) >= len(authorizationString) {
		return "", fmt.Errorf("Authorization string is not a Bearer token")
	}
	authToken := strings.Split(authorizationString, bearerSchema)[1]
	authToken = strings.TrimSpace(authToken)
	return authToken, nil
}
