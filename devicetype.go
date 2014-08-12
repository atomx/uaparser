package uaparser

const (
	PC          Id = iota + 2 // 2
	PHONE                     // 3
	TABLET                    // 4
	CONSOLE                   // 5
	SET_TOP_BOX               // 6
)

var deviceTypes = []pattern{
	// Set-top box
	pattern{
		SET_TOP_BOX,
		[]string{"media center pc"},
		[]string{},
		nil,
	},
	pattern{
		SET_TOP_BOX,
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
	pattern{
		TABLET,
		[]string{"tablet pc"},
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
}

// Will return the device type.
// Returns 1 when nothing matched.
func DeviceType(userAgent string) Id {
	id, _ := find(deviceTypes, userAgent)
	return id
}
