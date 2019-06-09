package gconfig

import (
	"log"
	"testing"
)

type Config struct {
	Db_host string
	Db_user string
	Db_pass string
	Db_port string
	Db_name string
}

var config *Config

func TestLoadYml(t *testing.T) {
	LoadYml("conf/config.yml", &config)

}

func TestshowConfig(t *testing.T) {
	log.Println("db user\t\t:", config.Db_user)
	log.Println("db pass\t\t:", config.Db_pass)
	log.Println("db host\t\t:", config.Db_host)
	log.Println("db port\t\t:", config.Db_port)
	log.Println("db name\t\t:", config.Db_name)
}

func TestMain(m *testing.M) {
	log.Println("start -------------------------")
	m.Run()
	log.Println("finish-------------------------")
}

func TestWorkFlow(t *testing.T) {
	t.Run("LoadYml", TestLoadYml)
	t.Run("showConfig", TestshowConfig)
}
