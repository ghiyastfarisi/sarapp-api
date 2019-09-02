package config

type (
	// Config is configuration struct
	Config struct {
		Application ApplicationConfig `mapstructure:"application"`
		Database    DatabaseConfig    `mapstructure:"database"`
	}
	// ApplicationConfig ...
	ApplicationConfig struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
		Port    int    `mapstructure:"port"`
	}
	// DatabaseConfig ...
	DatabaseConfig struct {
		Main DBStruct
	}
	// DBStruct ...
	DBStruct struct {
		Dsn           string
		Port          int
		Maxconnection int
	}
)
