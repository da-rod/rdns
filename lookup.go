package main

import (
	"fmt"
	"net"
	"net/netip"
	"sort"
)

type result struct {
	addr netip.Addr
	ptr  []string
}

func (r result) print() {
	fmt.Println(r.addr, r.ptr)
}

type results []result

func (r results) print() {
	sort.Slice(r, func(i, j int) bool {
		return r[i].addr.Compare(r[j].addr) < 0
	})
	for _, res := range r {
		if *zero && len(res.ptr) == 0 {
			continue
		}
		res.print()
	}
}

func lookup(input []netip.Addr) results {
	size := len(input)
	ch := make(chan result, size)

	for _, ipaddr := range input {
		ip := ipaddr
		go func() {
			rdns, _ := net.LookupAddr(ip.String())
			ch <- result{
				addr: ip,
				ptr:  rdns,
			}
		}()
	}

	res := make(results, 0, size)
	for range size {
		r := <-ch
		res = append(res, r)
	}
	return res
}
