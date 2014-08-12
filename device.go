
package uaparser


import (
    "regexp"
)


const (
  IPHONE Id = iota + 2 // 2
  IPAD                 // 3
  IPOD                 // 4
  GALAXY_NEXUS         // 5
  NOKIA_LUMIA          // 6
  NEXUS                // 7
)


var devices = []pattern{

  // iPhone
  pattern{
    IPHONE,
    []string{"iphone"},
    []string{},
    nil,
  },

  // iPad
  pattern{
    IPAD,
    []string{"ipad"},
    []string{},
    nil,
  },

  // iPod
  pattern{
    IPOD,
    []string{"ipod"},
    []string{},
    nil,
  },

  // Galaxy Nexus
  pattern{
    GALAXY_NEXUS,
    []string{"galaxy nexus"},
    []string{},
    nil,
  },

  // Nokia Lumia
  pattern{
    NOKIA_LUMIA,
    []string{"nokia", "lumia"},
    []string{},
    regexp.MustCompile(`lumia (\d+)`),
  },

  // Nexus
  pattern{
    NEXUS,
    []string{"nexus"},
    []string{},
    regexp.MustCompile(`nexus (\d+)`),
  },

}


// Will return the device id and version number.
// Returns 1,0 when nothing matched.
func Device(userAgent string) (Id, float32) {
    return find(devices, userAgent)
}

