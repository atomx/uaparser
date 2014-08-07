
package uaparser


import (
    "regexp"
)


var devices = []pattern{

  // iPhone
  pattern{
    2,
    []string{"iphone"},
    []string{},
    nil,
  },

  // iPad
  pattern{
    3,
    []string{"ipad"},
    []string{},
    nil,
  },

  // iPod
  pattern{
    4,
    []string{"ipod"},
    []string{},
    nil,
  },

  // Galaxy Nexus
  pattern{
    5,
    []string{"galaxy nexus"},
    []string{},
    nil,
  },

  // Nokia Lumia
  pattern{
    6,
    []string{"nokia", "lumia"},
    []string{},
    regexp.MustCompile(`lumia (\d+)`),
  },

}


// Will return the device id and version number.
// Returns 1,0 when nothing matched.
func Device(userAgent string) (Id, float32) {
    return find(devices, userAgent)
}

