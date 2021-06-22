package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Ip          string `json:"ip"`
	Port        int16  `json:"port"`
	Certificate string `json:"certificate"`
	Key         string `json:"key"`
}

var configuration *Configuration

func Load(filename string) (err error) {

	var newConfiguration Configuration

	configFile, err := os.Open(filename)
	defer configFile.Close()
	if nil != err {

		return
	}

	err = json.NewDecoder(configFile).Decode(&newConfiguration)
	if err != nil {

		return
	}

	configuration = &newConfiguration

	return
}

func GetConfiguration() Configuration {

	if nil != configuration {
		return *configuration
	} else {
		return GetDefault()
	}
}

var GetDefault = (func() func() Configuration {
	configuration := Configuration{
		Ip:          "::1",
		Port:        4343,
		Certificate: "server.crt",
		Key:         "server.key",
	}

	return (func() Configuration { return configuration })
})()
