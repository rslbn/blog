package web

import (
	"encoding/json"
	"net/http"

	customError "github.com/rslbn/blog/internal/errors"
	"github.com/rslbn/blog/internal/util"
)

type handler func(w http.ResponseWriter, r *http.Request) error

func HandlerAdapter(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			sendErrorResponse(w, err)
		}
	}
}

func JSONResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}
	return nil
}

func sendErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	errResponse, _ := util.EncodeJson(customError.ErrorHandler(err))
	w.Write(errResponse)
}
