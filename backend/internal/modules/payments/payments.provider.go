package payments

import (
	"database/sql"

	"github.com/lib/pq"
)

type PaymentProvider struct {
	Db *sql.DB
}

func (r *PaymentProvider) GetPayments(userID int) ([]Payment, error) {
	var payments []Payment
	var err error
	var rows *sql.Rows

	if rows, err = r.Db.Query(`
		SELECT pay.id, tem.name, pay.amount, pay.date, pay.payment_id, pay.description
		FROM balance.payments pay
		LEFT JOIN balance.payment_templates tem
		ON pay.payment_id = tem.id
		WHERE pay.user_id = $1
		ORDER BY pay.date DESC

	`, userID); err != nil {
		return payments, err
	}
	defer rows.Close()
	payments = make([]Payment, 0, 10)
	for rows.Next() {
		payment := Payment{}
		err := rows.Scan(&payment.ID, &payment.Name, &payment.Amount, &payment.Date, &payment.PaymentID, &payment.Description)
		if err != nil {
			return payments, err
		}
		payments = append(payments, payment)
	}

	return payments, err
}
func (r *PaymentProvider) UpdatePayment(payment PaymentUpdateDto, paymentID int, userID int) error {
	var err error

	_, err = r.Db.Exec(`
	UPDATE balance.payments
	SET
		payment_id = $1,
		amount = $2,
		description = $3
	WHERE
		id = $4 AND
		user_id = $5
	`, payment.PaymentID, payment.Amount, payment.Description, paymentID, userID)

	return err
}
func (r *PaymentProvider) CreatePayment(payment PaymentCreateDto, userID int) (int, error) {
	var paymentID int

	err := r.Db.QueryRow(`
	INSERT INTO balance.payments 
	(amount, payment_id, user_id, description) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id
	`, payment.Amount, payment.PaymentID, userID, payment.Description).Scan(&paymentID)

	return paymentID, err
}
func (r *PaymentProvider) DeletePayments(paymentIDs []int, userID int) error {
	var err error

	_, err = r.Db.Exec(`
	DELETE FROM balance.payments
	WHERE user_id = $1 AND id = ANY($2) 
	`, userID, pq.Array(paymentIDs))

	return err
}

func (r *PaymentProvider) GetPaymentTemplate(userID int) ([]PaymentTemplate, error) {
	var templates []PaymentTemplate
	var err error
	var rows *sql.Rows

	if rows, err = r.Db.Query(`
		SELECT id, name, amount
		FROM balance.payment_templates
		WHERE user_id = $1

	`, userID); err != nil {
		return templates, err
	}
	defer rows.Close()
	templates = make([]PaymentTemplate, 0, 10)
	for rows.Next() {
		template := PaymentTemplate{}
		err := rows.Scan(&template.ID, &template.Name, &template.Amount)
		if err != nil {
			return templates, err
		}
		templates = append(templates, template)
	}

	return templates, err
}
func (r *PaymentProvider) UpdatePaymentTemplate(template PaymentTemplateUpdateDto, templateID int, userID int) error {
	var err error

	_, err = r.Db.Exec(`
	UPDATE balance.payment_templates
	SET
		name = $1,
		amount = $2
	WHERE
		id = $3 AND
		user_id = $4
	`, template.Name, template.Amount, templateID, userID)

	return err
}
func (r *PaymentProvider) CreatePaymentTemplate(template PaymentTemplateCreateDto, userID int) (int, error) {
	var templateID int

	err := r.Db.QueryRow(`
	INSERT INTO balance.payment_templates
	(name, amount, user_id) 
	VALUES ($1, $2, $3) 
	RETURNING id
	`, template.Name, template.Amount, userID).Scan(&templateID)

	return templateID, err
}
func (r *PaymentProvider) DeletePaymentTemplates(paymentIDs []int, userID int) error {
	var err error

	_, err = r.Db.Exec(`
	DELETE FROM balance.payment_templates
	WHERE user_id = $1 AND id = ANY($2) 
	`, userID, pq.Array(paymentIDs))

	return err
}
