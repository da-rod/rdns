package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"sort"
)

type result struct {
	Addr net.IP
	Ptr  []string
}

type results []result

func main() {
	// Check args
	if len(os.Args) != 2 {
		printHelp()
		os.Exit(1)
	}

	// Parse input for lookup
	var entries []string
	input := os.Args[1]
	switch input {
	case "-":
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

func printHelp() {
	p := os.Args[0]
	fmt.Printf("Usage: %s <IP or CIDR>\n\n", p)
	fmt.Printf("Also accepts IP(s)/CIDR(s) from STDIN using: %s -\n", p)
	fmt.Printf("Example: cat ips.txt | %s -\n", p)
}

func printResults(results results) {
	sort.Slice(results, func(i, j int) bool {
		return bytes.Compare(results[i].Addr, results[j].Addr) < 0
	})
	for _, res := range results {
		fmt.Println(res.Addr, res.Ptr)
	}
}
