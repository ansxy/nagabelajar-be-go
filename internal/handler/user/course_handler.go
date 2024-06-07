package user

import (
	"net/http"
	"strconv"

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
	userId := r.Context().Value(constant.UserID).(string)
	params.UserID = userId
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
	userID := r.Context().Value(constant.UserID)
	userid, err := uuid.Parse(userID.(string))

	if err != nil {
		response.Error(w, err)
		return
	}

	params = &request.GetOneCourseRequest{
		CourseID: courseID,
		UserID:   &userid,
	}

	res, err := u.uc.FindOneCourse(ctx, params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, res)
}

func (h *userHandler) FindCourseDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	courseDetailID, _ := strconv.Atoi(chi.URLParam(r, "course_detail_id"))
	courseDetail, err := h.uc.FindCourseDetail(ctx, courseDetailID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, courseDetail)
}
