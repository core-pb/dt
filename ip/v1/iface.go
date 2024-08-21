package ip

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"errors"
	"fmt"
	"net"
	"net/netip"
)

var (
	_ driver.Valuer              = (*IP)(nil)
	_ sql.Scanner                = (*IP)(nil)
	_ encoding.BinaryMarshaler   = (*IP)(nil)
	_ encoding.BinaryUnmarshaler = (*IP)(nil)
	_ encoding.TextMarshaler     = (*IP)(nil)
	_ encoding.TextUnmarshaler   = (*IP)(nil)

	_ driver.Valuer              = (*AddrPort)(nil)
	_ sql.Scanner                = (*AddrPort)(nil)
	_ encoding.BinaryMarshaler   = (*AddrPort)(nil)
	_ encoding.BinaryUnmarshaler = (*AddrPort)(nil)
	_ encoding.TextMarshaler     = (*AddrPort)(nil)
	_ encoding.TextUnmarshaler   = (*AddrPort)(nil)
)

func (x *IP) MarshalText() ([]byte, error)            { return x.NetIP().MarshalText() }
func (x *IP) MarshalBinary() (data []byte, err error) { return x.NetIP().MarshalBinary() }

func (x *IP) UnmarshalText(text []byte) error {
	var addr netip.Addr
	if err := addr.UnmarshalText(text); err != nil {
		return err
	}
	x.SetIP(addr)
	return nil
}

func (x *IP) UnmarshalBinary(data []byte) error {
	var addr netip.Addr
	if err := addr.UnmarshalBinary(data); err != nil {
		return err
	}
	x.SetIP(addr)
	return nil
}

func (x *IP) Value() (driver.Value, error) {
	if addr := x.NetIP(); addr.IsValid() {
		return addr.String(), nil
	}
	return nil, errors.New("invalid Addr format")
}

func (x *IP) Scan(src any) error {
	switch src := src.(type) {
	case nil:
	case netip.Addr:
		x.SetIP(src)

	case net.IP:
		addr, ok := netip.AddrFromSlice(src)
		if !ok {
			return errors.New("invalid Addr format")
		}
		x.SetIP(addr)

	case []byte:
		addr, ok := netip.AddrFromSlice(src)
		if !ok {
			return errors.New("invalid Addr format")
		}
		x.SetIP(addr)

	case string:
		addr, err := netip.ParseAddr(src)
		if err != nil {
			return err
		}
		x.SetIP(addr)

	default:
		return fmt.Errorf("unexpected type %T", src)
	}
	return nil
}

func (x *AddrPort) Value() (driver.Value, error) { return x.NetAddrPort().String(), nil }

func (x *AddrPort) Scan(src any) error {
	switch src := src.(type) {
	case nil:
	case netip.AddrPort:
		x.SetAddrPort(src)

	case []byte:
		var ap netip.AddrPort
		if err := ap.UnmarshalText(src); err != nil {
			return err
		}
		x.SetAddrPort(ap)

	case string:
		ap, err := netip.ParseAddrPort(src)
		if err != nil {
			return err
		}
		x.SetAddrPort(ap)

	default:
		return fmt.Errorf("unexpected type %T", src)
	}
	return nil
}

func (x *AddrPort) MarshalText() ([]byte, error)            { return x.NetAddrPort().MarshalText() }
func (x *AddrPort) MarshalBinary() (data []byte, err error) { return x.NetAddrPort().MarshalBinary() }

func (x *AddrPort) UnmarshalText(text []byte) error {
	var addr netip.AddrPort
	if err := addr.UnmarshalText(text); err != nil {
		return err
	}
	x.SetAddrPort(addr)
	return nil
}

func (x *AddrPort) UnmarshalBinary(data []byte) error {
	var addr netip.AddrPort
	if err := addr.UnmarshalBinary(data); err != nil {
		return err
	}
	x.SetAddrPort(addr)
	return nil
}
