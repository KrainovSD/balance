package web

import (
	"encoding/json"
	"net/http"
)

func GetProto(r *http.Request, custom string) string {
	var proto string
	var proxyHeader = r.Header[http.CanonicalHeaderKey("x-forwarded-proto")]
	var scheme = r.URL.Scheme

	switch {
	case custom != "":
		proto = custom
	case len(proxyHeader) > 0:
		proto = proxyHeader[0]
	case scheme != "":
		proto = scheme
	case r.TLS != nil:
		proto = "https"
	default:
		proto = "http"
	}

	return proto
}

func GetHost(r *http.Request, custom string) string {
	var host string
	var queryHost = r.URL.Query().Get("frontend_host")

	switch {
	case queryHost != "":
		host = queryHost
	case custom != "":
		host = custom
	default:
		host = r.Host
	}

	return host
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  int    `json:"status"`
	Error   error  `json:"-"`
}

func SendError(w http.ResponseWriter, err ErrorResponse) {
	var message string
	var code int
	var status int

	status = err.Status
	if status == 0 {
		status = 500
	}

	if err.Message != "" {
		message = err.Message + ": " + err.Error.Error()
	} else {
		message = err.Error.Error()
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ErrorResponse{
		Message: message,
		Code:    code,
		Status:  status,
	})

}
