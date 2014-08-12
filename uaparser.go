
package uaparser


import (
    "regexp"
    "strconv"
    "strings"
)


type Id uint

type pattern struct {
    id             Id
    mustContain    []string
    mustNotContain []string
    version        *regexp.Regexp
}


const (
  UNKNOWN Id = 1
)


func find(patterns []pattern, userAgent string) (id Id, version float32) {
  userAgent = strings.ToLower(userAgent)

patterns:
  for _, pattern := range patterns {
    for _, token := range pattern.mustContain {
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

    if len(vs) > 2 {
      if v, err := strconv.ParseFloat(vs[1] + "." + vs[2], 32); err == nil {
        version = float32(v)
      }
    } else if len(vs) == 2 {
      if v, err := strconv.ParseFloat(vs[1], 32); err == nil {
        version = float32(v)
      }
    }

    return
  }

  // Unknown
  id = 1

  return
}

