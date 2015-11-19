package uaparser

// Device Type id's.
const (
	PC        = iota + 2 // 2
	PHONE                // 3
	TABLET               // 4
	CONSOLE              // 5
	SETTOPBOX            // 6
)

var DeviceTypes = map[uint]string{
	UNKNOWN:   "Unknown",
	PC:        "PC",
	PHONE:     "Phone",
	TABLET:    "Tablet",
	CONSOLE:   "Console",
	SETTOPBOX: "Set top box",
}

var deviceTypes = []pattern{
	// Set-top box
	pattern{
		SETTOPBOX,
		[]string{"media center pc"},
		[]string{},
		nil,
	},
	pattern{
		SETTOPBOX,
		[]string{"googletv"},
		[]string{},
		nil,
	},

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
		[]string{"ipad"}, // Firefox for iPad contains both.
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
	pattern{
		PHONE,
		[]string{"blackberry"},
		[]string{"playbook"},
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
	pattern{
		TABLET,
		[]string{"tablet pc"},
		[]string{},
		nil,
	},
	pattern{
		TABLET,
		[]string{"blackberry", "playbook"},
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
	pattern{
		CONSOLE,
		[]string{"xbox"},
		[]string{},
		nil,
	},
	pattern{
		CONSOLE,
		[]string{"nintendo"},
		[]string{},
		nil,
	},

	// When we don't know anything else we just mark all Opera Mini devices as phones.
	pattern{
		PHONE,
		[]string{"opera mini"},
		[]string{},
		nil,
	},
}

// DeviceType will return the device type.
// Returns 1 when nothing matched.
func DeviceType(userAgent string) uint {
	id, _, _ := find(deviceTypes, userAgent)
	return id
}
