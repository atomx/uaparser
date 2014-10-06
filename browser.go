package uaparser

import (
	"regexp"
)

// Browser ID's.
const (
	IE       ID = iota + 2 // 2
	CHROME                 // 3
	SAFARI                 // 4
	FIREFOX                // 5
	OPERA                  // 6
	CHROMIUM               // 7
)

var browsers = []pattern{

	// IE < 11
	pattern{
		IE,
		[]string{"msie"},
		[]string{"chromeframe"},
		regexp.MustCompile(`msie (\d+)\.(\d+)`),
	},

	// IE >= 11
	pattern{
		IE,
		[]string{"trident"},
		[]string{"chromeframe"},
		regexp.MustCompile(`rv:(\d+)\.(\d+)`),
	},

	// Chrome
	pattern{
		CHROME,
		[]string{"chrome"},
		[]string{"chromium", "chromeframe"},
		regexp.MustCompile(`chrome/(\d+)\.(\d+)`),
	},
	// Chrome on iOS
	pattern{
		CHROME,
		[]string{"crios"},
		[]string{},
		regexp.MustCompile(`crios/(\d+)\.(\d+)`),
	},

	// Safari
	pattern{
		SAFARI,
		[]string{"safari"},
		[]string{"chrome", "chromium", "crios"},
		regexp.MustCompile(`version/(\d+)\.(\d+)`),
	},

	// Firefox
	pattern{
		FIREFOX,
		[]string{"firefox"},
		[]string{"seamonkey"},
		regexp.MustCompile(`firefox/(\d+)\.(\d+)`),
	},

	// Opera
	pattern{
		OPERA,
		[]string{"opera"},
		[]string{},
		regexp.MustCompile(`opera/(\d+)\.(\d+)`),
	},

	// Chromium
	pattern{
		CHROMIUM,
		[]string{"chromium"},
		[]string{},
		regexp.MustCompile(`chromium/(\d+)\.(\d+)`),
	},
}

// Browser will return the browser id and version number.
// Returns 1,0 when nothing matched.
func Browser(userAgent string) (ID, int) {
	return find(browsers, userAgent)
}
