# go-config
golang读取配置文件，`json`|`yaml`配置文件包


```golang
type Config struct {
	Db_host string
	Db_user string
	Db_pass string
	Db_port string
	Db_name string
}

var config *Config

LoadYml("conf/config.yml", &config)
```



```golang
2018/07/25 14:50:09 db user		: root
2018/07/25 14:50:09 db pass		: 123456
2018/07/25 14:50:09 db host		: 127.0.0.1
2018/07/25 14:50:09 db port		: 3306
2018/07/25 14:50:09 db name		: test
PASS
ok  	gconfig	0.007s
```
