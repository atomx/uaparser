
package uaparser


const (
  PC Id = iota + 2 // 2
  PHONE            // 3
  TABLET           // 4
  CONSOLE          // 5
)


var deviceTypes = []pattern{

  // PC
  pattern{
    PC,
    []string{"windows"},
    []string{"windows phone"},
    nil,
  },
  pattern{
    PC,
    []string{"mac os x"},
    []string{"iphone", "ipad", "ipod"},
    nil,
  },
  pattern{
    PC,
    []string{"linux"},
    []string{"android"},
    nil,
  },
  pattern{
    PC,
    []string{"cros"},
    []string{},
    nil,
  },

  // Phone
  pattern{
    PHONE,
    []string{"iphone"},
    []string{},
    nil,
  },
  pattern{
    PHONE,
    []string{"windows phone"},
    []string{},
    nil,
  },
  pattern{
    PHONE,
    []string{"android", "mobile"},
    []string{},
    nil,
  },
  pattern{
    PHONE,
    []string{"bb"},
    []string{},
    nil,
  },

  // Tablet
  pattern{
    TABLET,
    []string{"ipad"},
    []string{},
    nil,
  },
  pattern{
    TABLET,
    []string{"android"},
    []string{"mobile", "cros"},
    nil,
  },
  pattern{
    TABLET,
    []string{"rim Tablet"},
    []string{},
    nil,
  },

  // Console
  pattern{
    CONSOLE,
    []string{"playstation"},
    []string{},
    nil,
  },

}


// Will return the device type.
// Returns 1 when nothing matched.
func DeviceType(userAgent string) (Id) {
  id, _ := find(deviceTypes, userAgent)
  return id
}

