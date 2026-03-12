package main

import (
	"context"
	"fmt"
	"net"
	"net/netip"
	"sort"
	"time"
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

	var resolver net.Resolver
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, ipaddr := range input {
		ip := ipaddr
		go func() {
			rdns, _ := resolver.LookupAddr(ctx, ip.String())
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
