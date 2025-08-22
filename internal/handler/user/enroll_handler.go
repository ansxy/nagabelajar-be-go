package user

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (u *userHandler) CreateEnrollment(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var req request.EnrollCourseRequest
	req.CourseID = chi.URLParam(r, "course_id")
	userID, _ := uuid.Parse(r.Context().Value(constant.UserID).(string))

	req.UserID = userID

	if err := u.uc.EnrollCourse(ctx, &req); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}
