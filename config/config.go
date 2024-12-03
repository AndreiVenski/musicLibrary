package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	Server       ServerConfig
	Postgres     PostgresConfig
	MusicService MusicServiceConfig
	HttpClient   HTTPClientConfig
	Logger       LoggerConfig
}

type ServerConfig struct {
	RunningPort string `envconfig:"SERVER_RUNNINGPORT"`
}

type LoggerConfig struct {
	LogLevel string `envconfig:"LOGGER_LOGLEVEL"`
}
type PostgresConfig struct {
	PostgresqlHost     string `envconfig:"POSTGRESQL_HOST"`
	PostgresqlPort     string `envconfig:"POSTGRESQL_PORT"`
	PostgresqlUser     string `envconfig:"POSTGRESQL_USER"`
	PostgresqlPassword string `envconfig:"POSTGRESQL_PASSWORD"`
	PostgresqlDbname   string `envconfig:"POSTGRESQL_DBNAME"`
}

type MusicServiceConfig struct {
	MusicAPIURL string `envconfig:"MUSICSERVICE_APIURL"`
}

type HTTPClientConfig struct {
	Timeout             time.Duration `envconfig:"HTTPCLIENT_TIMEOUT"`
	IdleConnTimeout     time.Duration `envconfig:"HTTPCLIENT_IDLECONNTIMEOUT"`
	DialTimeout         time.Duration `envconfig:"HTTPCLIENT_DIALTIMEOUT"`
	KeepAlive           time.Duration `envconfig:"HTTPCLIENT_KEEPALIVE"`
	MaxIdleConns        int           `envconfig:"HTTPCLIENT_MAXIDLECONNS"`
	MaxIdleConnsPerHost int           `envconfig:"HTTPCLIENT_MAXIDLECONNSPERHOST"`
	DisableCompression  bool          `envconfig:"HTTPCLIENT_DISABLECOMPRESSION"`
}

func InitConfig(path string) (*Config, error) {
	var cfg Config
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}
	err = envconfig.Process("", &cfg)
	return &cfg, err
}
