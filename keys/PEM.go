package keys

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

//LoadPrivateKeyFromPEM loads a crypto.PrivateKey or crypto.PublicKey from PEM string as []byte
func LoadPrivateKeyFromPEM(pemKey []byte) (crypto.PrivateKey, error) {
	block, rest := pem.Decode(pemKey)
	var priv crypto.PrivateKey
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		priv, err = x509.ParseECPrivateKey(block.Bytes)
		if err != nil {
			priv, err = x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				return nil, err
			}
		}
	}

	fmt.Printf("Got a %T, with remaining data: %q\n", priv, rest)
	return priv, nil
}

//LoadPublicKeyFromPEM loads a crypto.PublicKey or crypto.PublicKey from PEM string as []byte
func LoadPublicKeyFromPEM(pemKey []byte) (crypto.PublicKey, error) {
	block, rest := pem.Decode(pemKey)
	var publ crypto.PublicKey
	publ, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		publ, err = x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}
	fmt.Printf("Got a %T, with remaining data: %q\n", publ, rest)

	return publ, nil
}
