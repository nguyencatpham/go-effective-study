package config

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server *Server      `yaml:"server,omitempty"`
	DB     *Database    `yaml:"database,omitempty"`
	JWT    *JWT         `yaml:"jwt,omitempty"`
	App    *Application `yaml:"application,omitempty"`
}

// Database holds data necessery for database configuration
type Database struct {
	DBName       string `yaml:"dbName,omitempty" json:"dbname,omitempty"`
	PSN          string `yaml:"psn,omitempty" json:"psn,omitempty"`
	LogQueries   bool   `yaml:"log_queries,omitempty"`
	PSNBase      string `yaml:"psnBase,omitempty" json:"psnBase,omitempty"`
	Log          bool
	CreateSchema bool
	Timeout      int
	MaxRetries   int `yaml:"maxRetries,omitempty" json:"maxRetries,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	Host          string `yaml:"host,omitempty"`
	Port          string `yaml:"port,omitempty"`
	Debug         bool   `yaml:"debug,omitempty"`
	ReadTimeout   int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout  int    `yaml:"write_timeout_seconds,omitempty"`
	SwaggerUIPath string `yaml:"swagger_ui_path,omitempty"`
	SwaggerJSON   string `yaml:"swagger_json,omitempty"`
	Schemes 	  []string `yaml:"schemes,omitempty"`
}

// JWT holds data necessery for JWT configuration
type JWT struct {
	Secret           string `yaml:"secret,omitempty"`
	Duration         int    `yaml:"duration_minutes,omitempty"`
	RefreshDuration  int    `yaml:"refresh_duration_minutes,omitempty"`
	MaxRefresh       int    `yaml:"max_refresh_minutes,omitempty"`
	SigningAlgorithm string `yaml:"signing_algorithm,omitempty"`
}

// Application holds application configuration details
type Application struct {
	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
}

type ErrorDict struct {
	ErrorList []ErrorMessage `yaml:"error_list,omitempty" json:"errorList,omitempty"`
}

func LoadConfig(path string) (*Configuration, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

func LoadErrorList() (*ErrorDict, error) {
	path := "./config/errors.yaml"
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading error list file, %s", err)
	}
	var cfg = new(ErrorDict)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

type ErrorMessage struct {
	Type string `yaml:"type,omitempty" json:"type,omitempty"`
	Text string `yaml:"text,omitempty" json:"text,omitempty"`
}
