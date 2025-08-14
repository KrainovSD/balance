package lib

import (
	"encoding/json"
	"net/http"
	"os"
)

func GetProto(r *http.Request) string {
	var proto string
	var frontendProtocol = os.Getenv("FRONTEND_PROTOCOL")
	var proxyHeader = r.Header[http.CanonicalHeaderKey("x-forwarded-proto")]
	var scheme = r.URL.Scheme

	switch {
	case frontendProtocol != "":
		proto = frontendProtocol
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

func GetHost(r *http.Request) string {
	var host string
	var frontedHost = os.Getenv("FRONTEND_HOST")
	var queryHost = r.URL.Query().Get("frontend_host")

	switch {
	case queryHost != "":
		host = queryHost
	case frontedHost != "":
		host = frontedHost
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
