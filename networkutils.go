package networkutils

import "github.com/hiagomf/networkutils/functions"

func GetManufacturer(macAddress *string) (manufacturerData *functions.ManufacturerData, err error) {
	return functions.GetManufacturer(macAddress)
}
