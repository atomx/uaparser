
package uaparser


import (
    "testing"
)


type test struct {
  userAgent string

  deviceTypeId Id

  deviceId      Id
  deviceVersion float32

  operatingSystemId      Id
  operatingSystemVersion float32

  browserId      Id
  browserVersion float32
}


var (
    tests = []test{
        test{
            "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
            2, // PC
            1, // Unknown device
            0,
            2, // Windows
            6.1,
            2, // IE
            10.0,
        },
        test{
            "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36",
            2, // PC
            1, // Unknown device
            0,
            2, // Windows
            6.1,
            3, // Chrome
            36.0,
        },
        test{
            "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53",
            3, // Phone
            2, // Iphone
            0,
            3, // iOS
            7.1,
            4, // Safari
            7.0,
        },
        test{
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.77.4 (KHTML, like Gecko) Version/7.0.5 Safari/537.77.4",
            2, // PC
            1, // Unknown device
            0,
            4, // Mac OS X
            10.9,
            4, // Safari
            7.0,
        },
        test{
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36",
            2, // Dekstop
            1, // Unknown device
            0,
            4, // Mac OS X
            10.9,
            3, // Chrome
            36.0,
        },
        test{
            "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0",
            2, // PC
            1, // Unknown device
            0,
            2, // Windows
            6.1,
            5, // Firefox
            30.0,
        },
        test{
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:30.0) Gecko/20100101 Firefox/30.0",
            2, // PC
            1, // Unknown device
            0,
            4, // Mac OS X
            10.9,
            5, // Firefox
            30.0,
        },
        test{
            "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:31.0) Gecko/20100101 Firefox/31.0",
            2, // PC
            1, // Unknown device
            0,
            2, // Windows
            6.1,
            5, // Firefox
            31.0,
        },
        test{
            "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
            2, // PC
            1, // Unknown device
            0,
            2, // Windows
            6.1,
            2, // IE
            11.0,
        },
        test{
            "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36",
            2, // PC
            1, // Unknown device
            0,
            5, // Linux
            0,
            3, // Chrome
            35.0,
        },
        test{
            "Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19",
            3, // Phone
            5, // Galaxy Nexus
            0,
            6, // Android
            4.0,
            3, // Chrome
            18.0,
        },
        test{
            "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
            2, // PC
            1, // Unknown device
            0,
            2, // Windows
            6.1,
            2, // IE
            9.0,
        },
        test{
            "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.11 (KHTML like Gecko) Chrome/23.0.1271.95 Safari/537.11",
            2, // PC
            1, // Unknown device
            0,
            2, // Windows
            5.1,
            3, // Chrome
            23.0,
        },
        test{
            "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
            2, // PC
            1, // Unknown device
            0,
            2, // Windows
            6.1,
            2, // IE
            7.0,
        },
        test{
            "Opera/9.80 (Windows NT 6.2; Win64; x64) Presto/2.12 Version/12.16",
            2, // PC
            1, // Unknown device
            0,
            2, // Windows
            6.2,
            6, // Opera
            9.8,
        },
        test{
            "Mozilla/5.0 (iPad; CPU OS 613 like Mac OS X) AppleWebKit/536.26 (KHTML like Gecko) Version/6.0 Mobile/10B329 Safari/8536.25",
            4, // Tablet
            3, // iPad
            0,
            3, // iOS
            613,
            4, // Safari
            6.0,
        },
        test{
            "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)",
            2, // PC
            1, // Unknown device
            0,
            2, // Windows
            5.1,
            2, // IE
            7.0,
        },
        test{
            "Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 520)",
            3, // Phone
            6, // Nokia Lumia
            520,
            7, // Windows Phone
            8.0,
            2, // IE
            10.0,
        },
    }
)


func TestUserAgents(t *testing.T) {
    for _, test := range tests {
      deviceTypeId                              := DeviceType     (test.userAgent)
      operatingSystemId, operatingSystemVersion := OperatingSystem(test.userAgent)
      browserId        , browserVersion         := Browser        (test.userAgent)

      t.Log(test.userAgent)

      if deviceTypeId != test.deviceTypeId {
        t.Errorf(" - device type id %d does not match expected %d", deviceTypeId, test.deviceTypeId)
      }

      if operatingSystemId != test.operatingSystemId {
        t.Errorf(" - operating system id %d does not match expected %d", operatingSystemId, test.operatingSystemId)
      }
      if operatingSystemVersion != test.operatingSystemVersion {
        t.Errorf(" - operating system version %f does not match expected %f", operatingSystemVersion, test.operatingSystemVersion)
      }

      if browserId != test.browserId {
        t.Errorf(" - browser id %d does not match expected %d", browserId, test.browserId)
      }
      if browserVersion != test.browserVersion {
        t.Errorf(" - browser version %f does not match expected %f", browserVersion, test.browserVersion)
      }
    }
}


func BenchmarkUserAgents(b *testing.B) {
  for i := 0; i < b.N; i++ {
    for _, test := range tests {
      _    = DeviceType     (test.userAgent)
      _, _ = OperatingSystem(test.userAgent)
      _, _ = Browser        (test.userAgent)
    }
  }
}

