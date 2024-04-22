package user

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/response"
	custom_error "github.com/ansxy/nagabelajar-be-go/pkg/error"
)

func (h *userHandler) UploudFile(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024*10)
	if err := r.ParseMultipartForm(1024 * 1024 * 10); err != nil {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "File too large",
		})
		return
	}

	_, header, err := r.FormFile("file")
	if err != nil {
		response.Error(w, err)
		return
	}

	err = h.uc.UploadFile(r.Context(), header)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)

}
