package receipts

import (
	"errors"
)

type Receipt struct {
	ID          int    `json:"id"`
	Amount      int    `json:"amount"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ReceiptID   int    `json:"receiptId"`
	Date        string `json:"date"`
}

type ReceiptCreateDto struct {
	Amount      *int    `json:"amount"`
	ReceiptID   *int    `json:"receiptId"`
	Description *string `json:"description"`
}

func (r *ReceiptCreateDto) Validate() error {

	if r.ReceiptID == nil {
		return errors.New("receiptId is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}
	if r.Description == nil {
		return errors.New("description is required")
	}
	return nil
}

type ReceiptUpdateDto struct {
	Amount      *int    `json:"amount"`
	ReceiptID   *int    `json:"receiptId"`
	Description *string `json:"description"`
}

func (r *ReceiptUpdateDto) Validate() error {

	if r.ReceiptID == nil && r.Amount == nil && r.Description == nil {
		return errors.New("something for update is required")
	}

	return nil
}

type ReceiptTemplate struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ReceiptTemplateCreateDto struct {
	Name   *string `json:"name"`
	Amount *int    `json:"amount"`
}

func (r *ReceiptTemplateCreateDto) Validate() error {

	if r.Name == nil {
		return errors.New("name is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}

type ReceiptTemplateUpdateDto struct {
	Name   *string `json:"name"`
	Amount *int    `json:"amount"`
}

func (r *ReceiptTemplateUpdateDto) Validate() error {

	if r.Name == nil {
		return errors.New("name is required")
	}
	if r.Amount == nil {
		return errors.New("amount is required")
	}

	return nil
}
