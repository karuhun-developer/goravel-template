package helpers

import (
	"errors"
	"fmt"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"karuhundeveloper.com/gostarterkit/app/http/responses"
	"karuhundeveloper.com/gostarterkit/app/models/role"
)

func OrmFilter(ctx http.Context, query orm.Query, fields []string) orm.Query {
	search := ctx.Request().Query("search", "")

	// Apply if search query is provided
	if search != "" {
		// If search by is provided
		searchBy := ctx.Request().Query("search_by", "")

		if searchBy != "" {
			query = query.Where(searchBy + " LIKE ?", "%"+search+"%")
		} else {
			// Search in all fields
			for _, field := range fields {
				query = query.OrWhere(field + " LIKE ?", "%"+search+"%")
			}
		}
	}

	// Order by
	orderBy := ctx.Request().Query("order_by", "id")
	orderDirection := ctx.Request().Query("order", "desc")

	query = query.Order(orderBy + " " + orderDirection)

	fmt.Println(query.ToRawSql().Get(role.Role{}))

	return query
}

func PaginateHelper(page int, paginate int, total int64) (pagination responses.PaginationResponse, err error) {
	// Set pagination response
	lastPage := int((total + int64(paginate) - 1) / int64(paginate))
	nextPage := 0
	if page < lastPage {
		nextPage = page + 1
	}
	prevPage := 0
	if page > 1 && page <= lastPage {
		prevPage = page - 1
	}

	// If page exceeds last page, return error
	if page > lastPage && lastPage != 0 {
		err = errors.New("page exceeds last page")
		return
	}

	pagination = responses.PaginationResponse{
		Total:       total,
		PerPage:     paginate,
		CurrentPage: page,
		LastPage:    lastPage,
		NextPage:    nextPage,
		PrevPage:    prevPage,
	}

	return
}