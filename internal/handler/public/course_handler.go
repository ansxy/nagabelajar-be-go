package public

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/go-chi/chi/v5"
)

func (h *publicHandler) GetListCourse(w http.ResponseWriter, r *http.Request) {
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

func (h *publicHandler) GetOneCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := new(request.GetOneCourseRequest)
	params.CourseID = chi.URLParam(r, "course_id")
	course, err := h.uc.FindOneCourse(ctx, params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, course)

}
