package query

import (
	"fmt"
)

type IErr[I any] interface {
	Err(error) I
}

type ILimit[I any] interface {
	Limit(int) I
}

type IOffset[I any] interface {
	Offset(int) I
}

type IOrder[I any] interface {
	Order(...string) I
}

type IWhere[I any] interface {
	Where(query string, args ...any) I
}

type IWhereOr[I any] interface {
	WhereOr(query string, args ...any) I
}

func _fmt(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}
