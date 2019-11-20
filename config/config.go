package config

import "flag"

const (
	serviceName    = `UserService`
	defaultAddress = `localhost:12345`
)

type Config struct {
	ServiceName    string
	ServiceAddress string
}

func New() (cfg Config) {
	cfg.ServiceName = serviceName

	flag.StringVar(&cfg.ServiceAddress, "addr", defaultAddress, "Usage [host:port]")
	flag.Parse()
	return cfg
}
