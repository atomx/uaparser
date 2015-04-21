package uaparser

import (
	"testing"
)

type test struct {
	userAgent string

	deviceTypeId int

	operatingSystemId      int
	operatingSystemVersion int

	browserId      int
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
		test{
			"Mozilla/5.0 (Windows NT 6.4; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Safari/537.36 Edge/12.0",
			PC,
			WINDOWS, Version(6, 4),
			IE, Version(12, 0),
		},
		test{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.89 Vivaldi/1.0.83.38 Safari/537.36",
			PC,
			MACOSX, Version(10, 10),
			UNKNOWN, Version(0, 0),
		},
		test{
			"Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; SM-T210R Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30 UCBrowser/2.3.2.300",
			TABLET,
			ANDROID, Version(4, 1),
			UCBROWSER, Version(2, 3),
		},
		test{ // QQ Browser
			"Mozilla/5.0 (Linux; U; Android 4.4.4; zh-cn; HTC M8t Build/KTU84P) AppleWebKit/537.36 (KHTML like Gecko)Version/4.0 MQQBrowser/5.6 Mobile Safari/537.36",
			PHONE,
			ANDROID, Version(4, 4),
			UNKNOWN, Version(0, 0),
		},
		test{ // Opera on iOS using Opera Turbo.
			"Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) OPiOS/10.0.1.90729 Mobile/12F70 Safari/9537.53",
			PHONE,
			IOS, Version(8, 3),
			OPERA, Version(10, 0),
		},
	}
)

func TestUserAgents(t *testing.T) {
	for _, test := range tests {
		deviceTypeId := DeviceType(test.userAgent)
		operatingSystemId, operatingSystemVersion := OperatingSystem(test.userAgent)
		browserId, browserVersion := Browser(test.userAgent)

		t.Log(test.userAgent)

		if deviceTypeId != test.deviceTypeId {
			t.Errorf(" - device type %s does not match expected %s", DeviceTypes[deviceTypeId], DeviceTypes[test.deviceTypeId])
		}

		if operatingSystemId != test.operatingSystemId {
			t.Errorf(" - operating system id %s does not match expected %s", OperatingSystems[operatingSystemId], OperatingSystems[test.operatingSystemId])
		}
		if operatingSystemVersion != test.operatingSystemVersion {
			t.Errorf(" - operating system version %d does not match expected %d", operatingSystemVersion, test.operatingSystemVersion)
		}

		if browserId != test.browserId {
			t.Errorf(" - browser id %s does not match expected %s", Browsers[browserId], Browsers[test.browserId])
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
