package config

import (
	"github.com/essentiaone/ess-atomic-swap/server"
	"github.com/joeshaw/envdecode"
	"github.com/pkg/errors"
)

// Configuration is a struct that contains server struct
type Configuration struct {
	BitcoinNodeAddr  string `env:"BTC_ADDR,required"`
	EthereumNodeAddr string `env:"ETH_ADDR,required"`
	Server           server.Settings
	SCDBAddr         string `env:"SCDB_ADDR,required"`
	PrivateKet       string `env:"PRIVATE_KEY,required"`
	IsProduction     bool   `env:"PRODUCTION,default=false"`
}

// Init that read env va-les
func Init() (*Configuration, error) {
	conf := &Configuration{}
	if err := envdecode.Decode(conf); err != nil {
		return nil, errors.Wrap(err, "cannot read env variables")
	}
	return conf, nil
}
