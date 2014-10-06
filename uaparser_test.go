package uaparser

import (
	"testing"
)

type test struct {
	userAgent string

	deviceTypeID ID

	operatingSystemID      ID
	operatingSystemVersion int

	browserID      ID
	browserVersion int
}

var (
	tests = []test{
		test{
			"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
			PC,
			WINDOWS, Version(6, 1),
			IE, Version(10, 0),
		},
		test{
			"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36",
			PC,
			WINDOWS, Version(6, 1),
			CHROME, Version(36, 0),
		},
		test{
			"Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53",
			PHONE,
			IOS, Version(7, 1),
			SAFARI, Version(7, 0),
		},
		test{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.77.4 (KHTML, like Gecko) Version/7.0.5 Safari/537.77.4",
			PC,
			MACOSX, Version(10, 9),
			SAFARI, Version(7, 0),
		},
		test{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36",
			PC,
			MACOSX, Version(10, 9),
			CHROME, Version(36, 0),
		},
		test{
			"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0",
			PC,
			WINDOWS, Version(6, 1),
			FIREFOX, Version(30, 0),
		},
		test{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:30.0) Gecko/20100101 Firefox/30.0",
			PC,
			MACOSX, Version(10, 9),
			FIREFOX, Version(30, 0),
		},
		test{
			"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:31.0) Gecko/20100101 Firefox/31.0",
			PC,
			WINDOWS, Version(6, 1),
			FIREFOX, Version(31, 0),
		},
		test{
			"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
			PC,
			WINDOWS, Version(6, 1),
			IE, Version(11, 0),
		},
		test{
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36",
			PC,
			LINUX, Version(0, 0),
			CHROME, Version(35, 0),
		},
		test{
			"Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19",
			PHONE,
			ANDROID, Version(4, 0),
			CHROME, Version(18, 0),
		},
		test{
			"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
			PC,
			WINDOWS, Version(6, 1),
			IE, Version(9, 0),
		},
		test{
			"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.11 (KHTML like Gecko) Chrome/23.0.1271.95 Safari/537.11",
			PC,
			WINDOWS, Version(5, 1),
			CHROME, Version(23, 0),
		},
		test{
			"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
			SETTOPBOX,
			WINDOWS, Version(6, 1),
			IE, Version(7, 0),
		},
		test{
			"Opera/9.80 (Windows NT 6.2; Win64; x64) Presto/2.12 Version/12.16",
			PC,
			WINDOWS, Version(6, 2),
			OPERA, Version(9, 80),
		},
		test{
			"Mozilla/5.0 (iPad; CPU OS 613 like Mac OS X) AppleWebKit/536.26 (KHTML like Gecko) Version/6.0 Mobile/10B329 Safari/8536.25",
			TABLET,
			IOS, Version(613, 0),
			SAFARI, Version(6, 0),
		},
		test{
			"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)",
			PC,
			WINDOWS, Version(5, 1),
			IE, Version(7, 0),
		},
		test{
			"Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 520)",
			PHONE,
			WINDOWSPHONE, Version(8, 0),
			IE, Version(10, 0),
		},
		test{
			"Mozilla/5.0 (Linux; Android 4.4.4; Nexus 7 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.131 Safari/537.36",
			TABLET,
			ANDROID, Version(4, 4),
			CHROME, Version(36, 0),
		},
		test{
			"Mozilla/5.0 (X11; CrOS x86_64 5841.83.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.138 Safari/537.36",
			PC,
			CHROMEOS, Version(0, 0),
			CHROME, Version(36, 0),
		},
		test{
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/34.0.1847.116 Chrome/34.0.1847.116 Safari/537.36",
			PC,
			LINUX, Version(0, 0),
			CHROMIUM, Version(34, 0),
		},
		test{
			"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; chromeframe/11.0.660.0)",
			PC,
			WINDOWS, Version(5, 1),
			UNKNOWN, Version(0, 0), // Don't detect a chromeframe as IE or Chrome.
		},
		test{
			"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; chromeframe/11.0.660.0) AppleWebKit/534.18 (KHTML, like Gecko) Chrome/11.0.660.0 Safari/534.18",
			PC,
			WINDOWS, Version(5, 1),
			UNKNOWN, Version(0, 0), // Don't detect a chromeframe sub-requests as IE or Chrome.
		},
		test{
			"Mozilla/5.0 (X11; U; Linux i686; en-US) AppleWebKit/533.4 (KHTML, like Gecko) Chrome/5.0.375.127 Large Screen Safari/533.4 GoogleTV/ 162671",
			SETTOPBOX,
			LINUX, Version(0, 0),
			CHROME, Version(5, 0),
		},
		test{
			"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; OfficeLiveConnector.1.5; OfficeLivePatch.1.3; .NET4.0C; .NET4.0E)",
			SETTOPBOX,
			WINDOWS, Version(6, 0),
			IE, Version(7, 0),
		},
	}
)

func TestUserAgents(t *testing.T) {
	for _, test := range tests {
		deviceTypeID := DeviceType(test.userAgent)
		operatingSystemID, operatingSystemVersion := OperatingSystem(test.userAgent)
		browserID, browserVersion := Browser(test.userAgent)

		t.Log(test.userAgent)

		if deviceTypeID != test.deviceTypeID {
			t.Errorf(" - device type id %d does not match expected %d", deviceTypeID, test.deviceTypeID)
		}

		if operatingSystemID != test.operatingSystemID {
			t.Errorf(" - operating system id %d does not match expected %d", operatingSystemID, test.operatingSystemID)
		}
		if operatingSystemVersion != test.operatingSystemVersion {
			t.Errorf(" - operating system version %d does not match expected %d", operatingSystemVersion, test.operatingSystemVersion)
		}

		if browserID != test.browserID {
			t.Errorf(" - browser id %d does not match expected %d", browserID, test.browserID)
		}
		if browserVersion != test.browserVersion {
			t.Errorf(" - browser version %d does not match expected %d", browserVersion, test.browserVersion)
		}
	}
}

func BenchmarkUserAgents(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			_ = DeviceType(test.userAgent)
			_, _ = OperatingSystem(test.userAgent)
			_, _ = Browser(test.userAgent)
		}
	}
}
