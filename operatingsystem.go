package uaparser

import (
	"regexp"
)

const (
	WINDOWS       Id = iota + 2 // 2
	IOS                         // 3
	MAC_OS_X                    // 4
	LINUX                       // 5
	ANDROID                     // 6
	WINDOWS_PHONE               // 7
	CHROME_OS                   // 8
)

var operatingSystems = []pattern{

	// Windows
	pattern{
		WINDOWS,
		[]string{"windows"},
		[]string{"windows phone"},
		regexp.MustCompile(`windows nt[ ]?(\d+\.\d+)`),
	},

	// iOS
	pattern{
		IOS,
		[]string{"like mac os x"},
		[]string{},
		regexp.MustCompile(`os (\d+)(?:_(\d+))?`),
	},

	// Mac OS X
	pattern{
		MAC_OS_X,
		[]string{"mac os x"},
		[]string{"iphone", "ipad", "ipod"},
		regexp.MustCompile(`mac os x (\d+)[_\.](\d+)`), // Safari uses '_', Firefox uses '.'.
	},

	// Linux
	pattern{
		LINUX,
		[]string{"linux"},
		[]string{"android"},
		nil,
	},

	// Android
	pattern{
		ANDROID,
		[]string{"android"},
		[]string{},
		regexp.MustCompile(`android (\d+\.\d+)`),
	},

	// Windows Phone
	pattern{
		WINDOWS_PHONE,
		[]string{"windows phone"},
		[]string{},
		regexp.MustCompile(`phone (?:os )?(\d+\.\d+)`),
	},

	// Chrome OS
	pattern{
		CHROME_OS,
		[]string{"cros"},
		[]string{},
		nil,
	},
}

// Will return the operating system id and version number.
// Returns 1,0 when nothing matched.
func OperatingSystem(userAgent string) (Id, float32) {
	return find(operatingSystems, userAgent)
}
