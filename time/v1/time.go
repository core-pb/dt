package time

import (
	"time"
)

func Now() *Time             { return From(time.Now()) }
func From(t time.Time) *Time { return (&Time{}).Set(t) }

func (x *Time) Set(t time.Time) *Time { x.Seconds, x.Nanos = t.Unix(), int32(t.Nanosecond()); return x }
func (x *Time) AsTime() time.Time     { return time.Unix(x.GetSeconds(), int64(x.GetNanos())) }

func (x *Time) Parse(layout, value string) error {
	t, err := time.Parse(layout, value)
	if err != nil {
		return err
	}
	x.Set(t)
	return nil
}
