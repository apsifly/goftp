package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"regexp"

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
	re1 := regexp.MustCompile("^([0-9]+[.]){3}[0-9]+")
	if !re1.MatchString(config.Server.Host) {
		ips, _ := net.LookupIP(config.Server.Host)
		if len(ips) == 0 {
			log.Fatal("can not resolve provided host")
		}
		//we need to know IP to handle PASV command. Dynamic binding for PASV is not implemented
		//TODO: use net.TCPConn.Local..
		config.Server.Host = ips[0].String()
	}
}
