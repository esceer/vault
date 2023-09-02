package config

// Config is the structure for the configuration data read from the env
type Config struct {
	ServerAddress string `envconfig:"SERVER_ADDRESS" default:":8080"`

	DatabaseDriver string `envconfig:"DATABASE_DRIVER" default:"sqlite3"`
	DataSource     string `envconfig:"DATA_SOURCE" default:"vault.db"`

	LogLevel             int  `envconfig:"LOG_LEVEL" default:"0"` // -1: trace, 0: debug, 1: info, 2: warn
	HumanFriendlyLogging bool `envconfig:"HUMAN_FRIENDLY_LOGGING" default:"true"`
}

// NewConfig returns a config instance initialized with default values
func NewConfig() *Config {
	return &Config{}
}
