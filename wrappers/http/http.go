package fueltilityhttp

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const (
	ContentTypeHeader   = "Content-Type"
	JSONContentType     = "application/json"
	ControlOriginHeader = "Access-Control-Allow-Origin"
)

type Response[T any] struct {
	Success bool `json:"success"`
	Data    []T  `json:"data"`
}

type ErrorResponse struct {
	Success bool          `json:"success"`
	Error   *ErrorDetails `json:"error,omitempty"`
}

type ErrorDetails struct {
	Message string `json:"message"`
}

type ResponseWriter struct {
	W http.ResponseWriter
}

func (crw *ResponseWriter) SetCors(origin string) {
	if strings.Contains(origin, "www.fueltility.com") {
		crw.W.Header().Set(ControlOriginHeader, "https://fueltility.com")
	}

	if strings.Contains(origin, "127.0.0.1") {
		crw.W.Header().Set(ControlOriginHeader, "http://localhost:3000")
	}

	crw.W.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	crw.W.Header().Set("Access-Control-Allow-Methods", "*")
}

func (crw *ResponseWriter) SendJSONResponse(status int, payload interface{}) {
	crw.W.Header().Set(ContentTypeHeader, JSONContentType)
	crw.W.WriteHeader(status)

	jsonResp, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return
	}

	if _, err := crw.W.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
