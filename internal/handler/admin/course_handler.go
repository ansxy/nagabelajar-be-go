package admin

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
)

func (h *adminHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UpsertCourseRequest

	if err := h.v.ValidateStruct(r, &req); err != nil {
		response.Error(w, err)
		return
	}

	err := h.uc.CreateCourse(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}
