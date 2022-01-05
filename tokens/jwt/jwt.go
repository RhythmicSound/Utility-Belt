package jwt

import (
	"encoding/json"
	"fmt"

	jose "gopkg.in/square/go-jose.v2"
)

//GetJWSPayloadUnsafely unmarshals the JWS token string without verification
// and returns the payload
//
//Verification of the JWS should be done at gateway or elsewhere in the application
func GetJWSPayloadUnsafely(authToken string) (map[string]interface{}, error) {
	tokenElements, err := jose.ParseSigned(authToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal authToken: %v", err)

	}

	payloadJSON := tokenElements.UnsafePayloadWithoutVerification()
	payload := make(map[string]interface{})
	err = json.Unmarshal(payloadJSON, &payload)
	return payload, err
}
