package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	Google struct {
		ApiKey string
		Model  string
	}
	Server struct {
		Host string
	}
	DB struct {
		Database string
		Uri      string
	}
	JWT struct {
		Secret string
	}
	SSO struct {
		Issuer       string
		AuthUrl      string
		TokenUrl     string
		ClientId     string
		Scope        string
		RedirectUri  string
		ResponseType string
		LogoutUrl    string
		ClientSecret string
		GrantType    string
	}
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./server/config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Failed to load config", err)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		log.Println("Failed to unmarshal config", err)
	}

	cfg = &c

}

func Get() *Config {
	once.Do(loadConfig)
	return cfg
}
