package public

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	custom_error "github.com/ansxy/nagabelajar-be-go/pkg/error"
)

func (h *publicHandler) ValidateCertificate(w http.ResponseWriter, r *http.Request) {
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

	err = h.uc.ValidateCertificate(r.Context(), header)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *publicHandler) CreateCertificate(w http.ResponseWriter, r *http.Request) {
	var req request.CreateCertificateRequest
	req.FirebaseID = r.Context().Value(constant.FirebaseID).(string)

	if err := h.v.ValidateStruct(r, &req); err != nil {
		response.Error(w, err)
		return
	}

	err := h.uc.CreateCertificate(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *publicHandler) ValidateCertificateByAddress(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")

	res, err := h.uc.ValidateCertificateByAddress(r.Context(), address)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, res)
}

func (h *publicHandler) GetListCertificate(w http.ResponseWriter, r *http.Request) {
	var req request.ListCertificateRequest
	req.BaseQuery = request.BaseNewQuery(r)
	req.PerPage = -1
	req.Page = 1
	req.FirebaseID = r.Context().Value(constant.FirebaseID).(string)

	res, count, err := h.uc.GetListCertificate(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Pagination(w, res, req.Page, req.PerPage, count)
}
