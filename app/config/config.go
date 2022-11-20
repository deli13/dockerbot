package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var Config Configuration

type Configuration struct {
	Credentials Credentials `yaml:"credentials"`
	Project     Project     `yaml:"project"`
}

type Credentials struct {
	TgCredentials string `yaml:"tgcredentials"`
}

type Project struct {
	AllowUser []int64 `yaml:"allowUser"`
}

func LoadConfig(path string) {
	file, err := ioutil.ReadFile(path)
	panicError(err)
	err = yaml.Unmarshal(file, &Config)
	panicError(err)
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}
