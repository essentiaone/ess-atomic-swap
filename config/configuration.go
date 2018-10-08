package config

import (
	"github.com/essentiaone/ess-atomic-swap/server"
	"github.com/joeshaw/envdecode"
)

//Configuration is a struct that contains server struct
type Configuration struct {
	Server server.Settings
}

//Init - func that read env va-les
func Init() (Configuration, error) {
	conf := Configuration{}
	if err := envdecode.Decode(&conf); err != nil {
		panic(err)
	}
	return conf, nil
}
