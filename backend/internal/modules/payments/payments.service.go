package payments

type PaymentService struct {
	PaymentProvider *PaymentProvider
}

func (r *PaymentService) GetPayments(userID int) ([]Payment, error) {
	return r.PaymentProvider.GetPayments(userID)
}
func (r *PaymentService) UpdatePayment(payment PaymentUpdateDto, paymentID int, userID int) error {
	return r.PaymentProvider.UpdatePayment(payment, paymentID, userID)

}
func (r *PaymentService) CreatePayment(payment PaymentCreateDto, userID int) (int, error) {
	return r.PaymentProvider.CreatePayment(payment, userID)

}
func (r *PaymentService) DeletePayments(paymentIDs []int, userID int) error {
	return r.PaymentProvider.DeletePayments(paymentIDs, userID)

}

func (r *PaymentService) GetPaymentTemplates(userID int) ([]PaymentTemplate, error) {
	return r.PaymentProvider.GetPaymentTemplate(userID)
}
func (r *PaymentService) UpdatePaymentTemplate(template PaymentTemplateUpdateDto, templateID int, userID int) error {
	return r.PaymentProvider.UpdatePaymentTemplate(template, templateID, userID)

}
func (r *PaymentService) CreatePaymentTemplate(template PaymentTemplateCreateDto, userID int) (int, error) {
	return r.PaymentProvider.CreatePaymentTemplate(template, userID)

}
func (r *PaymentService) DeletePaymentTemplates(templateIDs []int, userID int) error {
	return r.PaymentProvider.DeletePaymentTemplates(templateIDs, userID)

}
