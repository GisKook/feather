package conf

import (
	"encoding/json"
	"os"
)

type HttpConf struct {
	Addr     string
	AppID    string
	Interval int
}

type DBConf struct {
	Host   string
	Port   string
	User   string
	Passwd string
	DbName string
}

type Configuration struct {
	Http *HttpConf
	DB   *DBConf
}

var G_conf *Configuration

func ReadConfig(confpath string) (*Configuration, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	G_conf = &config

	return &config, err
}

func GetConf() *Configuration {
	return G_conf
}
