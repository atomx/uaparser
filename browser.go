package uaparser

import (
	"regexp"
)

// Browser int's.
const (
	IE        = iota + 2 // 2
	CHROME               // 3
	SAFARI               // 4
	FIREFOX              // 5
	OPERA                // 6
	CHROMIUM             // 7
	UCBROWSER            // 8
)

var Browsers = map[int]string{
	UNKNOWN:   "Unknown",
	IE:        "IE",
	CHROME:    "Chrome",
	SAFARI:    "Safari",
	FIREFOX:   "Firefox",
	OPERA:     "Opera",
	CHROMIUM:  "Chromium",
	UCBROWSER: "UC Browser",
}

var browsers = []pattern{

	// IE < 11
	pattern{
		IE,
		[]string{"msie"},
		[]string{"chromeframe"},
		regexp.MustCompile(`msie (\d+)\.(\d+)`),
	},

	// IE 11
	pattern{
		IE,
		[]string{"trident"},
		[]string{"chromeframe"},
		regexp.MustCompile(`rv:(\d+)\.(\d+)`),
	},

	// IE 12
	// Looks like:
	// Mozilla/5.0 (Windows NT 6.4; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Safari/537.36 Edge/12.0
	pattern{
		IE,
		[]string{"applewebkit", "chrome", "safari", "edge"},
		[]string{},
		regexp.MustCompile(`edge/(\d+)\.(\d+)`),
	},

	// Chrome
	pattern{
		CHROME,
		[]string{"chrome"},
		[]string{"chromium", "chromeframe", "edge", "vivaldi"}, // Edge is from IE12 preview.
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
		[]string{"chrome", "chromium", "crios", "ucbrowser", "qqbrowser"},
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

	// UC Browser
	// See: https://play.google.com/store/apps/details?id=com.UCMobile.intl&hl=en
	// According to http://gs.statcounter.com/#mobile_browser-ww-monthly-201402-201502-bar it's the 5th biggest browser
	// with about the same number of users as Opera.
	pattern{
		UCBROWSER,
		[]string{"ucbrowser"},
		[]string{},
		regexp.MustCompile(`ucbrowser/(\d+)\.(\d+)`),
	},
}

// Browser will return the browser id and version number.
// Returns 1,0 when nothing matched.
func Browser(userAgent string) (int, int) {
	return find(browsers, userAgent)
}
