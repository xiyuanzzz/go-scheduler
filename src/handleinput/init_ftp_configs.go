package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"runtime"
)

type FtpConfigs struct {
	FtpIp string `yaml:"ftp_ip"`
	FtpPort string `yaml:"ftp_port"`
	FtpUsername string `yaml:"ftp_username"`
	FtpPassword string `yaml:"ftp_password"`
	FileRootPath string `yaml:"file_root_path"`
}

var ftpConfigs FtpConfigs

func init() {
	ftpConfigs.GetConfigs("ftp_configs.yml")
}

func (conf *FtpConfigs) GetConfigs(configFileName string) {
	fmt.Println("... 正在获取FTP连接信息")

	var delimiter string
	osType := runtime.GOOS
	if osType == "linux" {
		delimiter = "/"
	} else if osType == "windows" {
		delimiter = "\\"
	}

	path := "src" + delimiter + "configs" + delimiter + configFileName
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Read db configs file error: %v", err)
	}

	err = yaml.Unmarshal(file, conf)
	if err != nil {
		log.Printf("Resolve db configs file error: %v", err)
	}
}

/*func main() {
        fmt.Printf("%+v\n", dbConfigs)
}*/
