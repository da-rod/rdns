package main

import (
	"strings"
)

func isValid(input string) (bool, string) {
	switch {
	// Check for empty lines and comments
	case len(input) == 0, strings.HasPrefix(input, "#"):
		return false, ""
	// Check if input is an IP addr or an IP cidr
	case isIP(input), isCIDR(input):
		return true, input
	// Check if input is an accepted truncated IP addr (x.y.z)
	// In which case, we assume it means x.y.z.0/24
	case len(strings.Split(input, ".")) == 3:
		input = input + ".0/24"
		return isCIDR(input), input
	default:
		return false, ""
	}
}

func buildHostsList(entries []string) []string {
	var list []string
	for _, entry := range entries {
		switch {
		// Simple case: entry is an ip addr
		case isIP(entry):
			list = append(list, entry)
		// Default: entry is an ip range -> expand it to ip addresses
		default:
			list = append(list, getHosts(entry)...)
		}
	}
	return list
}
