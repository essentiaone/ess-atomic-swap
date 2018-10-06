package config

import (
	"github.com/caarlos0/env"
	"github.com/essentiaone/ess-atomic-swap/server"
)

//Configuration is a struct that contains server struct
type Configuration struct {
	Server server.Settings
}

//Init - func that read env va-les
func Init() (Configuration, error) {
	conf := Configuration{}
	env.Parse(&conf)
	return conf, nil
}
