
package uaparser


import (
    "regexp"
)


var browsers = []pattern{

  // IE < 11
  pattern{
    2,
    []string{"msie"},
    []string{},
    regexp.MustCompile(`msie (\d+\.\d+)`),
  },

  // IE >= 11
  pattern{
    2,
    []string{"trident"},
    []string{},
    regexp.MustCompile(`rv:(\d+\.\d+)`),
  },

  // Chrome
  pattern{
    3,
    []string{"chrome"},
    []string{"chromium"},
    regexp.MustCompile(`chrome/(\d+\.\d+)`),
  },
  // Chrome on iOS
  pattern{
    3,
    []string{"crios"},
    []string{},
    regexp.MustCompile(`crios/(\d+\.\d+)`),
  },

  // Safari
  pattern{
    4,
    []string{"safari"},
    []string{"chrome", "chromium", "crios"},
    regexp.MustCompile(`version/(\d+\.\d+)`),
  },

  // Firefox
  pattern{
    5,
    []string{"firefox"},
    []string{"seamonkey"},
    regexp.MustCompile(`firefox/(\d+\.\d+)`),
  },

  // Opera
  pattern{
    6,
    []string{"opera"},
    []string{},
    regexp.MustCompile(`opera/(\d+\.\d+)`),
  },

}


// Will return the browser id and version number.
// Returns 1,0 when nothing matched.
func Browser(userAgent string) (Id, float32) {
  return find(browsers, userAgent)
}

