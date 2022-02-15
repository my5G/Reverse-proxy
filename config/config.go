package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

type Config struct {
	Sctp struct {
		Ip   string `yaml: "ip"`
		Port int    `yaml: "port"`
	} `yaml: "sctp"`
	Http struct {
		Ip   string `yaml: "ip"`
		Port int    `yaml: "port"`
	} `yaml: "http"`
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func (c *Config) GetConf() *Config {
	Ddir := RootDir()

	configPath, err := filepath.Abs(Ddir + "/config/config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
