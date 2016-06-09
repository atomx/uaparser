package uaparser

// Device Type id's.
const (
	PC        = iota + 2 // 2
	PHONE                // 3
	TABLET               // 4
	CONSOLE              // 5
	SETTOPBOX            // 6
)

var DeviceTypes = map[int64]string{
	UNKNOWN:   "Unknown",
	PC:        "PC",
	PHONE:     "Phone",
	TABLET:    "Tablet",
	CONSOLE:   "Console",
	SETTOPBOX: "Set top box",
}

var deviceTypes = []pattern{
	// Set-top box
	{
		SETTOPBOX,
		[]string{"media center pc"},
		[]string{},
		nil,
	},
	{
		SETTOPBOX,
		[]string{"googletv"},
		[]string{},
		nil,
	},

	// PC
	{
		PC,
		[]string{"windows"},
		[]string{"windows phone"},
		nil,
	},
	{
		PC,
		[]string{"mac os x"},
		[]string{"iphone", "ipad", "ipod"},
		nil,
	},
	{
		PC,
		[]string{"linux"},
		[]string{"android"},
		nil,
	},
	{
		PC,
		[]string{"cros"},
		[]string{},
		nil,
	},

	// Phone
	{
		PHONE,
		[]string{"iphone"},
		[]string{"ipad"}, // Firefox for iPad contains both.
		nil,
	},
	{
		PHONE,
		[]string{"windows phone"},
		[]string{},
		nil,
	},
	{
		PHONE,
		[]string{"android", "mobile"},
		[]string{},
		nil,
	},
	{
		PHONE,
		[]string{"blackberry"},
		[]string{"playbook"},
		nil,
	},

	// Tablet
	{
		TABLET,
		[]string{"ipad"},
		[]string{},
		nil,
	},
	{
		TABLET,
		[]string{"android"},
		[]string{"mobile", "cros"},
		nil,
	},
	{
		TABLET,
		[]string{"rim Tablet"},
		[]string{},
		nil,
	},
	{
		TABLET,
		[]string{"tablet pc"},
		[]string{},
		nil,
	},
	{
		TABLET,
		[]string{"blackberry", "playbook"},
		[]string{},
		nil,
	},

	// Console
	{
		CONSOLE,
		[]string{"playstation"},
		[]string{},
		nil,
	},
	{
		CONSOLE,
		[]string{"xbox"},
		[]string{},
		nil,
	},
	{
		CONSOLE,
		[]string{"nintendo"},
		[]string{},
		nil,
	},

	// When we don't know anything else we just mark all Opera Mini devices as phones.
	{
		PHONE,
		[]string{"opera mini"},
		[]string{},
		nil,
	},
}

// DeviceType will return the device type.
// Returns 1 when nothing matched.
func DeviceType(userAgent string) int64 {
	id, _, _ := find(deviceTypes, userAgent)
	return id
}
