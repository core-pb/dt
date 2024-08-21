package uuid

import (
	"database/sql"
	"database/sql/driver"
	"encoding"

	"github.com/google/uuid"
)

var (
	_ driver.Valuer              = (*UUID)(nil)
	_ sql.Scanner                = (*UUID)(nil)
	_ encoding.BinaryMarshaler   = (*UUID)(nil)
	_ encoding.BinaryUnmarshaler = (*UUID)(nil)
	_ encoding.TextMarshaler     = (*UUID)(nil)
	_ encoding.TextUnmarshaler   = (*UUID)(nil)
)

func (x *UUID) Value() (driver.Value, error) {
	u, err := x.ToUUID()
	if err != nil {
		return nil, err
	}
	return u.Value()
}
func (x *UUID) Scan(src any) error {
	var u uuid.UUID
	if err := u.Scan(src); err != nil {
		return err
	}
	x.FromUUID(u)
	return nil
}

func (x *UUID) MarshalBinary() ([]byte, error) {
	u, err := x.ToUUID()
	if err != nil {
		return nil, err
	}
	return u.MarshalBinary()
}
func (x *UUID) UnmarshalBinary(data []byte) error {
	var u uuid.UUID
	if err := u.UnmarshalBinary(data); err != nil {
		return err
	}
	x.FromUUID(u)
	return nil
}

func (x *UUID) MarshalText() ([]byte, error) {
	u, err := x.ToUUID()
	if err != nil {
		return nil, err
	}
	return u.MarshalText()
}
func (x *UUID) UnmarshalText(text []byte) error {
	var u uuid.UUID
	if err := u.UnmarshalText(text); err != nil {
		return err
	}
	x.FromUUID(u)
	return nil
}
