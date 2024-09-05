package ut

import (
	"github.com/uptrace/bun"
)

func (x *BunHelp) Where(query string, args ...any) *BunHelp {
	return x.Help(func(q any) {
		switch q := q.(type) {
		case *bun.SelectQuery:
			q.Where(query, args...)
		case *bun.InsertQuery:
			q.Where(query, args...)
		case *bun.UpdateQuery:
			q.Where(query, args...)
		case *bun.DeleteQuery:
			q.Where(query, args...)
		}
	})
}

func (x *BunHelp) WhereOr(query string, args ...any) *BunHelp {
	return x.Help(func(q any) {
		switch q := q.(type) {
		case *bun.SelectQuery:
			q.WhereOr(query, args...)
		case *bun.InsertQuery:
			q.WhereOr(query, args...)
		case *bun.UpdateQuery:
			q.WhereOr(query, args...)
		case *bun.DeleteQuery:
			q.WhereOr(query, args...)
		}
	})
}

func (x *BunHelp) WherePK(cols ...string) *BunHelp {
	return x.Help(func(q any) {
		switch q := q.(type) {
		case *bun.SelectQuery:
			q.WherePK(cols...)
		case *bun.InsertQuery:
			panic("not support")
		case *bun.UpdateQuery:
			q.WherePK(cols...)
		case *bun.DeleteQuery:
			q.WherePK(cols...)
		}
	})
}

func (x *BunHelp) WhereDeleted() *BunHelp {
	return x.Help(func(q any) {
		switch q := q.(type) {
		case *bun.SelectQuery:
			q.WhereDeleted()
		case *bun.InsertQuery:
			panic("not support")
		case *bun.UpdateQuery:
			q.WhereDeleted()
		case *bun.DeleteQuery:
			q.WhereDeleted()
		}
	})
}

func (x *BunHelp) WhereAllWithDeleted() *BunHelp {
	return x.Help(func(q any) {
		switch q := q.(type) {
		case *bun.SelectQuery:
			q.WhereAllWithDeleted()
		case *bun.InsertQuery:
			panic("not support")
		case *bun.UpdateQuery:
			q.WhereAllWithDeleted()
		case *bun.DeleteQuery:
			q.WhereAllWithDeleted()
		}
	})
}
