package uaparser

import (
	"regexp"
)

// Operating System id's.
const (
	WINDOWS      = iota + 2 // 2
	IOS                     // 3
	MACOSX                  // 4
	LINUX                   // 5
	ANDROID                 // 6
	WINDOWSPHONE            // 7
	CHROMEOS                // 8
	RIMOS                   // 9
)

var OperatingSystems = map[uint]string{
	WINDOWS:      "Windows",
	IOS:          "iOS",
	MACOSX:       "Mac OS X",
	LINUX:        "Linux",
	ANDROID:      "Android",
	WINDOWSPHONE: "Windows Phone",
	CHROMEOS:     "Chrome OS",
	RIMOS:        "RIM OS",
}

var operatingSystems = []pattern{

	// RIM OS
	pattern{
		RIMOS,
		[]string{"blackberry"},
		[]string{},
		nil,
	},

	// Windows
	pattern{
		WINDOWS,
		[]string{"windows"},
		[]string{"windows phone"},
		regexp.MustCompile(`windows nt[ ]?(\d+)\.(\d+)`),
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
		MACOSX,
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
		regexp.MustCompile(`android (\d+)\.(\d+)`),
	},

	// Windows Phone
	pattern{
		WINDOWSPHONE,
		[]string{"windows phone"},
		[]string{},
		regexp.MustCompile(`phone (?:os )?(\d+)\.(\d+)`),
	},

	// Chrome OS
	pattern{
		CHROMEOS,
		[]string{"cros"},
		[]string{},
		nil,
	},
}

// OperatingSystem will return the operating system id and version number.
// Returns 1,0 when nothing matched.
func OperatingSystem(userAgent string) (uint, int) {
	return find(operatingSystems, userAgent)
}
