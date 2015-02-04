package uaparser

import (
	"regexp"
	"strconv"
	"strings"
)

// ID type.
type ID uint

type pattern struct {
	id             ID
	mustContain    []string
	mustNotContain []string
	version        *regexp.Regexp
}

const (
	// UNKNOWN is the unknown ID.
	UNKNOWN ID = 1
)

// Version converts a major,minor version pair into a single number using (major * 10000) + minor.
func Version(major, minor int) int {
	return (major * 10000) + minor
}

func Unversion(version int) (int, int) {
	minor := version % 10000
	return (version - minor) / 10000, minor
}

func find(patterns []pattern, userAgent string) (id ID, version int) {
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
		major := 0
		minor := 0

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

		version = Version(major, minor)

		return
	}

	// Unknown
	id = 1

	return
}
