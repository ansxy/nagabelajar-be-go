package admin

import (
	"net/http"
	"strconv"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/go-chi/chi/v5"
)

func (h *adminHandler) CreateCourseDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UpserCourseDetailRequest
	req.CourseID = chi.URLParam(r, "course_id")
	if err := h.v.ValidateStruct(r, &req); err != nil {
		response.Error(w, err)
		return
	}

	err := h.uc.CreateCourseDetail(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *adminHandler) FindCourseDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	courseID, _ := strconv.Atoi(chi.URLParam(r, "course_id"))
	courseDetail, err := h.uc.FindCourseDetail(ctx, courseID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, courseDetail)
}
