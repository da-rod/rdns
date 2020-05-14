package main

import "net"

func isIP(s string) bool {
	return net.ParseIP(s) != nil
}

func isCIDR(s string) bool {
	_, _, err := net.ParseCIDR(s)
	return err == nil
}

func getHosts(network string) []string {
	var ips []string
	ip, iprange, _ := net.ParseCIDR(network)
	for ip := ip.Mask(iprange.Mask); iprange.Contains(ip); increment(ip) {
		ips = append(ips, ip.String())
	}
	return ips
}

func increment(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
