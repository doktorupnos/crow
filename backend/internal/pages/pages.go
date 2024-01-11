package pages

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type PaginationParams struct {
	PageNumber int
	PageSize   int
}

func Paginate(p PaginationParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := p.PageNumber * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}

func ExtractPage(r *http.Request) int {
	p, _ := strconv.Atoi(r.URL.Query().Get("page"))
	return p
}

func ExtractLimit(r *http.Request) int {
	l, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	return l
}
