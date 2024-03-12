package configs

import "github.com/BurntSushi/toml"

const defaultName = "./configs/app.toml"

type Config struct {
	Name string              `toml:"name"`
	HTTP *HTTP               `toml:"http"`
	Env  map[string]Variable `toml:"env"`
}

type HTTP struct {
	Port        string `toml:"port"`
	SwaggerPath string `toml:"swagger_path"`
}

type Variable struct {
	Value string `toml:"value"`
}

func Parse() (*Config, error) {
	var config Config

	if _, err := toml.DecodeFile(defaultName, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
