package functions

import (
	"errors"
	"sort"
	"strings"

	"github.com/hiagomf/networkutils/data/mac_address/singleton"
)

func GetManufacturer(macAddress *string) (manufacturerData *ManufacturerData, err error) {
	if macAddress == nil {
		return manufacturerData, errors.New("mac address cannot be nil")
	}

	manufacturerData = new(ManufacturerData) // init pointer
	macAddressStr := *macAddress
	selectedManufacturerRange := macAddressStr[0:8]

	lines := singleton.GetInstance() // Buscando no singleton os dados

	// Percorrendo linha a linha
	for i := 0; i < len(lines); i++ {
		manufacturerRef := strings.Split(lines[i], "\t")
		mainManufacturer := strings.Split(lines[i+1], "\t")

		toCompare := []string{selectedManufacturerRange, manufacturerRef[0], mainManufacturer[0]}
		sort.Strings(toCompare)

		if toCompare[1] == selectedManufacturerRange && selectedManufacturerRange != mainManufacturer[0] {
			previousManufacturer := strings.Split(lines[i-1], "\t")
			nextManufacturer := strings.Split(lines[i+2], "\t")

			manufacturerData.Mac = macAddressStr
			manufacturerData.Manufacturer = manufacturerRef[0] + " " + manufacturerRef[1] + " " + manufacturerRef[2]
			manufacturerData.PreviousManufacturer = previousManufacturer[0] + " " + previousManufacturer[1] + " " + previousManufacturer[2]
			manufacturerData.NextManufacturer = nextManufacturer[0] + " " + nextManufacturer[1] + " " + nextManufacturer[2]
			break
		}
	}
	return manufacturerData, nil
}
