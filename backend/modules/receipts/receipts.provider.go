package receipts

import (
	"database/sql"

	"github.com/lib/pq"
)

type ReceiptProvider struct {
	Db *sql.DB
}

func (r *ReceiptProvider) GetReceipts(userID int) ([]Receipt, error) {
	var receipts []Receipt
	var err error
	var rows *sql.Rows

	if rows, err = r.Db.Query(`
		SELECT rec.id, tem.name, rec.amount, rec.date 
		FROM balance.receipts rec
		LEFT JOIN balance.receipt_templates tem
		ON rec.receipt_id = tem.id
		WHERE rec.user_id = $1

	`, userID); err != nil {
		return receipts, err
	}
	defer rows.Close()
	receipts = make([]Receipt, 0, 10)
	for rows.Next() {
		receipt := Receipt{}
		err := rows.Scan(&receipt.ID, &receipt.Name, &receipt.Amount, &receipt.Date)
		if err != nil {
			return receipts, err
		}
		receipts = append(receipts, receipt)
	}

	return receipts, err
}
func (r *ReceiptProvider) UpdateReceipt(receipt ReceiptUpdate, receiptID int, userID int) error {
	var err error

	_, err = r.Db.Exec(`
	UPDATE balance.receipts
	SET
		receipt_id = $1,
		amount = $2
	WHERE
		id = $3 AND
		user_id = $4
	`, receipt.ReceiptID, receipt.Amount, receiptID, userID)

	return err
}
func (r *ReceiptProvider) CreateReceipt(receipt ReceiptCreate, userID int) (int, error) {
	var receiptID int

	err := r.Db.QueryRow(`
	INSERT INTO balance.receipts 
	(amount, receipt_id, user_id) 
	VALUES ($1, $2, $3) 
	RETURNING id
	`, receipt.Amount, receipt.ReceiptID, userID).Scan(&receiptID)

	return receiptID, err
}
func (r *ReceiptProvider) DeleteReceipts(receiptIDs []int, userID int) error {
	var err error

	_, err = r.Db.Exec(`
	DELETE FROM balance.receipts
	WHERE user_id = $1 AND id = ANY($2) 
	`, userID, pq.Array(receiptIDs))

	return err
}

func (r *ReceiptProvider) GetReceiptTemplate(userID int) ([]ReceiptTemplate, error) {
	var templates []ReceiptTemplate
	var err error
	var rows *sql.Rows

	if rows, err = r.Db.Query(`
		SELECT id, name, amount
		FROM balance.receipt_templates
		WHERE user_id = $1

	`, userID); err != nil {
		return templates, err
	}
	defer rows.Close()
	templates = make([]ReceiptTemplate, 0, 10)
	for rows.Next() {
		template := ReceiptTemplate{}
		err := rows.Scan(&template.ID, &template.Name, &template.Amount)
		if err != nil {
			return templates, err
		}
		templates = append(templates, template)
	}

	return templates, err
}
func (r *ReceiptProvider) UpdateReceiptTemplate(template ReceiptTemplateUpdate, templateID int, userID int) error {
	var err error

	_, err = r.Db.Exec(`
	UPDATE balance.receipt_templates
	SET
		name = $1,
		amount = $2
	WHERE
		id = $3 AND
		user_id = $4
	`, template.Name, template.Amount, templateID, userID)

	return err
}
func (r *ReceiptProvider) CreateReceiptTemplate(template ReceiptTemplateCreate, userID int) (int, error) {
	var templateID int

	err := r.Db.QueryRow(`
	INSERT INTO balance.receipt_templates
	(name, amount, user_id) 
	VALUES ($1, $2, $3) 
	RETURNING id
	`, template.Name, template.Amount, userID).Scan(&templateID)

	return templateID, err
}
func (r *ReceiptProvider) DeleteReceiptTemplates(receiptIDs []int, userID int) error {
	var err error

	_, err = r.Db.Exec(`
	DELETE FROM balance.receipt_templates
	WHERE user_id = $1 AND id = ANY($2) 
	`, userID, pq.Array(receiptIDs))

	return err
}
