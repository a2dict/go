package cry

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"github.com/a2dict/go/str"
)

var (
	ErrMalformedPem = errors.New("malformed pem")
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

// ParseBase64PKCS1PrivateKey ...
func ParseBase64PKCS1PrivateKey(s string) (priv *rsa.PrivateKey, err error) {
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return
	}
	priv, err = x509.ParsePKCS1PrivateKey(bs)
	if err != nil {
		return
	}
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

// VerifyRsaSignWithHash ...
func VerifyRsaSignWithHash(content, sign string, hash crypto.Hash, pub *rsa.PublicKey) error {
	h := hash.New()
	h.Write([]byte(content))
	digest := h.Sum(nil)
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pub, hash, digest, signature)
}

// ParsePemPKCS8PrivateKey ...
func ParsePemPKCS8PrivateKey(s string) (priv *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(s))
	k, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	priv = k.(*rsa.PrivateKey)
	return
}

// ParsePemPKCS1PrivateKey ...
func ParsePemPKCS1PrivateKey(s string) (priv *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(s))
	if block == nil {
		return nil, ErrMalformedPem
	}
	priv, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return
}

// ParsePemPKIXPublicKey ...
func ParsePemPKIXPublicKey(s string) (pub *rsa.PublicKey, err error) {
	block, _ := pem.Decode([]byte(s))
	if block == nil {
		return nil, ErrMalformedPem
	}
	k, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	pub = k.(*rsa.PublicKey)
	return
}

// RsaKeyPair ...
type RsaKeyPair struct {
	Priv         *rsa.PrivateKey
	Pub          *rsa.PublicKey
	PrivPKCS1Pem string
	PubPKIXPem   string
}

// GenRsaKey ...
func GenRsaKey(bits int) (*RsaKeyPair, error) {
	priv, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	var privBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}
	privPkcs1PemBytes := pem.EncodeToMemory(privBlock)

	pub := &priv.PublicKey
	pubBlockBytes, err := x509.MarshalPKIXPublicKey(pub)

	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBlockBytes,
	}
	pubPkixPemBytes := pem.EncodeToMemory(pubBlock)
	return &RsaKeyPair{
		Priv:         priv,
		Pub:          pub,
		PrivPKCS1Pem: string(privPkcs1PemBytes),
		PubPKIXPem:   string(pubPkixPemBytes),
	}, nil
}

// VerifyKeyPair ...
func VerifyKeyPair(priv *rsa.PrivateKey, pub *rsa.PublicKey) error {
	content := str.RandStrWithCharset(32, str.Letters)
	s, err := RsaSignWithHash(content, crypto.SHA256, priv)
	if err != nil {
		return err
	}
	err = VerifyRsaSignWithHash(content, s, crypto.SHA256, pub)
	if err != nil {
		return err
	}
	return nil
}
