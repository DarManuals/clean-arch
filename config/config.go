package config

const (
	ServiceName = `UserService`
	defaultPort = `:12345`
)

type Config struct {
	ServiceName string
	ServicePort string
}

func New() (cfg Config) {

	cfg.ServiceName = ServiceName
	cfg.ServicePort = defaultPort

	return cfg
}
