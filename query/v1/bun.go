package query

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/known/structpb"
)

// InOrEqSafe use
//
// InOrEqSafe[uint64](nil, bun.Ident("key"), []uint64{1, 2})
// InOrEqSafe[uint64](nil, bun.Ident("key"), []uint64{1, 2}, func(a any) any { return bun.In(a) })
func InOrEqSafe[I, T any](tx IWhere[I], key any, val []T, format ...func(any) any) I {
	switch len(val) {
	case 0:
		return tx.(I)
	case 1:
		return tx.Where(`? = ?`, key, val[0])
	default:
		if len(format) > 0 && format[0] != nil {
			return tx.Where(`? IN (?)`, key, format[0](val))
		}
		return tx.Where(`? IN (?)`, key, val)
	}
}

// InOrEq use
//
// InOrEq[uint64](nil, "key", []uint64{1, 2})
// InOrEq[uint64](nil, "key", []uint64{1, 2}, func(a any) any { return bun.In(a) })
func InOrEq[I, T any](tx IWhere[I], key string, val []T, format ...func(any) any) I {
	switch len(val) {
	case 0:
		return tx.(I)
	case 1:
		return tx.Where(_fmt(`%s = ?`, key), val[0])
	default:
		if len(format) > 0 && format[0] != nil {
			return tx.Where(_fmt(`%s IN (?)`, key), format[0](val))
		}
		return tx.Where(_fmt(`%s IN (?)`, key), val)
	}
}

func OrLikeEq[I any](tx IWhereOr[I], key string, val []string, format ...func(any) any) I {
	var precise []string

	for i := range val {
		if val[i] == "" || val[i] == "*" {
			continue
		}
		if strings.ContainsAny(val[i], "*") {
			tx.WhereOr(_fmt(`%s LIKE ?`, key), strings.ReplaceAll(val[i], "*", "%"))
			continue
		}

		precise = append(precise, val[i])
	}

	switch len(precise) {
	case 0:
		return tx.(I)
	case 1:
		return tx.WhereOr(_fmt(`%s = ?`, key), val[0])
	default:
		if len(format) > 0 && format[0] != nil {
			return tx.WhereOr(_fmt(`%s IN (?)`, key), format[0](val))
		}
		return tx.WhereOr(_fmt(`%s IN (?)`, key), val)
	}
}

func WhereStruct[I any](tx interface {
	IWhere[I]
	IErr[I]
}, qk string, val *structpb.Struct) I {
	for k, v := range val.GetFields() {
		switch v := v.Kind.(type) {
		case nil:
			continue
		case *structpb.Value_NullValue:
			tx.Where(fmt.Sprintf("%s->'%s' = 'null'", qk, k))
			continue
		case *structpb.Value_StructValue:
			WhereStruct(tx, fmt.Sprintf("%s->'%s'", qk, k), v.StructValue)
			continue
		case *structpb.Value_ListValue:
			if len(v.ListValue.Values) == 0 {
				continue
			}
		}

		b, err := v.MarshalJSON()
		if err != nil {
			return tx.Err(fmt.Errorf("json: marshal err, %s", err))
		}
		tx.Where(fmt.Sprintf("%s->'%s' = '%s'", qk, k, string(b)))
	}

	return tx.(I)
}
