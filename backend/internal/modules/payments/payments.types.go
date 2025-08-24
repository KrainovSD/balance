package payments

import (
	"errors"
)

type Payment struct {
	ID          int    `json:"id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Name        string `json:"name"`
	PaymentID   int    `json:"paymentId"`
	Date        string `json:"date"`
}

type PaymentCreateDto struct {
	Amount      *int    `json:"amount"`
	PaymentID   *int    `json:"paymentId"`
	Description *string `json:"description"`
}

func (r *PaymentCreateDto) Validate() error {

	if r.PaymentID == nil {
		return errors.New("paymentId is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}
	if r.Description == nil {
		return errors.New("description is required")
	}

	return nil
}

type PaymentUpdateDto struct {
	Amount      *int    `json:"amount"`
	PaymentID   *int    `json:"paymentId"`
	Description *string `json:"description"`
}

func (r *PaymentUpdateDto) Validate() error {

	if r.PaymentID == nil && r.Amount == nil && r.Description == nil {
		return errors.New("something for update is required")
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
