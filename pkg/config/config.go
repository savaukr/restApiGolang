package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	HttpAddr string
	DBAddr string

}

func NewConfig() (*Config, error) {
	var s Config
    err := envconfig.Process("MR", &s)
    if err != nil {
        return nil, err
    }
	return s, nil
}