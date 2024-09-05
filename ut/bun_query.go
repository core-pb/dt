package ut

import (
	"errors"
	"fmt"

	"github.com/core-pb/dt/query/v1"
	"github.com/uptrace/bun"
	"go.x2ox.com/sorbifolia/pyrokinesis"
	"google.golang.org/protobuf/types/known/structpb"
)

func (x *BunHelp) Pagination(p *query.Pagination) *BunHelp {
	if p == nil {
		return x
	}

	var (
		tx             = x.query.(*bun.SelectQuery)
		page, pageSize = int(p.Page), int(p.PageSize)
	)
	if page < 1 {
		page = 1
	}

	switch {
	case pageSize == -1:
		return x
	case pageSize < 1:
		pageSize = 1
	case pageSize > 1000:
		pageSize = 1000
	}

	switch {
	case page*pageSize > 100*10000: // query cost to big
		tx.Err(errors.New("page too big"))
	case page == 1:
		tx.Limit(pageSize)
	default:
		tx.Limit(pageSize).Offset((page - 1) * pageSize)
	}

	return x
}

func (x *BunHelp) Sort(s []*query.Sort) *BunHelp {
	arr := make([]string, 0, len(s))

	for _, v := range s {
		var (
			desc bool
			key  string
		)
		switch v := v.GetKey().(type) {
		case *query.Sort_Desc:
			desc, key = true, v.Desc
		case *query.Sort_Asc:
			key = v.Asc
		}

		switch {
		case key == "":
		case desc:
			arr = append(arr, key+" DESC")
		default:
			arr = append(arr, key)
		}
	}

	x.query.(*bun.SelectQuery).Order(arr...)
	return x
}

func (x *BunHelp) QueryStruct(qk string, val *structpb.Struct) *BunHelp {
	for k, v := range val.GetFields() {
		switch v := v.Kind.(type) {
		case nil:
			continue
		case *structpb.Value_NullValue:
			x.Where(fmt.Sprintf("%s->'%s' = 'null'", qk, k))
			continue
		case *structpb.Value_StructValue:
			x.QueryStruct(fmt.Sprintf("%s->'%s'", qk, k), v.StructValue)
			continue
		case *structpb.Value_ListValue:
			if len(v.ListValue.Values) == 0 {
				continue
			}
		}

		b, err := v.MarshalJSON()
		if err != nil {
			return x.Err(err)
		}
		x.Where(fmt.Sprintf("%s->'%s' = '%s'", qk, k, pyrokinesis.Bytes.ToString(b)))
	}
	return x
}
