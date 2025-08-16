package receipts

type ReceiptService struct {
	ReceiptProvider ReceiptProvider
}

func (r *ReceiptService) GetReceipts(userID int) ([]Receipt, error) {
	return r.ReceiptProvider.GetReceipts(userID)
}
func (r *ReceiptService) UpdateReceipt(receipt ReceiptUpdate, receiptID int, userID int) error {
	return r.ReceiptProvider.UpdateReceipt(receipt, receiptID, userID)

}
func (r *ReceiptService) CreateReceipt(receipt ReceiptCreate, userID int) (int, error) {
	return r.ReceiptProvider.CreateReceipt(receipt, userID)

}
func (r *ReceiptService) DeleteReceipts(receiptIDs []int, userID int) error {
	return r.ReceiptProvider.DeleteReceipts(receiptIDs, userID)

}

func (r *ReceiptService) GetReceiptTemplates(userID int) ([]ReceiptTemplate, error) {
	return r.ReceiptProvider.GetReceiptTemplate(userID)
}
func (r *ReceiptService) UpdateReceiptTemplate(template ReceiptTemplateUpdate, templateID int, userID int) error {
	return r.ReceiptProvider.UpdateReceiptTemplate(template, templateID, userID)

}
func (r *ReceiptService) CreateReceiptTemplate(template ReceiptTemplateCreate, userID int) (int, error) {
	return r.ReceiptProvider.CreateReceiptTemplate(template, userID)

}
func (r *ReceiptService) DeleteReceiptTemplates(templateIDs []int, userID int) error {
	return r.ReceiptProvider.DeleteReceiptTemplates(templateIDs, userID)

}
