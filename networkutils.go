package networkutils

import "github.com/hiagomf/mac-utils/functions"

func GetManufacturer(macAddress *string) (manufacturerData *functions.ManufacturerData, err error) {
	return functions.GetManufacturer(macAddress)
}
