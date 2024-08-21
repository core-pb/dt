package ip

import (
	"fmt"
	"net/netip"
)

func (x *IP) IsIPv4() bool { return x.GetV4() != nil }
func (x *IP) IsIPv6() bool { return x.GetV6() != nil }

func (x *IP) NetIP() netip.Addr {
	if v := x.GetV4(); v != nil {
		return netip.AddrFrom4([4]byte{
			byte(v.Ip >> 24), byte(v.Ip >> 16), byte(v.Ip >> 8), byte(v.Ip),
		})
	}
	if v := x.GetV6(); v != nil {
		return netip.AddrFrom16([16]byte{
			byte(v.Hi >> 56), byte(v.Hi >> 48), byte(v.Hi >> 40), byte(v.Hi >> 32),
			byte(v.Hi >> 24), byte(v.Hi >> 16), byte(v.Hi >> 8), byte(v.Hi),

			byte(v.Lo >> 56), byte(v.Lo >> 48), byte(v.Lo >> 40), byte(v.Lo >> 32),
			byte(v.Lo >> 24), byte(v.Lo >> 16), byte(v.Lo >> 8), byte(v.Lo),
		})
	}
	return netip.Addr{}
}

func (x *IP) SetIP(ip netip.Addr) {
	b := ip.AsSlice()

	switch len(b) {
	case 4:
		x.Addr = &IP_V4{V4: &IPv4{
			Ip: uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24,
		}}
	case 16:
		x.Addr = &IP_V6{V6: &IPv6{
			Hi: uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
				uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56,
			Lo: uint64(b[15]) | uint64(b[14])<<8 | uint64(b[13])<<16 | uint64(b[12])<<24 |
				uint64(b[11])<<32 | uint64(b[10])<<40 | uint64(b[9])<<48 | uint64(b[8])<<56,
		}}
	default:
		panic(fmt.Sprintf("invalid IP address: %s", ip))
	}
}

func (x *AddrPort) NetAddrPort() netip.AddrPort {
	return netip.AddrPortFrom(x.Addr.NetIP(), uint16(x.Port))
}

func (x *AddrPort) SetAddrPort(ap netip.AddrPort) {
	x.Addr.SetIP(ap.Addr())
	x.Port = uint32(ap.Port())
}

func (x *Prefix) NetPrefix() netip.Prefix {
	return netip.PrefixFrom(x.Addr.NetIP(), int(x.Bits))
}

func (x *Prefix) SetPrefix(p netip.Prefix) {
	x.Addr.SetIP(p.Addr())
	x.Bits = int32(p.Bits())
}
