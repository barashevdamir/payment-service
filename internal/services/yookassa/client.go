package yookassa

import (
	"net/url"
)

type Formatter func() (string, *url.URL)

type endpoint struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

type yookassaClient struct {
	Scheme    string `yaml:"scheme"`
	Host      string `yaml:"host"`
	SecretKey string
	AccountId string
	Endpoints yookassaEndpoints `yaml:"endpoints"`
}

type yookassaEndpoints struct {
	Payments endpoint `yaml:"payment"`
}

type Amount struct {
	Value    float64
	Currency string "RUB"
}

type Confirmation struct {
	Type      string "redirect"
	Locale    string
	ReturnUrl string
	Enforce   bool
}

func NewYookassaClient(scheme, host, secretKey, accountId string, endpoints yookassaEndpoints) *yookassaClient {
	return &yookassaClient{
		Scheme:    scheme,
		Host:      host,
		SecretKey: secretKey,
		AccountId: accountId,
		Endpoints: endpoints,
	}
}
