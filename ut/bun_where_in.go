package ut

import (
	"reflect"
	"strings"

	"github.com/uptrace/bun"
)

func (x *BunHelp) WhereIn(key string, val any) *BunHelp {
	if val == nil {
		return x
	}

	rv := reflect.ValueOf(val)
	if rv.Len() == 1 {
		return x.Where("? = ?", bun.Ident(key), rv.Index(0).Interface())
	}
	return x.Where("? IN (?)", bun.Ident(key), bun.In(val))
}

func (x *BunHelp) WhereInX(key string, val any) *BunHelp {
	if val == nil {
		return x
	}

	rv := reflect.ValueOf(val)
	if rv.Len() == 1 {
		return x.Where(key+" = ?", rv.Index(0).Interface())
	}
	return x.Where(key+" IN (?)", bun.In(val))
}

func (x *BunHelp) WhereEqOrLike(key string, val string) *BunHelp {
	if val == "" || val == "*" {
		return x
	}

	if !strings.ContainsAny(val, "*") {
		x.Where("? = ?", bun.Ident(key), val)
		return x
	}
	x.Where(`? LIKE ?`, bun.Ident(key), strings.ReplaceAll(val, "*", "%"))
	return x
}

func (x *BunHelp) WhereInEqOrLike(key string, val []string) *BunHelp {
	var precise []string

	for i := range val {
		if val[i] == "" || val[i] == "*" {
		} else if strings.ContainsAny(val[i], "*") {
			x.Where(`? LIKE ?`, bun.Ident(key), strings.ReplaceAll(val[i], "*", "%"))
		} else {
			precise = append(precise, val[i])
		}
	}

	switch len(precise) {
	case 0:
	case 1:
		x.WhereOr(`? = ?`, bun.Ident(key), val[0])
	default:
		x.WhereOr(`? IN (?)`, bun.Ident(key), bun.In(val))
	}
	return x
}
