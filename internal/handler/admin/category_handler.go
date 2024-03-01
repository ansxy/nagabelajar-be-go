package admin

import (
	"net/http"
	"strconv"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/go-chi/chi/v5"
)

func (h *adminHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UpsertCategoryRequest

	if err := h.v.ValidateStruct(r, &req); err != nil {
		response.Error(w, err)
		return
	}

	err := h.uc.CreateCategory(ctx, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *adminHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query, err := strconv.Atoi(chi.URLParam(r, "category_id"))

	if err != nil {
		response.Error(w, err)
		return

	}
	err = h.uc.DeleteOneCategory(ctx, query)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *adminHandler) FindListCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var params request.ListCategoryRequest
	params.BaseQuery = request.BaseNewQuery(r)

	res, cnt, err := h.uc.FindListCategory(ctx, &params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Pagination(w, res, params.Page, params.PerPage, cnt)
}
