package singleton

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var lock = &sync.Mutex{}
var manufacturersInstance []string

func GetInstance() []string {
	if len(manufacturersInstance) == 0 {
		manufacturersInstance = refreshManufacturerInstace()
	}
	return manufacturersInstance
}

func refreshManufacturerInstace() []string {
	lock.Lock()
	defer lock.Unlock()

	url := "https://gitlab.com/wireshark/wireshark/-/raw/master/manuf"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	archive, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	initData := strings.SplitN(string(archive), "00:00:00", 2)
	manufacturersInstance = strings.Split(initData[1], "\n")
	manufacturersInstance[0] = strings.Replace(manufacturersInstance[0], "\t", "", 1)
	return manufacturersInstance
}
