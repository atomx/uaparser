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

var OperatingSystems = map[int64]string{
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
	{
		RIMOS,
		[]string{"blackberry"},
		[]string{},
		nil,
	},

	// Windows
	{
		WINDOWS,
		[]string{"windows"},
		[]string{"windows phone"},
		regexp.MustCompile(`windows nt[ ]?(\d+)\.(\d+)`),
	},

	// iOS
	{
		IOS,
		[]string{"like mac os x"},
		[]string{},
		regexp.MustCompile(`os (\d+)(?:_(\d+))?`),
	},

	// Mac OS X
	{
		MACOSX,
		[]string{"mac os x"},
		[]string{"iphone", "ipad", "ipod"},
		regexp.MustCompile(`mac os x (\d+)[_\.](\d+)`), // Safari uses '_', Firefox uses '.'.
	},

	// Linux
	{
		LINUX,
		[]string{"linux"},
		[]string{"android"},
		nil,
	},

	// Android
	{
		ANDROID,
		[]string{"android"},
		[]string{"windows phone"}, // IE 12 on windows phone says it's android :(
		regexp.MustCompile(`android (\d+)\.(\d+)`),
	},

	// Windows Phone
	{
		WINDOWSPHONE,
		[]string{"windows phone"},
		[]string{},
		regexp.MustCompile(`phone (?:os )?(\d+)\.(\d+)`),
	},

	// Chrome OS
	{
		CHROMEOS,
		[]string{"cros"},
		[]string{},
		nil,
	},
}

// OperatingSystem will return the operating system id and version number.
// Returns 1,0 when nothing matched.
func OperatingSystem(userAgent string) (int64, int, int) {
	return find(operatingSystems, userAgent)
}
