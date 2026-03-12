package main

import (
	"net/netip"
)

func isIP(s string) (netip.Addr, bool) {
	ip, err := netip.ParseAddr(s)
	return ip, err == nil
}

func isCIDR(s string) (netip.Prefix, bool) {
	cidr, err := netip.ParsePrefix(s)
	return cidr.Masked(), err == nil
}

func getHosts(items []any) []netip.Addr {
	var ips []netip.Addr
	for _, item := range items {
		switch input := item.(type) {
		case netip.Addr:
			ips = append(ips, input)
		case netip.Prefix:
			for ip := input.Addr(); input.Contains(ip); ip = ip.Next() {
				ips = append(ips, ip)
			}
		}
	}
	return ips
}
