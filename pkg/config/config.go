package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config type
type Config struct {
	Server ServerConfiguration
	DB     DB
}

// DB Database connection
type DB struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Database     string `yaml:"database"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Sslmode      string `yaml:"sslmode"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	Driver       string `yaml:"driver"`
}

// ServerConfiguration server configuration
type ServerConfiguration struct {
	Port string
}

// ConnectURL returns the one liner connection string for postgres.
func (c *Config) ConnectURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database)
}

// Load load file of configuration
func Load(file string) (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName(file)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
