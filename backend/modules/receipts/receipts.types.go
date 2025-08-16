package receipts

import (
	"errors"
)

type Receipt struct {
	ID     int    `json:"id"`
	Amount int    `json:"amount"`
	Name   string `json:"name"`
	Date   string `json:"date"`
}

type ReceiptCreate struct {
	Amount    *int `json:"amount"`
	ReceiptID *int `json:"receiptId"`
}

func (r *ReceiptCreate) Validate() error {

	if r.ReceiptID == nil {
		return errors.New("receiptId is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}

type ReceiptUpdate struct {
	Amount    *int `json:"amount"`
	ReceiptID *int `json:"receiptId"`
}

func (r *ReceiptUpdate) Validate() error {

	if r.ReceiptID == nil {
		return errors.New("receiptId is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}

type ReceiptTemplate struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ReceiptTemplateCreate struct {
	Name   *string `json:"name"`
	Amount *int    `json:"amount"`
}

func (r *ReceiptTemplateCreate) Validate() error {

	if r.Name == nil {
		return errors.New("name is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}

type ReceiptTemplateUpdate struct {
	Name   *string `json:"name"`
	Amount *int    `json:"amount"`
}

func (r *ReceiptTemplateUpdate) Validate() error {

	if r.Name == nil {
		return errors.New("name is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}
