package main

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var config struct {
	Server struct {
		Host string
		Port int
	}
	Users map[string]struct {
		Password  string
		Disabled  bool `yaml:"disabled,omitempty"`
		WritePath []string
		ReadPath  []string
	}
}

var configPath = flag.String("c", "goftp.yml", "config file location")

func readConfig(fileName string) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.UnmarshalStrict(f, &config)
	if err != nil {
		log.Fatal(err)
	}

}
