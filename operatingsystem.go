package uaparser

import (
	"regexp"
)

// Operating System ID's.
const (
	WINDOWS      ID = iota + 2 // 2
	IOS                        // 3
	MACOSX                     // 4
	LINUX                      // 5
	ANDROID                    // 6
	WINDOWSPHONE               // 7
	CHROMEOS                   // 8
)

var OperatingSystems = map[ID]string{
	WINDOWS:      "Windows",
	IOS:          "iOS",
	MACOSX:       "Max OS X",
	LINUX:        "Linux",
	ANDROID:      "Android",
	WINDOWSPHONE: "Windows Phone",
	CHROMEOS:     "Chrome OS",
}

var operatingSystems = []pattern{

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
func OperatingSystem(userAgent string) (ID, int) {
	return find(operatingSystems, userAgent)
}
