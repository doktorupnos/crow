package pages

import "gorm.io/gorm"

type PaginationParams struct {
	PageNumber int
	PageSize   int
}

func Paginate(params PaginationParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (params.PageNumber - 1) * params.PageSize
		return db.Offset(offset).Limit(params.PageSize)
	}
}
