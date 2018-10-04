package config

import (
	"github.com/essentiaone/ess-atomic-swap/server"
	// "github.com/spf13/viper"
)

//Configuration is a struct that contains server struct
type Configuration struct {
	Server server.Settings
}
