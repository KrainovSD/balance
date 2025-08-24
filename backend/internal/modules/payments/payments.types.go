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

type PaymentCreateDto struct {
	Amount    *int `json:"amount"`
	PaymentID *int `json:"paymentId"`
}

func (r *PaymentCreateDto) Validate() error {

	if r.PaymentID == nil {
		return errors.New("paymentId is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}

type PaymentUpdateDto struct {
	Amount    *int `json:"amount"`
	PaymentID *int `json:"paymentId"`
}

func (r *PaymentUpdateDto) Validate() error {

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

type PaymentTemplateCreateDto struct {
	Name   *string `json:"name"`
	Amount *int    `json:"amount"`
}

func (r *PaymentTemplateCreateDto) Validate() error {

	if r.Name == nil {
		return errors.New("name is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}

type PaymentTemplateUpdateDto struct {
	Name   *string `json:"name"`
	Amount *int    `json:"amount"`
}

func (r *PaymentTemplateUpdateDto) Validate() error {

	if r.Name == nil {
		return errors.New("name is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}
