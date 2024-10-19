package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		ConnectionString string `yaml:"connectionString"`
	} `yaml:"database"`
	LogLevel      string `yaml:"logLevel"`
	LeadgidApiKey string `yaml:"leadgidApiKey"`
	ChatGPTApiKey string `yaml:"chatGPTApiKey"`
}

type Service struct {
	Host         string `yaml:"host"`
	HttpPort     string `yaml:"httpPort"`
	GrpcPort     string `yaml:"grpcPort"`
	RateLimit    int    `yaml:"rateLimit"`
	WorkersCount int    `yaml:"workersCount"`
}

const (
	pathToConfig = "config.yaml"

	AppName = "banks-api"
)

func LoadConfig() (*Config, error) {
	rawYaml, err := os.ReadFile(pathToConfig)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(rawYaml, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) GetCheckoutPgConnectionString() string {
	return c.Database.ConnectionString
}

func (c *Config) GetLeadgidApiKey() string {
	return c.LeadgidApiKey
}

func (c *Config) GetChatGPTApiKey() string {
	return c.ChatGPTApiKey
}
