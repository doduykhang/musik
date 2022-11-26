package dto

type Pagination struct {
	Page int `schema:"page" validate:"gt=0"`
	Size int `schema:"size" validate:"gt=0"`
}
