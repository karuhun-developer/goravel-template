package v1

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type PaginationRequest struct {
	Page *int `form:"page" json:"page" query:"page"`
	Paginate *int `form:"paginate" json:"paginate" query:"paginate"`
	Search *string `form:"search" json:"search" query:"search"`
	SearchBy *string `form:"search_by" json:"search_by" query:"search_by"`
	Order *string `form:"order" json:"order" query:"order"`
	OrderBy *string `form:"order_by" json:"order_by" query:"order_by"`
}

func (r *PaginationRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *PaginationRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PaginationRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"page":     "required_with:page|integer|min:1",
		"paginate": "required_with:paginate|integer|min:1|max:100",
		"search":   "required_with:search|string",
		"search_by":"required_with:search_by|string",
		"order":    "required_with:order|string|in:asc,desc",
		"order_by": "required_with:order_by|string",
	}
}

func (r *PaginationRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PaginationRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PaginationRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
