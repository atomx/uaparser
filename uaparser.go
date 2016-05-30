package uaparser

import (
	"regexp"
	"strconv"
	"strings"
)

type pattern struct {
	id             int64
	mustContain    []string
	mustNotContain []string
	version        *regexp.Regexp
}

const (
	// UNKNOWN is the unknown id.
	UNKNOWN = 1
)

func find(patterns []pattern, userAgent string) (id int64, major int, minor int) {
	userAgent = strings.ToLower(userAgent)

patterns:
	for _, pattern := range patterns {
		for _, token := range pattern.mustContain {
			// TODO: We can speed this up by copying http://golang.org/src/strings/strings.go#L159
			// and precalculating the hashsep.
			if !strings.Contains(userAgent, token) {
				continue patterns
			}
		}

		for _, token := range pattern.mustNotContain {
			if strings.Contains(userAgent, token) {
				continue patterns
			}
		}

		id = pattern.id

		if pattern.version == nil {
			return
		}

		vs := pattern.version.FindStringSubmatch(userAgent)

		if len(vs) > 1 {
			if v, err := strconv.ParseInt(vs[1], 10, 32); err == nil {
				major = int(v)
			}
		}
		if len(vs) > 2 {
			if v, err := strconv.ParseInt(vs[2], 10, 32); err == nil {
				minor = int(v)
			}
		}

		return
	}

	// Unknown
	id = 1

	return
}
