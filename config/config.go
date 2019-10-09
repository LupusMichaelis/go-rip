package config

type Configuration struct {
	Ip          string
	Port        int16
	Certificate string
	Key         string
}

var GetConfiguration = (func() func() Configuration {
	configuration := Configuration{
		Ip:          "::1",
		Port:        4343,
		Certificate: "server.crt",
		Key:         "server.key",
	}

	return (func() Configuration { return configuration })
})()
