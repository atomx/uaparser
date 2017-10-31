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
	EDGE                      // 13
	QQBROWSER                 // 14
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
	EDGE:           "Edge",
	QQBROWSER:      "QQ Browser",
}

var browsers = []pattern{

	{
		TWITTER,
		[]string{"twitter for iphone"},
		[]string{},
		nil,
	},
	{
		TWITTER,
		[]string{"twitterandroid"},
		[]string{},
		nil,
	},

	// BlackBerry
	{
		BLACKBERRY,
		[]string{"blackberry"},
		[]string{"opera mini"},
		regexp.MustCompile(` blackberry (\d+)`),
	},

	// Edge
	{
		EDGE,
		[]string{" edge/"},
		[]string{},
		regexp.MustCompile(` edge/(\d+)\.(\d+)`),
	},

	// IE < 11
	{
		IE,
		[]string{"msie"},
		[]string{"chromeframe"},
		regexp.MustCompile(` msie (\d+)\.(\d+)`),
	},

	// IE 11
	{
		IE,
		[]string{"trident"},
		[]string{"chromeframe"},
		regexp.MustCompile(` rv:(\d+)\.(\d+)`),
	},

	// IE 12 (edge)
	// Looks like:
	// Mozilla/5.0 (Windows NT 6.4; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Safari/537.36 Edge/12.0
	{
		IE,
		[]string{"applewebkit", "chrome", "safari", "edge"},
		[]string{},
		regexp.MustCompile(` edge/(\d+)\.(\d+)`),
	},

	// Chrome
	{
		CHROME,
		[]string{" chrome/"},
		[]string{
			"chromium",
			"chromeframe",
			"edge", // IE12 preview.
			"vivaldi",
			" opr/",
			"fbav/",
		},
		regexp.MustCompile(` chrome/(\d+)\.(\d+)`),
	},
	// Chrome on iOS
	{
		CHROME,
		[]string{" crios/"},
		[]string{},
		regexp.MustCompile(` crios/(\d+)\.(\d+)`),
	},

	// Safari
	{
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
			"samsungbrowser",
			"valve",
		},
		regexp.MustCompile(` version/(\d+)\.(\d+)`),
	},
	{
		SAFARI,
		[]string{"ipad", "applewebkit", "mobile"},
		[]string{
			"chrome",
			"chromium",
			"crios",
			"opios", // Opera on iOS using Opera Turbo.
			"ucbrowser",
			"qqbrowser",
			"android",
			"fxios", // Firefox on iOS.
			"fbav/",
			" qq/",
		},
		nil,
	},
	{
		SAFARI,
		[]string{"iphone", "applewebkit", "mobile"},
		[]string{
			"chrome",
			"chromium",
			"crios",
			"opios", // Opera on iOS using Opera Turbo.
			"ucbrowser",
			"qqbrowser",
			"android",
			"fxios", // Firefox on iOS.
			"fbav/",
			" qq/",
		},
		nil,
	},

	// Android
	{
		ANDROIDBROWSER,
		[]string{"android", "mobile safari"},
		[]string{"qqbrowser", "fbav/"},
		regexp.MustCompile(` version/(\d+)\.(\d+)`),
	},

	// Firefox
	{
		FIREFOX,
		[]string{"firefox"},
		[]string{"seamonkey"},
		regexp.MustCompile(` firefox/(\d+)\.(\d+)`),
	},
	{ // On iPhone.
		FIREFOX,
		[]string{"fxios"},
		[]string{},
		regexp.MustCompile(` fxios/(\d+)\.(\d+)`),
	},

	// Opera
	{ // Mini
		OPERA,
		[]string{"opera mini"},
		[]string{},
		regexp.MustCompile(`opera/(\d+)\.(\d+)`),
	},
	{
		OPERA,
		[]string{"opera"},
		[]string{},
		regexp.MustCompile(` version/(\d+)\.(\d+)`),
	},
	{ // Turbo on iOS
		OPERA,
		[]string{"opios"},
		[]string{},
		regexp.MustCompile(` opios/(\d+)\.(\d+)`),
	},
	{ // Opera Next
		OPERA,
		[]string{" opr/"},
		[]string{},
		regexp.MustCompile(` opr/(\d+)\.(\d+)`),
	},

	// Chromium
	{
		CHROMIUM,
		[]string{"chromium"},
		[]string{},
		regexp.MustCompile(` chromium/(\d+)\.(\d+)`),
	},

	// UC Browser
	// See: https://play.google.com/store/apps/details?id=com.UCMobile.intl&hl=en
	// According to http://gs.statcounter.com/#mobile_browser-ww-monthly-201402-201502-bar it's the 5th biggest browser
	// with about the same number of users as Opera.
	{
		UCBROWSER,
		[]string{"ucbrowser"},
		[]string{},
		regexp.MustCompile(` ucbrowser/(\d+)\.(\d+)`),
	},

	{
		FACEBOOK,
		[]string{"fbav/"},
		[]string{},
		regexp.MustCompile(`fbav/(\d+)\.(\d+)`),
	},

	// QQ Browser
	{
		QQBROWSER,
		[]string{" qq/"},
		[]string{},
		regexp.MustCompile(` qq/(\d+)\.(\d+)`),
	},
}

// Browser will return the browser id and version number.
// Returns 1,0 when nothing matched.
func Browser(userAgent string) (int64, int, int) {
	return find(browsers, userAgent)
}
