package ut

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type BunHelp struct {
	idb   bun.IDB
	query any
}

func Bun(idb bun.IDB) *BunHelp { return &BunHelp{idb: idb} }

func (x *BunHelp) Select() *BunHelp { x.query = x.idb.NewSelect(); return x }
func (x *BunHelp) Insert() *BunHelp { x.query = x.idb.NewInsert(); return x }
func (x *BunHelp) Update() *BunHelp { x.query = x.idb.NewUpdate(); return x }
func (x *BunHelp) Delete() *BunHelp { x.query = x.idb.NewDelete(); return x }

func (x *BunHelp) RawSelect() *bun.SelectQuery { return x.query.(*bun.SelectQuery) }
func (x *BunHelp) RawInsert() *bun.InsertQuery { return x.query.(*bun.InsertQuery) }
func (x *BunHelp) RawUpdate() *bun.UpdateQuery { return x.query.(*bun.UpdateQuery) }
func (x *BunHelp) RawDelete() *bun.DeleteQuery { return x.query.(*bun.DeleteQuery) }

func (x *BunHelp) Help(f func(q any)) *BunHelp { f(x.query); return x }

func (x *BunHelp) Model(model any) *BunHelp {
	return x.Help(func(q any) {
		switch q := q.(type) {
		case *bun.SelectQuery:
			q.Model(model)
		case *bun.InsertQuery:
			q.Model(model)
		case *bun.UpdateQuery:
			q.Model(model)
		case *bun.DeleteQuery:
			q.Model(model)
		}
	})
}

func (x *BunHelp) Table(table ...string) *BunHelp {
	return x.Help(func(q any) {
		switch q := q.(type) {
		case *bun.SelectQuery:
			q.Table(table...)
		case *bun.InsertQuery:
			q.Table(table...)
		case *bun.UpdateQuery:
			q.Table(table...)
		case *bun.DeleteQuery:
			q.Table(table...)
		}
	})
}

func (x *BunHelp) Exec(ctx context.Context, dest ...any) (sql.Result, error) {
	switch q := x.query.(type) {
	case *bun.SelectQuery:
		return q.Exec(ctx, dest...)
	case *bun.InsertQuery:
		return q.Exec(ctx, dest...)
	case *bun.UpdateQuery:
		return q.Exec(ctx, dest...)
	case *bun.DeleteQuery:
		return q.Exec(ctx, dest...)
	}
	panic("not implemented")
}

func (x *BunHelp) Scan(ctx context.Context, dest ...any) error {
	switch q := x.query.(type) {
	case *bun.SelectQuery:
		return q.Scan(ctx, dest...)
	case *bun.InsertQuery:
		return q.Scan(ctx, dest...)
	case *bun.UpdateQuery:
		return q.Scan(ctx, dest...)
	case *bun.DeleteQuery:
		return q.Scan(ctx, dest...)
	}
	panic("not implemented")
}

func (x *BunHelp) Err(err error) *BunHelp {
	switch q := x.query.(type) {
	case *bun.SelectQuery:
		q.Err(err)
	case *bun.InsertQuery:
		q.Err(err)
	case *bun.UpdateQuery:
		q.Err(err)
	case *bun.DeleteQuery:
		q.Err(err)
	}
	return x
}
