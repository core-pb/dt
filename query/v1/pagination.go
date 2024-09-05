package query

import (
	"errors"
)

var ErrPageTooBig = errors.New("page too big")

func UsePagination[I any](x *Pagination, tx interface {
	IErr[I]
	ILimit[I]
	IOffset[I]
}) I {
	if x == nil {
		return tx.(I)
	}

	page, pageSize := int(x.Page), int(x.PageSize)
	if page < 1 {
		page = 1
	}

	switch {
	case pageSize == -1:
		return tx.(I)
	case pageSize < 1:
		pageSize = 1
	case pageSize > 1000:
		pageSize = 1000
	}

	if page*pageSize > 100*10000 {
		return tx.Err(ErrPageTooBig)
	}

	tx.Limit(pageSize)

	if page != 1 {
		return tx.Offset((page - 1) * pageSize)
	}
	return tx.(I)
}
