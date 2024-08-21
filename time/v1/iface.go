package time

import (
	"database/sql"
	"encoding"
	"encoding/json"
	"time"
)

var (
	_ sql.Scanner = (*Time)(nil)

	_ json.Marshaler             = (*Time)(nil)
	_ json.Unmarshaler           = (*Time)(nil)
	_ encoding.BinaryMarshaler   = (*Time)(nil)
	_ encoding.BinaryUnmarshaler = (*Time)(nil)
	_ encoding.TextMarshaler     = (*Time)(nil)
	_ encoding.TextUnmarshaler   = (*Time)(nil)
)

func (x *Time) Scan(src any) error {
	switch src := src.(type) {
	case nil:
	case string:
		return x.Parse(time.RFC3339Nano, src)

	case []byte:
		t := &time.Time{}
		if err := t.UnmarshalText(src); err != nil {
			return err
		}
		x.Set(*t)
	}

	return nil
}

func (x *Time) MarshalJSON() ([]byte, error)   { return x.AsTime().MarshalJSON() }
func (x *Time) MarshalBinary() ([]byte, error) { return x.AsTime().MarshalBinary() }
func (x *Time) MarshalText() ([]byte, error)   { return x.AsTime().MarshalText() }

func (x *Time) UnmarshalJSON(b []byte) error {
	var t time.Time
	if err := t.UnmarshalJSON(b); err != nil {
		return err
	}
	x.Set(t)
	return nil
}

func (x *Time) UnmarshalBinary(b []byte) error {
	var t time.Time
	if err := t.UnmarshalBinary(b); err != nil {
		return err
	}
	x.Set(t)
	return nil
}
func (x *Time) UnmarshalText(b []byte) error {
	var t time.Time
	if err := t.UnmarshalText(b); err != nil {
		return err
	}
	x.Set(t)
	return nil
}
