package handler

import (
	apimodel "advdiploma/server/api/model"
	"net/http"
)

func (h *Handler) SyncList(w http.ResponseWriter, r *http.Request) {
	user := h.getUserDataFromContext(r)

	list, err := h.svcSecret.GetUserSyncList(r.Context(), user.UserID)
	if err != nil {
		h.writeError(w, err)
	}

	result := apimodel.SyncResponse{
		List: list,
	}

	h.writeJSONResponse(w, http.StatusOK, result)
}
