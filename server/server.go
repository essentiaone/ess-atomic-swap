package server

//Settings struct is a simple settings for the server
type Settings struct {
	SSL  bool
	Port string `env:"ESS_ATOMIC_SWAP_APP_PORT"`
}
