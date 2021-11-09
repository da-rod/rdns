package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
)

type result struct {
	addr net.IP
	ptr  []string
}

type results []result

var zero *bool

func init() {
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), "Reverse DNS lookup tool.\n\n")
		fmt.Fprint(flag.CommandLine.Output(), "Takes an IP address or CIDR as argument (or from stdin (one per line)).\n")
		fmt.Fprint(flag.CommandLine.Output(), "(Also accepts \"truncated\" CIDRs x.y.z which are treated as x.y.z.0/24.)\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <IP or CIDR>\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nExamples:\n  - %s 140.82.112.32/29\n  - %s < ips.txt\n  - cat ips.txt | %s\n", os.Args[0], os.Args[0], os.Args[0])
	}
	zero = flag.Bool("z", false, "Filter out IP addresses without reverse DNS")
}

func main() {
	flag.Parse()

	// Parse input for lookup
	var entries []string
	input := flag.Arg(0)
	switch input {
	case "":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			in := scanner.Text()
			if ok, val := isValid(in); ok {
				entries = append(entries, val)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	default:
		if ok, val := isValid(input); ok {
			entries = append(entries, val)
		}
	}

	out := lookup(buildHostsList(entries))
	printResults(out)
}

func printResults(results results) {
	sort.Slice(results, func(i, j int) bool {
		return bytes.Compare(results[i].addr, results[j].addr) < 0
	})
	for _, res := range results {
		if *zero && len(res.ptr) == 0 {
			continue
		}
		fmt.Println(res.addr, res.ptr)
	}
}
