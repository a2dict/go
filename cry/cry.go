package cry

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

// ParseBase64PKCS8PrivateKey ...
func ParseBase64PKCS8PrivateKey(s string) (priv *rsa.PrivateKey, err error) {
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return
	}
	k, err := x509.ParsePKCS8PrivateKey(bs)
	if err != nil {
		return
	}
	priv = k.(*rsa.PrivateKey)
	return
}

// ParseBase64PKIXPublicKey ...
func ParseBase64PKIXPublicKey(s string) (pub *rsa.PublicKey, err error) {
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return
	}
	k, err := x509.ParsePKIXPublicKey(bs)
	if err != nil {
		return
	}
	pub = k.(*rsa.PublicKey)
	return
}

// RsaSignWithHash ...
func RsaSignWithHash(content string, hash crypto.Hash, priv *rsa.PrivateKey) (string, error) {
	h := hash.New()
	h.Write([]byte(content))
	digest := h.Sum(nil)
	sig, err := rsa.SignPKCS1v15(rand.Reader, priv, hash, digest)
	if err != nil {
		return "", err
	}
	s := base64.StdEncoding.EncodeToString(sig)
	return s, nil
}

// RsaSignVerifyWithHash ...
func RsaSignVerifyWithHash(content, sign string, hash crypto.Hash, pub *rsa.PublicKey) error {
	h := hash.New()
	h.Write([]byte(content))
	digest := h.Sum(nil)
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pub, hash, digest, signature)
}
