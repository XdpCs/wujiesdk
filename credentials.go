package wujiesdk

// @Title        credentials.go
// @Description  sign request
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2023-10-10 20:47

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

type Credentials struct {
	AppID         string
	PrivateKey    string
	cache         *cache.Cache
	RsaPrivateKey *rsa.PrivateKey
}

func (c *Credentials) Sign(req *http.Request) (*http.Request, error) {
	auth, found := c.cache.Get(HTTPHeaderAuthorization)
	if !found {
		sign, err := c.sign()
		if err != nil {
			return nil, fmt.Errorf("c.sign(): sign fail: %w", err)
		}
		c.cache.Set(HTTPHeaderAuthorization, sign, DefaultExpiration)
		auth = sign
	}
	req.Header.Set(HTTPHeaderAuthorization, auth.(string))
	return req, nil
}

func (c *Credentials) sign() (string, error) {
	var signContent struct {
		AppID     string `json:"appId"`
		Timestamp int64  `json:"timestamp"`
	}
	signContent.AppID = c.AppID
	signContent.Timestamp = time.Now().Unix()
	data, err := json.Marshal(&signContent)
	if err != nil {
		return "", fmt.Errorf("json.Marshal: sign content: %v, marshal sign content fail: %w", signContent, err)
	}

	hash := sha256.New()
	hash.Write(data)
	digest := hash.Sum(nil)
	signBytes, err := rsa.SignPKCS1v15(rand.Reader, c.RsaPrivateKey, crypto.SHA256, digest)
	if err != nil {
		return "", fmt.Errorf("rsa.SignPKCS1v15: digest: %v, sign fail: %w", digest, err)
	}
	signString := base64.StdEncoding.EncodeToString(signBytes)
	authorization := map[string]string{
		"sign":             signString,
		"secretKeyVersion": "1",
		"appId":            c.AppID,
		"original":         string(data),
	}
	auth, err := json.Marshal(authorization)
	if err != nil {
		return "", fmt.Errorf("json.Marshal: authorization: %v, marshal authorization fail: %w, ", authorization, err)
	}
	return string(auth), nil
}

func NewCredentials(appID, privateKey string) (*Credentials, error) {
	c := &Credentials{
		AppID:      appID,
		PrivateKey: privateKey,
		cache:      cache.New(DefaultExpiration, 10*time.Minute),
	}
	pkBytes, err := base64.StdEncoding.DecodeString(c.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("base64.StdEncoding.DecodeString: private key: %v, decode private key fail: %w", c.PrivateKey, err)
	}
	RsaPtr, err := x509.ParsePKCS8PrivateKey(pkBytes)
	if err != nil {
		return nil, fmt.Errorf("x509.ParsePKCS8PrivateKey: private bytes: %v, parse private key fail: %w", pkBytes, err)
	}
	RsaPk := RsaPtr.(*rsa.PrivateKey)
	c.RsaPrivateKey = RsaPk
	return c, nil
}

func (c *Credentials) BeforeRequest(req *http.Request) error {
	_, err := c.Sign(req)
	return err
}

func (c *Credentials) AfterRequest(_ *http.Response, _ error) {}
