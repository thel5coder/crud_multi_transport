package jwe

import (
	"github.com/lestrrat/go-jwx/jwa"
	"github.com/lestrrat/go-jwx/jwe"
)

// Credential ...
type Credential struct {
	KeyLocation string
	Passphrase  string
}

// Generate ...
func (cred *Credential) GenerateJwePayload(value string) (res string, err error) {
	privkey, err := rsaConfigSetup(cred.KeyLocation, cred.Passphrase)
	if err != nil {
		return res, err
	}

	// Generate JWE
	jweRes, err := jwe.Encrypt([]byte(value), jwa.RSA1_5, &privkey.PublicKey, jwa.A128CBC_HS256, jwa.Deflate)
	res = string(jweRes)

	return res, err
}

// Rollback ...
func (cred *Credential) Rollback(value string) (res string, err error) {
	privkey, err := rsaConfigSetup(cred.KeyLocation, cred.Passphrase)
	if err != nil {
		println(err.Error())
		return res, err
	}

	decrypted, err := jwe.Decrypt([]byte(value), jwa.RSA1_5, privkey)
	if err != nil {
		return res, err
	}
	res = string(decrypted)

	return res, err
}
