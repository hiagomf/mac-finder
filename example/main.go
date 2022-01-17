package main

import (
	"log"

	"github.com/hiagomf/mac-utils/functions"
	"github.com/hiagomf/mac-utils/utils"
)

func main() {
	v1, _ := functions.GetManufacturer(utils.StringPointer("D0:94:66:DF:5C:89"))
	log.Println(
		"\n------- "+v1.Manufacturer+" -----------",
		"\nMAC:         ", v1.Mac,
		"\nAnterior:    ", v1.PreviousManufacturer,
		"\nSelecionado: ", v1.Manufacturer,
		"\nPróximo:     ", v1.NextManufacturer,
	)

	v2, _ := functions.GetManufacturer(utils.StringPointer("58:10:8C:97:59:96"))
	log.Println(
		"\n------- "+v2.Manufacturer+" -----------",
		"\nMAC:         ", v2.Mac,
		"\nAnterior:    ", v2.PreviousManufacturer,
		"\nSelecionado: ", v2.Manufacturer,
		"\nPróximo:     ", v2.NextManufacturer,
	)
}
