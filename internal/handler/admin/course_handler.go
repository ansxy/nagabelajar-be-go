package admin

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/go-chi/chi/v5"
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

func (h *adminHandler) GetListCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := new(request.ListCourseRequest)
	params.BaseQuery = request.BaseNewQuery(r)
	params.Keyword = r.URL.Query().Get("keyword")
	courses, _, err := h.uc.FindListCourse(ctx, params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, courses)
}

func (h *adminHandler) GetOneCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	courseID := chi.URLParam(r, "course_id")
	course, err := h.uc.FindOneCourse(ctx, &request.GetOneCourseRequest{CourseID: courseID})
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, course)
}
