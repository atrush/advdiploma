package handler

import (
	apimodel "advdiploma/server/api/model"
	"advdiploma/server/model"
	"encoding/json"
	"net/http"
)

func (h *Handler) writeError(w http.ResponseWriter, err error) {
	switch err {

	case model.ErrorParamNotValid, model.ErrorItemNotFound, model.ErrorVersionToLow, model.ErrorItemNotFound:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getUserDataFromContext(r *http.Request) apimodel.UserContextData {
	ctxData := r.Context().Value(apimodel.ContextKeyUserID).(apimodel.UserContextData)

	return ctxData
}

func (h *Handler) writeJSONResponse(w http.ResponseWriter, status int, result interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	resp, err := json.Marshal(result)
	if err != nil {
		h.writeError(w, err)
		return
	}

	if _, err := w.Write(resp); err != nil {
		h.writeError(w, err)
		return
	}
}
