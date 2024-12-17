package yookassa

import (
	"net/url"
)

type Formatter func() (string, *url.URL)

type endpoint struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

type YookassaClient struct {
	Scheme    string `yaml:"scheme"`
	Host      string `yaml:"host"`
	SecretKey string
	AccountId string
	Endpoints yookassaEndpoints `yaml:"endpoints"`
}

type yookassaEndpoints struct {
	CreatePayment endpoint `yaml:"create_payment"`
	PaymentsList  endpoint `yaml:"payments_list"`
}

type Amount struct {
	Value    float64
	Currency string
}

type Confirmation struct {
	Type      string
	Locale    string
	ReturnUrl string
	Enforce   bool
}

type Recipient struct {
	AccountId string
	GatewayId string
}

type PaymentMethod struct {
	Type           string
	Id             string
	Saved          bool
	Title          string
	DiscountAmount Amount
	LoanOption     string
}

type CancellationDetails struct {
	Party  string
	Reason string
}

type ThreeDSecure struct {
	Applied bool
}
type AuthorizationDetails struct {
	RRN          string
	AuthCode     string
	ThreeDSecure ThreeDSecure
}

type Transfer struct {
	AccountId         string
	Amount            Amount
	Status            string
	PlatformFeeAmount Amount
	Description       string
	Metadata          map[string]string
}

type Settlement struct {
	Type   string
	Amount Amount
}

type Deal struct {
	Id          string
	Settlements []Settlement
}

type InvoiceDetails struct {
	Id string
}

type Payment struct {
	Id                    string
	Status                string
	Amount                Amount
	IncomeAmount          Amount
	Description           string
	Recipient             Recipient
	PaymentMethod         PaymentMethod
	CapturedAt            string
	CreatedAt             string
	ExpiresAt             string
	Confirmation          Confirmation
	Test                  bool
	RefundedAmount        Amount
	Paid                  bool
	Refundable            bool
	RecipientRegistration string
	Metadata              map[string]string
	CancellationDetails   CancellationDetails
	AuthorizationDetails  AuthorizationDetails
	Transfers             []Transfer
	Deal                  Deal
	MerchantCustomerId    string
	InvoiceDetails        InvoiceDetails
}

func NewYookassaClient(scheme string, host string, secretKey string, accountId string) *YookassaClient {
	endpoints := yookassaEndpoints{
		CreatePayment: endpoint{
			Path:   "/payments",
			Method: "POST",
		},
		PaymentsList: endpoint{
			Path:   "/payments",
			Method: "GET",
		},
	}
	return &YookassaClient{
		Scheme:    scheme,
		Host:      host,
		SecretKey: secretKey,
		AccountId: accountId,
		Endpoints: endpoints,
	}
}
