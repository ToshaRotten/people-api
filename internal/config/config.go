package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

// Config ..
type Config struct {
	Env        string `yaml:"env" env:"APP_ENV" envDefault:"local"`
	Storage    `yaml:"storage" env-required:"true"`
	HTTPServer `yaml:"http_server" env-required:"true"`
}

// HTTPServer ..
type HTTPServer struct {
	Host        string        `yaml:"host" env:"HTTP_HOST" env-default:"localhost" env-required:"true"`
	Port        string        `yaml:"port" env:"HTTP_PORT" env-default:"8080" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" env-default:"5s" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"HTTP_IDLE_TIMEOUT" env-default:"5s" env-required:"true"`
}

// Storage ..
type Storage struct {
	Connection string `yaml:"connection" env:"DB_CONNECTION" env-required:"true"`
	Host       string `yaml:"host" env:"DB_HOST" env-required:"true"`
	Port       string `yaml:"port" env:"DB_PORT" env-required:"true"`
	Database   string `yaml:"database" env:"DB_DATABASE" env-required:"true"`
	Username   string `yaml:"username" env:"DB_USERNAME" env-required:"true"`
	Password   string `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
}

// MustLoadConfig - returns a pointer to Config, may panic
func MustLoadConfig(path string) *Config {
	if path == "" {
		log.Fatal("path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
