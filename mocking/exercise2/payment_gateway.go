package exercise2

import (
	"errors"
	"github.com/Pallinder/go-randomdata"
	"math/rand"
	"strings"
	"time"
)

const (
	txn = "TXN"
)

// PaymentGateway is an interface for interacting with the payment gateway.
type PaymentGateway interface {
	ProcessPayment(amount float64, cardNumber string) (string, error)
}

// RealPaymentGateway is the implementation of PaymentGateway that uses a real payment gateway.
type RealPaymentGateway struct {
	// Some real payment gateway implementation here...
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1) == 1
}

// ProcessPayment processes a payment using the payment gateway.
func (pg *RealPaymentGateway) ProcessPayment(amount float64, cardNumber string) (string, error) {
	if !RandBool() {
		return "", errors.New("payment failed")
	}
	return strings.ToUpper(randomdata.Alphanumeric(12)), nil
}

// NewRealPaymentGateway creates a new instance of RealPaymentGateway.
func NewRealPaymentGateway() *RealPaymentGateway {
	return &RealPaymentGateway{}
}

func MakePayment(paymentGateway PaymentGateway, amount float64, cardNo string) (string, error) {
	txnId, err := paymentGateway.ProcessPayment(amount, cardNo)
	if err != nil {
		return "", err
	}

	return txnId, nil
}
