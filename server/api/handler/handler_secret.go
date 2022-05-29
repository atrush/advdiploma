package handler

import (
	apimodel "advdiploma/server/api/model"
	"advdiploma/server/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//  SecretUpload adds or updates secret, returns id and ver
//  if request id is nil adds secret and returns new id and ver 1
//  if request id not nil updates secret and returns id and incremented server version

//  200 - if secret addedd or updated succefully
//  422 - if secret not founded, is deleted, low version to update, request data not valid
//  400 - if cant parse request
//  500 - internal error
func (h *Handler) SecretUpload(w http.ResponseWriter, r *http.Request) {
	user := h.getUserDataFromContext(r)

	var req apimodel.SecretRequest
	if !h.isSecretBodyRead(w, r, &req) {
		return
	}

	secret := model.Secret{
		ID:        req.ID,
		UserID:    user.UserID,
		Ver:       req.Ver,
		Data:      req.Data,
		IsDeleted: false,
	}

	id, ver, err := h.svcSecret.AddUpdate(r.Context(), secret)
	if err != nil {
		h.writeError(w, err)
		return
	}

	resp := apimodel.SecretRequest{
		ID:  id,
		Ver: ver,
	}

	h.writeJSONResponse(w, http.StatusOK, resp)
}

//  SecretDelete deletes secret

//  200 - if deleted succefully
//  422 - if secret not founded, request data not valid
//  400 - if cant parse request
//  500 - internal error
func (h *Handler) SecretDelete(w http.ResponseWriter, r *http.Request) {
	user := h.getUserDataFromContext(r)

	var req apimodel.SecretRequest
	if !h.isSecretBodyRead(w, r, &req) {
		return
	}

	err := h.svcSecret.Delete(r.Context(), req.ID, user.UserID)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeJSONResponse(w, http.StatusOK, nil)
}

//  SecretGet deletes secret

//  200 - if succefully
//  422 - if secret not founded
//  400 - if cant parse request
//  500 - internal error
func (h *Handler) SecretGet(w http.ResponseWriter, r *http.Request) {
	user := h.getUserDataFromContext(r)

	var req apimodel.SecretRequest
	if !h.isSecretBodyRead(w, r, &req) {
		return
	}

	secret, err := h.svcSecret.Get(r.Context(), req.ID, user.UserID)
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeJSONResponse(w, http.StatusOK, secret)
}

func (h *Handler) isSecretBodyRead(w http.ResponseWriter, r *http.Request, data *apimodel.SecretRequest) bool {
	// read withdraw from request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	if err := json.Unmarshal(body, data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	return true
}

//func (h *Handler) SecretGet(w http.ResponseWriter, r *http.Request) {
//	// context must contain user id, if not its internal error
//	userID, err := h.GetUserIDFromContext(r)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	//  202 if order saved
//	w.WriteHeader(http.StatusAccepted)
//
//}
//
//func (h *Handler) SecretDelete(w http.ResponseWriter, r *http.Request) {
//	// context must contain user id, if not its internal error
//	userID, err := h.GetUserIDFromContext(r)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	//  202 if order saved
//	w.WriteHeader(http.StatusAccepted)
//
//}
