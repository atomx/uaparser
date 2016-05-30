package uaparser

import (
	"regexp"
)

// Browser id's.
const (
	IE             = iota + 2 // 2
	CHROME                    // 3
	SAFARI                    // 4
	FIREFOX                   // 5
	OPERA                     // 6
	CHROMIUM                  // 7
	UCBROWSER                 // 8
	ANDROIDBROWSER            // 9
	BLACKBERRY                // 10
	FACEBOOK                  // 11
	TWITTER                   // 12
)

var Browsers = map[int64]string{
	UNKNOWN:        "Unknown",
	IE:             "IE",
	CHROME:         "Chrome",
	SAFARI:         "Safari",
	FIREFOX:        "Firefox",
	OPERA:          "Opera",
	CHROMIUM:       "Chromium",
	UCBROWSER:      "UC Browser",
	ANDROIDBROWSER: "Android Browser",
	BLACKBERRY:     "BlackBerry Browser",
	FACEBOOK:       "Facebook Browser",
	TWITTER:        "Twitter Browser",
}

var browsers = []pattern{

	pattern{
		TWITTER,
		[]string{"twitter for iphone"},
		[]string{},
		nil,
	},
	pattern{
		TWITTER,
		[]string{"twitterandroid"},
		[]string{},
		nil,
	},

	// BlackBerry
	pattern{
		BLACKBERRY,
		[]string{"blackberry"},
		[]string{"opera mini"},
		regexp.MustCompile(`blackberry (\d+)`),
	},

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

	// IE 12 (edge)
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
		[]string{
			"chromium",
			"chromeframe",
			"edge", // IE12 preview.
			"vivaldi",
			" opr/",
			"fbav/",
		},
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
		[]string{
			"chrome",
			"chromium",
			"crios",
			"opios", // Opera on iOS using Opera Turbo.
			"ucbrowser",
			"qqbrowser",
			"android",
			"fxios", // Firefox on iOS.
		},
		regexp.MustCompile(`version/(\d+)\.(\d+)`),
	},

	pattern{
		ANDROIDBROWSER,
		[]string{"android", "mobile safari"},
		[]string{"qqbrowser", "fbav/"},
		regexp.MustCompile(`version/(\d+)\.(\d+)`),
	},

	// Firefox
	pattern{
		FIREFOX,
		[]string{"firefox"},
		[]string{"seamonkey"},
		regexp.MustCompile(`firefox/(\d+)\.(\d+)`),
	},
	pattern{ // On iPhone.
		FIREFOX,
		[]string{"fxios"},
		[]string{},
		regexp.MustCompile(`fxios/(\d+)\.(\d+)`),
	},

	// Opera
	pattern{ // Mini
		OPERA,
		[]string{"opera mini"},
		[]string{},
		regexp.MustCompile(`opera/(\d+)\.(\d+)`),
	},
	pattern{
		OPERA,
		[]string{"opera"},
		[]string{},
		regexp.MustCompile(`version/(\d+)\.(\d+)`),
	},
	pattern{ // Turbo on iOS
		OPERA,
		[]string{"opios"},
		[]string{},
		regexp.MustCompile(`opios/(\d+)\.(\d+)`),
	},
	pattern{ // Opera Next
		OPERA,
		[]string{" opr/"},
		[]string{},
		regexp.MustCompile(` opr/(\d+)\.(\d+)`),
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

	pattern{
		FACEBOOK,
		[]string{"fbav/"},
		[]string{},
		regexp.MustCompile(`fbav/(\d+)\.(\d+)`),
	},
}

// Browser will return the browser id and version number.
// Returns 1,0 when nothing matched.
func Browser(userAgent string) (int64, int, int) {
	return find(browsers, userAgent)
}
