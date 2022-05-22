package handler

import (
	"advdiploma/server/model"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) SecretAdd(w http.ResponseWriter, r *http.Request) {
	// context must contain user id, if not its internal error
	userData := h.GetUserDataFromContext(r)

	// read withdraw from request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	secret, err := h.svcSecret.Add(r.Context(), model.Secret{
		UserID:   userData.UserID,
		DeviceID: userData.DeviceID,
		Data:     string(body),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//  202 if order saved
	w.WriteHeader(http.StatusAccepted)

	if _, err := w.Write([]byte(secret.ID.String())); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
