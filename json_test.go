package gconfig

import (
	"log"
)

type JsonConfig struct {
	DbHost string
	DbUser string
	DbPass string
	DbPort string
	DbName string
}

var jsonConfig *JsonConfig

func init() {
	LoadJson("conf/config.json", &jsonConfig)
	log.Println(jsonConfig)
}
