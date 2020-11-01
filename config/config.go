package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	Db dbURI
}

type dbURI struct {
	Driver   string `yaml:"driver"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Options  string `yaml:"options"`
}

// GetDbURI returns database URI formatted from config file fields
// TODO: implement option to get config from env var
func GetDbURI(filePath string) (string, error) {
	yamlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var c config
	err = yaml.Unmarshal(yamlContent, &c)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if c.Db.Driver == "" || c.Db.User == "" || c.Db.Password == "" || c.Db.Host == "" {
		return "", errors.New("missing uri values")
	}

	uri := fmt.Sprintf("%v://%v:%v@%v", c.Db.Driver, c.Db.User, c.Db.Password, c.Db.Host)
	if c.Db.Options != "" {
		uri = fmt.Sprintf("%v/?%v", uri, c.Db.Options)
	}
	return uri, nil
}
