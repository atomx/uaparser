
package uaparser


var deviceTypes = []pattern{

  // PC
  pattern{
    2,
    []string{"windows"},
    []string{"windows phone"},
    nil,
  },
  pattern{
    2,
    []string{"mac os x"},
    []string{"iphone", "ipad", "ipod"},
    nil,
  },
  pattern{
    2,
    []string{"linux"},
    []string{"android"},
    nil,
  },

  // Phone
  pattern{
    3,
    []string{"iphone"},
    []string{},
    nil,
  },
  pattern{
    3,
    []string{"windows phone"},
    []string{},
    nil,
  },
  pattern{
    3,
    []string{"android", "mobile"},
    []string{},
    nil,
  },

  // Tablet
  pattern{
    4,
    []string{"ipad"},
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

