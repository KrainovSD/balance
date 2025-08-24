package receipts

import (
	"balance/internal/lib/web"
	oauthPlugin "balance/internal/plugins/oauth"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type ReceiptControllerOptions struct {
	Db   *sql.DB
	Auth *oauthPlugin.Auth
}
type ReceiptInterface interface {
	Init(mux *http.ServeMux)
}

func CreateReceiptController(options ReceiptControllerOptions) ReceiptInterface {
	return &ReceiptController{
		Auth: options.Auth,
		ReceiptService: &ReceiptService{
			ReceiptProvider: &ReceiptProvider{
				Db: options.Db,
			},
		},
	}
}

type ReceiptController struct {
	ReceiptService *ReceiptService
	Auth           *oauthPlugin.Auth
}

func (r *ReceiptController) GetReceipts(w http.ResponseWriter, req *http.Request) {
	var receipts []Receipt
	var userID int
	var err error
	userID, _ = oauthPlugin.GetUserId(req)

	if receipts, err = r.ReceiptService.GetReceipts(userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receipts)

}
func (r *ReceiptController) UpdateReceipt(w http.ResponseWriter, req *http.Request) {
	var userID int
	var receipt ReceiptUpdateDto
	var receiptID int
	var err error
	var body []byte
	userID, _ = oauthPlugin.GetUserId(req)

	if receiptID, err = strconv.Atoi(req.PathValue("ID")); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	if body, err = io.ReadAll(req.Body); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	defer req.Body.Close()

	if err = json.Unmarshal(body, &receipt); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if err = receipt.Validate(); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if err = r.ReceiptService.UpdateReceipt(receipt, receiptID, userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}
func (r *ReceiptController) CreateReceipt(w http.ResponseWriter, req *http.Request) {
	var userID int
	var receiptID int
	var receipt ReceiptCreateDto
	var err error
	var body []byte
	userID, _ = oauthPlugin.GetUserId(req)

	if body, err = io.ReadAll(req.Body); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	defer req.Body.Close()

	if err = json.Unmarshal(body, &receipt); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	if err = receipt.Validate(); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if receiptID, err = r.ReceiptService.CreateReceipt(receipt, userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receiptID)
}
func (r *ReceiptController) DeleteReceipts(w http.ResponseWriter, req *http.Request) {
	var userID int
	var err error
	var receiptIDs []int
	var body []byte
	userID, _ = oauthPlugin.GetUserId(req)

	if body, err = io.ReadAll(req.Body); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	defer req.Body.Close()

	if err = json.Unmarshal(body, &receiptIDs); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if err = r.ReceiptService.DeleteReceipts(receiptIDs, userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

func (r *ReceiptController) GetReceiptTemplates(w http.ResponseWriter, req *http.Request) {
	var templates []ReceiptTemplate
	var userID int
	var err error
	userID, _ = oauthPlugin.GetUserId(req)

	if templates, err = r.ReceiptService.GetReceiptTemplates(userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templates)

}
func (r *ReceiptController) UpdateReceiptTemplate(w http.ResponseWriter, req *http.Request) {
	var userID int
	var template ReceiptTemplateUpdateDto
	var templateID int
	var body []byte
	var err error
	userID, _ = oauthPlugin.GetUserId(req)

	if templateID, err = strconv.Atoi(req.PathValue("ID")); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	if body, err = io.ReadAll(req.Body); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	defer req.Body.Close()

	if err = json.Unmarshal(body, &template); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if err = template.Validate(); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if err = r.ReceiptService.UpdateReceiptTemplate(template, templateID, userID); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}
func (r *ReceiptController) CreateReceiptTemplate(w http.ResponseWriter, req *http.Request) {
	var userID int
	var templateID int
	var template ReceiptTemplateCreateDto
	var err error
	var body []byte
	userID, _ = oauthPlugin.GetUserId(req)

	if body, err = io.ReadAll(req.Body); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	defer req.Body.Close()

	if err = json.Unmarshal(body, &template); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	if err = template.Validate(); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if templateID, err = r.ReceiptService.CreateReceiptTemplate(template, userID); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templateID)
}
func (r *ReceiptController) DeleteReceiptTemplates(w http.ResponseWriter, req *http.Request) {
	var userID int
	var err error
	var templateIDs []int
	var body []byte
	userID, _ = oauthPlugin.GetUserId(req)

	if body, err = io.ReadAll(req.Body); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	defer req.Body.Close()

	if err = json.Unmarshal(body, &templateIDs); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if err = r.ReceiptService.DeleteReceiptTemplates(templateIDs, userID); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

func (r *ReceiptController) Init(mux *http.ServeMux) {

	mux.Handle("/api/v1/receipts", r.Auth.Middleware(true)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			r.CreateReceipt(w, req)
		case "GET":
			r.GetReceipts(w, req)
		case "DELETE":
			r.DeleteReceipts(w, req)
		default:
			w.WriteHeader(405)
		}
	})))
	mux.Handle("/api/v1/receipts/{ID}", r.Auth.Middleware(true)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "PATCH":
			r.UpdateReceipt(w, req)
		default:
			w.WriteHeader(405)
		}
	})))

	mux.Handle("/api/v1/receipt_templates", r.Auth.Middleware(true)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			r.CreateReceiptTemplate(w, req)
		case "GET":
			r.GetReceiptTemplates(w, req)
		case "DELETE":
			r.DeleteReceiptTemplates(w, req)
		default:
			w.WriteHeader(405)
		}
	})))
	mux.Handle("/api/v1/receipt_templates/{ID}", r.Auth.Middleware(true)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "PATCH":
			r.UpdateReceiptTemplate(w, req)
		default:
			w.WriteHeader(405)
		}
	})))
}
