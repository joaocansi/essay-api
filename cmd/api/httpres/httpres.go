package httpres

import (
	"encoding/json"
	"net/http"
)

type httpRes struct {
	w http.ResponseWriter
}

func New(w http.ResponseWriter) *httpRes {
	w.Header().Add("Content-Type", "application/json")
	return &httpRes{
		w: w,
	}
}

func (he *httpRes) Send(status int, obj any) {
	body, err := json.Marshal(obj)

	if err != nil {
		he.InternalServerError()
		return
	}

	he.w.WriteHeader(status)
	he.w.Write(body)
}

func (he *httpRes) InternalServerError() {
	he.Error(http.StatusInternalServerError, "internal server error")
}

func (he *httpRes) Error(status int, message string) {
	he.Send(status, map[string]string{"message": message})
}
