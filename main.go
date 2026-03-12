package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var zero *bool

func init() {
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprint(out, "Reverse DNS lookup tool.\n\n")
		fmt.Fprint(out, "Takes an IP address or CIDR as argument (or from stdin (one per line)).\n")
		fmt.Fprint(out, "(Also accepts \"truncated\" CIDRs x.y.z which are treated as x.y.z.0/24.)\n\n")
		fmt.Fprintf(out, "Usage: %s <IP or CIDR>\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(out, "\nExamples:\n  - %s 140.82.112.32/29\n  - %s < ips.txt\n  - cat ips.txt | %s\n", os.Args[0], os.Args[0], os.Args[0])
	}
	zero = flag.Bool("z", false, "Filter out IP addresses without reverse DNS")
}

func main() {
	flag.Parse()

	// Parse input for lookup
	var entries []any

	switch input := flag.Arg(0); input {
	case "", "-":
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if val, ok := isValid(scanner.Text()); ok {
				entries = append(entries, val)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	default:
		if val, ok := isValid(input); ok {
			entries = append(entries, val)
		}
	}

	ips := getHosts(entries)
	res := lookup(ips)
	res.print()
}

func isValid(input string) (any, bool) {
	// remove empty/comment lines / trim inline comments
	input, _, _ = strings.Cut(input, "#")
	input = strings.TrimSpace(input)
	// refang defanged strings
	input = strings.ReplaceAll(input, "[.]", ".")
	// accept truncated cidr x.y.z, consider it as x.y.z.0/24
	octets := strings.Split(input, ".")
	switch {
	case len(octets) == 4 && octets[len(octets)-1] == "":
		input = strings.TrimSuffix(input, ".")
		fallthrough
	case len(octets) == 3:
		input = input + ".0/24"
	}
	if ip, ok := isIP(input); ok {
		return ip, ok
	}
	if cidr, ok := isCIDR(input); ok {
		return cidr, ok
	}
	return nil, false
}
