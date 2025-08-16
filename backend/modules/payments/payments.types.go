package payments

import (
	"errors"
)

type Payment struct {
	ID     int    `json:"id"`
	Amount int    `json:"amount"`
	Name   string `json:"name"`
	Date   string `json:"date"`
}

type PaymentCreate struct {
	Amount    *int `json:"amount"`
	PaymentID *int `json:"paymentId"`
}

func (r *PaymentCreate) Validate() error {

	if r.PaymentID == nil {
		return errors.New("paymentId is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}

type PaymentUpdate struct {
	Amount    *int `json:"amount"`
	PaymentID *int `json:"paymentId"`
}

func (r *PaymentUpdate) Validate() error {

	if r.PaymentID == nil {
		return errors.New("paymentId is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}

type PaymentTemplate struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type PaymentTemplateCreate struct {
	Name   *string `json:"name"`
	Amount *int    `json:"amount"`
}

func (r *PaymentTemplateCreate) Validate() error {

	if r.Name == nil {
		return errors.New("name is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}

type PaymentTemplateUpdate struct {
	Name   *string `json:"name"`
	Amount *int    `json:"amount"`
}

func (r *PaymentTemplateUpdate) Validate() error {

	if r.Name == nil {
		return errors.New("name is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}
