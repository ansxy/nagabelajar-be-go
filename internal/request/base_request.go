package request

import (
	"log"
	"net/http"
	"strconv"
)

type BaseQuery struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Keyword string `json:"keyword"`
	Sort    string `json:"sort"`
}

const (
	DEFAULT_PAGE    = 1
	DEFAULT_PERPAGE = 10
)

func BaseNewQuery(r *http.Request) BaseQuery {
	page := DEFAULT_PAGE
	perPage := DEFAULT_PERPAGE

	log.Println(r.URL.Query().Get("page"))
	if r.URL.Query().Get("page") != "" {
		page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	}

	if r.URL.Query().Get("per_page") != "" {
		perPage, _ = strconv.Atoi(r.URL.Query().Get("per_page"))
	}

	return BaseQuery{
		Page:    page,
		PerPage: perPage,
		Keyword: r.URL.Query().Get("keyword"),
		Sort:    r.URL.Query().Get("sort"),
	}
}
