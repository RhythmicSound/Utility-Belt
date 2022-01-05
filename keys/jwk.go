package keys

import (
	"crypto"
	"encoding/json"
	"fmt"

	"github.com/lestrrat-go/jwx/jwk"
)

//NewJWK returns a new JWK set comprising of the given Public Key and furnished with
// the given config settings as per the RFC7517 standard
func NewJWK(publicKey crypto.PublicKey, configs map[string]interface{}) ([]byte, error) {
	//get jwk from public key
	key, err := jwk.New(publicKey)
	for k, v := range configs {
		key.Set(k, v)
	}

	keychain := jwk.NewSet()
	ok := keychain.Add(key)
	if !ok {
		return nil, fmt.Errorf("Could not add key to keychain")
	}

	//json representation
	jwkJSON, err := json.Marshal(keychain)
	if err != nil {
		return nil, fmt.Errorf("Could not marshal public key jwk: %+v", err)
	}
	return jwkJSON, nil
}
