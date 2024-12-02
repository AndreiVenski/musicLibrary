package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Kafka    KafkaConfig
}

type ServerConfig struct {
	JWTSecret          string `envconfig:"SERVER_JWTSECRET"`
	RunningPort        string `envconfig:"SERVER_RUNNINGPORT"`
	RefreshTokenSecret string `envconfig:"SERVER_REFRESHTOKENSECRET"`
	Vector4            string `envconfig:"SERVER_VEKTOR4"`
}

type PostgresConfig struct {
	PostgresqlHost     string `envconfig:"POSTGRESQL_HOST"`
	PostgresqlPort     string `envconfig:"POSTGRESQL_PORT"`
	PostgresqlUser     string `envconfig:"POSTGRESQL_USER"`
	PostgresqlPassword string `envconfig:"POSTGRESQL_PASSWORD"`
	PostgresqlDbname   string `envconfig:"POSTGRESQL_DBNAME"`
}

type KafkaConfig struct {
	SendMessageTimeout int    `envconfig:"KAFKA_SENDMESSAGETIMEOUT"`
	UserTopic          string `envconfig:"KAFKA_USERTOPIC"`
	Port               string `envconfig:"KAFKA_PORT"`
	Host               string `envconfig:"KAFKA_HOST"`
	PingTopic          string `envconfig:"Kafka_PINGTOPIC"`
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
