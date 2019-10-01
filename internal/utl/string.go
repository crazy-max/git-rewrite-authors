package utl

import (
	"errors"
	"fmt"
	"regexp"
)

// ParseAddress parses a single RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"
func ParseAddress(address string) (name string, email string, err error) {
	re := regexp.MustCompile("(.*) <(.*)>")
	match := re.FindStringSubmatch(address)
	if match == nil || len(match) != 3 {
		return "", "", errors.New(fmt.Sprintf("Cannot parse %s", address))
	}
	return match[1], match[2], nil
}
