package utils

import (
	"encoding/json"
	"net/http"
)

type HTTPSetup struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

type HTTPResponse struct {
	Code   int         `json:"code"`
	Error  string      `json:"error,omitempty"`
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (h HTTPSetup) setHeader() {
	h.Writer.Header().Set("Server", Config.App.Name+" "+Config.App.Version)
	h.Writer.Header().Set("Version", Config.App.Version)
	if Config.Security.EnabledCors {
		switch origin := h.Request.Header.Get("Origin"); origin {
		case Config.Security.WhitelistHost:
			h.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			h.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			h.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, secure")
		}
	} else {
		h.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		h.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		h.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, secure")
	}
	h.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	h.Writer.Header().Set("Content-Type", "application/json")
}

func (h HTTPSetup) ErrorResp(m string) {
	h.setHeader()
	h.Writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(h.Writer).Encode(HTTPResponse{
		Error: m,
		Code:  http.StatusBadRequest,
	})
}

func (h HTTPSetup) SuccessResp(resp interface{}) {
	h.setHeader()
	h.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(h.Writer).Encode(resp)
}
