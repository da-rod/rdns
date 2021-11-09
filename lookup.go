package main

import "net"

func lookup(input []string) results {
	size := len(input)
	ch := make(chan result, size)

	for _, ipaddr := range input {
		ip := ipaddr
		go func() {
			rdns, _ := net.LookupAddr(ip)
			ch <- result{
				addr: net.ParseIP(ip),
				ptr:  rdns,
			}
		}()
	}

	res := make(results, 0, size)
	for i := 0; i < size; i++ {
		r := <-ch
		res = append(res, r)
	}
	return res
}
