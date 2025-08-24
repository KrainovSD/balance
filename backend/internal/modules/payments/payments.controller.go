package payments

import (
	"balance/internal/lib/web"
	oauthPlugin "balance/internal/plugins/oauth"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type PaymentControllerOptions struct {
	Db   *sql.DB
	Auth *oauthPlugin.Auth
}
type PaymentInterface interface {
	Init(mux *http.ServeMux)
}

func CreatePaymentController(options PaymentControllerOptions) PaymentInterface {
	return &PaymentController{
		PaymentService: &PaymentService{
			PaymentProvider: &PaymentProvider{
				Db: options.Db,
			},
		},
		Auth: options.Auth,
	}
}

type PaymentController struct {
	PaymentService *PaymentService
	Auth           *oauthPlugin.Auth
}

func (r *PaymentController) GetPayments(w http.ResponseWriter, req *http.Request) {
	var payments []Payment
	var userID int
	var err error
	userID, _ = oauthPlugin.GetUserId(req)

	if payments, err = r.PaymentService.GetPayments(userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)

}
func (r *PaymentController) UpdatePayment(w http.ResponseWriter, req *http.Request) {
	var userID int
	var payment PaymentUpdateDto
	var paymentID int
	var err error
	var body []byte
	userID, _ = oauthPlugin.GetUserId(req)

	if paymentID, err = strconv.Atoi(req.PathValue("ID")); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	if body, err = io.ReadAll(req.Body); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	defer req.Body.Close()

	if err = json.Unmarshal(body, &payment); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if err = payment.Validate(); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if err = r.PaymentService.UpdatePayment(payment, paymentID, userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}
func (r *PaymentController) CreatePayment(w http.ResponseWriter, req *http.Request) {
	var userID int
	var paymentID int
	var payment PaymentCreateDto
	var err error
	var body []byte
	userID, _ = oauthPlugin.GetUserId(req)

	if body, err = io.ReadAll(req.Body); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	defer req.Body.Close()

	if err = json.Unmarshal(body, &payment); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	if err = payment.Validate(); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if paymentID, err = r.PaymentService.CreatePayment(payment, userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentID)
}
func (r *PaymentController) DeletePayments(w http.ResponseWriter, req *http.Request) {
	var userID int
	var err error
	var paymentIDs []int
	var body []byte
	userID, _ = oauthPlugin.GetUserId(req)

	if body, err = io.ReadAll(req.Body); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}
	defer req.Body.Close()

	if err = json.Unmarshal(body, &paymentIDs); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	if err = r.PaymentService.DeletePayments(paymentIDs, userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

func (r *PaymentController) GetPaymentTemplates(w http.ResponseWriter, req *http.Request) {
	var templates []PaymentTemplate
	var userID int
	var err error
	userID, _ = oauthPlugin.GetUserId(req)

	if templates, err = r.PaymentService.GetPaymentTemplates(userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templates)

}
func (r *PaymentController) UpdatePaymentTemplate(w http.ResponseWriter, req *http.Request) {
	var userID int
	var template PaymentTemplateUpdateDto
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

	if err = r.PaymentService.UpdatePaymentTemplate(template, templateID, userID); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}
func (r *PaymentController) CreatePaymentTemplate(w http.ResponseWriter, req *http.Request) {
	var userID int
	var templateID int
	var template PaymentTemplateCreateDto
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

	if templateID, err = r.PaymentService.CreatePaymentTemplate(template, userID); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templateID)
}
func (r *PaymentController) DeletePaymentTemplates(w http.ResponseWriter, req *http.Request) {
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

	if err = r.PaymentService.DeletePaymentTemplates(templateIDs, userID); err != nil {
		web.SendError(w, web.ErrorResponse{Error: err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

func (r *PaymentController) Init(mux *http.ServeMux) {

	mux.Handle("/api/v1/payments", r.Auth.Middleware(true)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			r.CreatePayment(w, req)
		case "GET":
			r.GetPayments(w, req)
		case "DELETE":
			r.DeletePayments(w, req)
		default:
			w.WriteHeader(405)
		}
	})))
	mux.Handle("/api/v1/payments/{ID}", r.Auth.Middleware(true)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "PATCH":
			r.UpdatePayment(w, req)
		default:
			w.WriteHeader(405)
		}
	})))

	mux.Handle("/api/v1/payment_templates", r.Auth.Middleware(true)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			r.CreatePaymentTemplate(w, req)
		case "GET":
			r.GetPaymentTemplates(w, req)
		case "DELETE":
			r.DeletePaymentTemplates(w, req)
		default:
			w.WriteHeader(405)
		}
	})))
	mux.Handle("/api/v1/payment_templates/{ID}", r.Auth.Middleware(true)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "PATCH":
			r.UpdatePaymentTemplate(w, req)
		default:
			w.WriteHeader(405)
		}
	})))
}
