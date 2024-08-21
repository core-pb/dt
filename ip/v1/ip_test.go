package ip

import (
	"net/netip"
	"testing"
)

func TestIP_SetIP(t *testing.T) {
	for _, v := range []string{
		"192.168.0.1", "0.0.0.0", "255.255.255.255",
		"2400::1",
	} {
		ip, _ := netip.ParseAddr(v)
		x := &IP{}
		x.SetIP(ip)

		if x.NetIP() != ip {
			t.Errorf("x.NetIP()=%v, want %v", x.NetIP(), ip)
		}
	}
}
