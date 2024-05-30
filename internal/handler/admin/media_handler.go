package admin

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	custom_error "github.com/ansxy/nagabelajar-be-go/pkg/error"
)

func (h *adminHandler) UploadMedia(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, constant.MAX_FILE_SIZE)
	if err := r.ParseMultipartForm(constant.MAX_FILE_SIZE); err != nil {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "Ukuran file melebihi batas maksimum",
		})

		response.Error(w, err)
		return
	}

	_, header, err := r.FormFile("file")
	if err != nil {
		response.Error(w, err)
		return
	}

	res, err := h.uc.UploadMedia(r.Context(), header)

	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, res)

}
