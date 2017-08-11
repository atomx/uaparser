package uaparser

import (
	"testing"
)

type test struct {
	userAgent string

	deviceTypeId int64

	operatingSystemId           int64
	operatingSystemVersionMajor int
	operatingSystemVersionMinor int

	browserId           int64
	browserVersionMajor int
	browserVersionMinor int
}

var (
	tests = []test{
		{
			"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
			PC,
			WINDOWS, 6, 1,
			IE, 10, 0,
		},
		{
			"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36",
			PC,
			WINDOWS, 6, 1,
			CHROME, 36, 0,
		},
		{
			"Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53",
			PHONE,
			IOS, 7, 1,
			SAFARI, 7, 0,
		},
		{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.77.4 (KHTML, like Gecko) Version/7.0.5 Safari/537.77.4",
			PC,
			MACOSX, 10, 9,
			SAFARI, 7, 0,
		},
		{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36",
			PC,
			MACOSX, 10, 9,
			CHROME, 36, 0,
		},
		{
			"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0",
			PC,
			WINDOWS, 6, 1,
			FIREFOX, 30, 0,
		},
		{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:30.0) Gecko/20100101 Firefox/30.0",
			PC,
			MACOSX, 10, 9,
			FIREFOX, 30, 0,
		},
		{
			"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:31.0) Gecko/20100101 Firefox/31.0",
			PC,
			WINDOWS, 6, 1,
			FIREFOX, 31, 0,
		},
		{
			"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
			PC,
			WINDOWS, 6, 1,
			IE, 11, 0,
		},
		{
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36",
			PC,
			LINUX, 0, 0,
			CHROME, 35, 0,
		},
		{
			"Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19",
			PHONE,
			ANDROID, 4, 0,
			CHROME, 18, 0,
		},
		{
			"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
			PC,
			WINDOWS, 6, 1,
			IE, 9, 0,
		},
		{
			"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.11 (KHTML like Gecko) Chrome/23.0.1271.95 Safari/537.11",
			PC,
			WINDOWS, 5, 1,
			CHROME, 23, 0,
		},
		{
			"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
			SETTOPBOX,
			WINDOWS, 6, 1,
			IE, 7, 0,
		},
		{
			"Opera/9.80 (Windows NT 6.2; Win64; x64) Presto/2.12 Version/12.16",
			PC,
			WINDOWS, 6, 2,
			OPERA, 12, 16,
		},
		{
			"Mozilla/5.0 (iPad; CPU OS 613 like Mac OS X) AppleWebKit/536.26 (KHTML like Gecko) Version/6.0 Mobile/10B329 Safari/8536.25",
			TABLET,
			IOS, 613, 0,
			SAFARI, 6, 0,
		},
		{
			"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)",
			PC,
			WINDOWS, 5, 1,
			IE, 7, 0,
		},
		{
			"Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 520)",
			PHONE,
			WINDOWSPHONE, 8, 0,
			IE, 10, 0,
		},
		{
			"Mozilla/5.0 (Linux; Android 4.4.4; Nexus 7 Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.131 Safari/537.36",
			TABLET,
			ANDROID, 4, 4,
			CHROME, 36, 0,
		},
		{
			"Mozilla/5.0 (X11; CrOS x86_64 5841.83.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.138 Safari/537.36",
			PC,
			CHROMEOS, 0, 0,
			CHROME, 36, 0,
		},
		{
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/34.0.1847.116 Chrome/34.0.1847.116 Safari/537.36",
			PC,
			LINUX, 0, 0,
			CHROMIUM, 34, 0,
		},
		{
			"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; chromeframe/11.0.660.0)",
			PC,
			WINDOWS, 5, 1,
			UNKNOWN, 0, 0, // Don't detect a chromeframe as IE or Chrome.
		},
		{
			"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; chromeframe/11.0.660.0) AppleWebKit/534.18 (KHTML, like Gecko) Chrome/11.0.660.0 Safari/534.18",
			PC,
			WINDOWS, 5, 1,
			UNKNOWN, 0, 0, // Don't detect a chromeframe sub-requests as IE or Chrome.
		},
		{
			"Mozilla/5.0 (X11; U; Linux i686; en-US) AppleWebKit/533.4 (KHTML, like Gecko) Chrome/5.0.375.127 Large Screen Safari/533.4 GoogleTV/ 162671",
			SETTOPBOX,
			LINUX, 0, 0,
			CHROME, 5, 0,
		},
		{
			"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; OfficeLiveConnector.1.5; OfficeLivePatch.1.3; .NET4.0C; .NET4.0E)",
			SETTOPBOX,
			WINDOWS, 6, 0,
			IE, 7, 0,
		},
		{
			"Mozilla/5.0 (Windows NT 6.4; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Safari/537.36 Edge/12.0",
			PC,
			WINDOWS, 6, 4,
			EDGE, 12, 0,
		},
		{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.89 Vivaldi/1.0.83.38 Safari/537.36",
			PC,
			MACOSX, 10, 10,
			UNKNOWN, 0, 0,
		},
		{
			"Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; SM-T210R Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30 UCBrowser/2.3.2.300",
			TABLET,
			ANDROID, 4, 1,
			UCBROWSER, 2, 3,
		},
		{ // QQ Browser
			"Mozilla/5.0 (Linux; U; Android 4.4.4; zh-cn; HTC M8t Build/KTU84P) AppleWebKit/537.36 (KHTML like Gecko)Version/4.0 MQQBrowser/5.6 Mobile Safari/537.36",
			PHONE,
			ANDROID, 4, 4,
			UNKNOWN, 0, 0,
		},
		{ // Opera on iOS using Opera Turbo.
			"Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) OPiOS/10.0.1.90729 Mobile/12F70 Safari/9537.53",
			PHONE,
			IOS, 8, 3,
			OPERA, 10, 0,
		},
		{
			"Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; sdk Build/MR1) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
			PHONE,
			ANDROID, 4, 0,
			ANDROIDBROWSER, 4, 0,
		},
		{
			"Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; mmi_apps_sdk Build/6.7.1_22) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
			PHONE,
			ANDROID, 4, 0,
			ANDROIDBROWSER, 4, 0,
		},
		{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.82 Safari/537.36 OPR/29.0.1795.41 (Edition beta)",
			PC,
			MACOSX, 10, 10,
			OPERA, 29, 0,
		},
		{
			"Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; SM-T210R Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30 UCBrowser/2.3.2.300",
			TABLET,
			ANDROID, 4, 1,
			UCBROWSER, 2, 3,
		},
		{
			"Mozilla/5.0 (BlackBerry; U; BlackBerry 9900; en) AppleWebKit/534.11+ (KHTML, like Gecko) Version/7.1.0.346 Mobile Safari/534.11+",
			PHONE,
			RIMOS, 0, 0,
			BLACKBERRY, 9900, 0,
		},
		{
			// Test a really high (but realistic) minor version number.
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10136",
			PC,
			WINDOWS, 10, 0,
			EDGE, 12, 10136,
		},
		{
			"Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; en) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Mobile Safari/537.36 Edge/12.10316",
			PHONE,
			WINDOWSPHONE, 10, 0,
			EDGE, 12, 10316,
		},
		{
			"Opera/9.80 (BlackBerry; Opera Mini/8.0.35667/37.6897; U; en) Presto/2.12.423 Version/12.16",
			PHONE,
			RIMOS, 0, 0,
			OPERA, 9, 80,
		},
		{
			"Opera/9.80 (Series 60; Opera Mini/7.1.32453/37.6897; U; en) Presto/2.12.423 Version/12.16",
			PHONE,
			UNKNOWN, 0, 0,
			OPERA, 9, 80,
		},
		{
			"Raptr RaptrDesktopApp/4.5.0 RaptrBuild-AMD RaptrControlCenter",
			UNKNOWN,
			UNKNOWN, 0, 0,
			UNKNOWN, 0, 0,
		},
		{
			"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/43.0.2357.121 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/35.0.0.48.273;]",
			PHONE,
			ANDROID, 5, 0,
			FACEBOOK, 35, 0,
		},
		{
			"Mozilla/5.0 (iPad; CPU OS 9_2_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Mobile/13D15 [FBAN/FBIOS;FBAV/51.0.0.52.154;FBBV/25351091;FBDV/iPad2,4;FBMD/iPad;FBSN/iPhone OS;FBSV/9.2.1;FBSS/1; FBCR/;FBID/tablet;FBLC/en_US;FBOP/1]",
			TABLET,
			IOS, 9, 2,
			FACEBOOK, 51, 0,
		},
		{
			"Mozilla/5.0 (iPad; CPU OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12F69 [FBAN/FBIOS;FBAV/51.0.0.52.154;FBBV/25351091;FBDV/iPad2,5;FBMD/iPad;FBSN/iPhone OS;FBSV/8.3;FBSS/1; FBCR/;FBID/tablet;FBLC/en_US;FBOP/1]",
			TABLET,
			IOS, 8, 3,
			FACEBOOK, 51, 0,
		},
		{
			"Mozilla/5.0 (iPhone; CPU iPhone OS 5_1_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Mobile/9B206 Twitter for iPhone",
			PHONE,
			IOS, 5, 1,
			TWITTER, 0, 0,
		},
		{
			"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 10 Build/LMY48T; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/46.0.2490.76 Safari/537.36 TwitterAndroid",
			TABLET,
			ANDROID, 5, 1,
			TWITTER, 0, 0,
		},
		{
			"mozilla/5.0 (iphone; cpu iphone os 8_1_2 like mac os x) applewebkit/600.1.4 (khtml, like gecko) mobile/12b440 qq/5.3.0.319 nettype/wifi mem/205",
			PHONE,
			IOS, 8, 1,
			QQBROWSER, 5, 3,
		},
		{
			"Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mobile/10A403",
			TABLET,
			IOS, 6, 0,
			SAFARI, 0, 0,
		},
	}
)

func TestUserAgents(t *testing.T) {
	for _, test := range tests {
		t.Run(test.userAgent, func(t *testing.T) {
			deviceTypeId := DeviceType(test.userAgent)
			operatingSystemId, operatingSystemVersionMajor, operatingSystemVersionMinor := OperatingSystem(test.userAgent)
			browserId, browserVersionMajor, browserVersionMinor := Browser(test.userAgent)

			if deviceTypeId != test.deviceTypeId {
				t.Errorf(" - device type %s does not match expected %s", DeviceTypes[deviceTypeId], DeviceTypes[test.deviceTypeId])
			}

			if operatingSystemId != test.operatingSystemId {
				t.Errorf(" - operating system id %s does not match expected %s", OperatingSystems[operatingSystemId], OperatingSystems[test.operatingSystemId])
			}
			if operatingSystemVersionMajor != test.operatingSystemVersionMajor ||
				operatingSystemVersionMinor != test.operatingSystemVersionMinor {
				t.Errorf(" - operating system version %d.%d does not match expected %d.%d", operatingSystemVersionMajor, operatingSystemVersionMinor, test.operatingSystemVersionMajor, test.operatingSystemVersionMinor)
			}

			if browserId != test.browserId {
				t.Errorf(" - browser id %s does not match expected %s", Browsers[browserId], Browsers[test.browserId])
			}
			if browserVersionMajor != test.browserVersionMajor ||
				browserVersionMinor != test.browserVersionMinor {
				t.Errorf(" - browser version %d.%d does not match expected %d.%d", browserVersionMajor, browserVersionMinor, test.browserVersionMajor, test.browserVersionMinor)
			}
		})
	}
}

func BenchmarkUserAgents(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			_ = DeviceType(test.userAgent)
			_, _, _ = OperatingSystem(test.userAgent)
			_, _, _ = Browser(test.userAgent)
		}
	}
}
