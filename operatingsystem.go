
package uaparser


import (
    "regexp"
)


var operatingSystems = []pattern{

  // Windows 
  pattern{
    2,
    []string{"windows"},
    []string{"windows phone"},
    regexp.MustCompile(`windows nt[ ]?(\d+\.\d+)`),
  },

  // iOS
  pattern{
    3,
    []string{"like mac os x"},
    []string{},
    regexp.MustCompile(`os (\d+)(?:_(\d+))?`),
  },

  // Mac OS X
  pattern{
    4,
    []string{"mac os x"},
    []string{"iphone", "ipad", "ipod"},
    regexp.MustCompile(`mac os x (\d+)[_\.](\d+)`), // Safari uses '_', Firefox uses '.'.
  },

  // Linux
  pattern{
    5,
    []string{"linux"},
    []string{"android"},
    nil,
  },

  // Android
  pattern{
    6,
    []string{"android"},
    []string{},
    regexp.MustCompile(`android (\d+\.\d+)`),
  },

  // Windows Phone
  pattern{
    7,
    []string{"windows phone"},
    []string{},
    regexp.MustCompile(`phone (\d+\.\d+)`),
  },

}


// Will return the operating system id and version number.
// Returns 1,0 when nothing matched.
func OperatingSystem(userAgent string) (Id, float32) {
    return find(operatingSystems, userAgent)
}

