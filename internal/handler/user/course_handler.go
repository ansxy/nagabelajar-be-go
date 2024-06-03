package user

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (u *userHandler) GetListCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := new(request.ListCourseRequest)
	params.BaseQuery = request.BaseNewQuery(r)
	params.Keyword = r.URL.Query().Get("keyword")
	res, _, err := u.uc.FindListCourse(ctx, params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, res)
}

func (u *userHandler) GetOneCourse(w http.ResponseWriter, r *http.Request) {
	var params *request.GetOneCourseRequest
	ctx := r.Context()
	courseID := chi.URLParam(r, "course_id")
	userID, err := uuid.Parse(r.Context().Value(constant.UserID).(string))
	if err != nil {
		response.Error(w, err)
		return

	}

	params = &request.GetOneCourseRequest{
		CourseID: courseID,
		UserID:   &userID,
	}

	res, err := u.uc.FindOneCourse(ctx, params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, res)
}
