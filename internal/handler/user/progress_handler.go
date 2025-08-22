package user

import (
	"net/http"
	"strconv"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	"github.com/go-chi/chi/v5"
)

func (h *userHandler) UpdateProgress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UpdateProgressRequest
	req.ProgressID, _ = strconv.Atoi(chi.URLParam(r, "progress_id"))
	req.UserID = r.Context().Value(constant.UserID).(string)

	if err := h.uc.UpdateProgress(ctx, &req); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}
