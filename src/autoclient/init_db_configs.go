package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

type DbConfigs struct {
	DbHostIP string `yaml:"db_host_ip"`
	DbHostPort string `yaml:"db_host_port"`
	LoginUsername string `yaml:"login_username"`
	LoginPassword string `yaml:"login_password"`
	DbName string `yaml:"db_name"`
}

var dbConfigs DbConfigs

func init() {
	dbConfigs.GetConfigs("db_configs.yml")
}

func (conf *DbConfigs) GetConfigs(configFileName string) {
	fmt.Println("... 正在获取数据库连接信息")
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Get current working directory error: %v", err)
	}

	osType := runtime.GOOS
	var delimiter string
	if osType == "linux" {
		delimiter = "/"
	} else if osType == "windows" {
		delimiter = "\\"
	}

	path := wd + delimiter + "configs" + delimiter + configFileName
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Read db configs file error: %v", err)
	}

	err = yaml.Unmarshal(file, conf)
	if err != nil {
		log.Printf("Resolve db configs file error: %v", err)
	}
	fmt.Println("*** 数据库连接信息已获取 ***")
}

/*func main() {
        fmt.Printf("%+v\n", dbConfigs)
}*/
