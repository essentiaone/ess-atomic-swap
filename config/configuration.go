package config

import (
	"github.com/essentiaone/ess-atomic-swap/server"
	"github.com/joeshaw/envdecode"
)

// Configuration is a struct that contains server struct
type Configuration struct {
	BitcoinNodeAddr  string `env:"BTC_ADDR, required"`
	EthereumNodeAddr string `env:"ETH_ADDR, required"`
	Server           server.Settings
}

// Init that read env va-les
func Init() (Configuration, error) {
	conf := Configuration{}
	if err := envdecode.Decode(&conf); err != nil {
		panic(err)
	}
	return conf, nil
}
