package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"log"
	"strings"
)

func connectFtp() *ftp.ServerConn {
	fmt.Println("... 正在等待FTP连接")

	ftpAddr := ftpConfigs.FtpIp + ":" + ftpConfigs.FtpPort
	conn, err := ftp.Dial(ftpAddr)
	if err != nil {
		log.Printf("Dial ftp server error: %v", err)
	}

	err = conn.Login(ftpConfigs.FtpUsername, ftpConfigs.FtpPassword)
	if err != nil {
		log.Printf("Login ftp server error: %v", err)
	}

	fmt.Println("*** FTP连接正常 ***")
	return conn
}

func resolvePath(path string) (dir, filename string) {
	fmt.Println("\n... 正在解析镜像文件下载路径")

	pathSlice := strings.Split(path, "/")

	ftpserver := pathSlice[2]
	for i := 3; i < len(pathSlice)-1; i++ {
		dir += pathSlice[i] + "/"
	}
	filename = pathSlice[len(pathSlice)-1]

	fmt.Printf("FTP服务器：%s\n", ftpserver)
	fmt.Printf("FTP文件路径：%s\n", dir)
	fmt.Printf("FTP文件名：%s\n", filename)
	return
}

func retrFile(conn *ftp.ServerConn, dir, filename string) *ftp.Response {
	fmt.Println("\n... 正在查找镜像文件")

	err := conn.ChangeDir(dir)
	if err != nil {
		log.Printf("Change ftp directory error: %v", err)
	}

	resp, err := conn.Retr(filename)
	if err != nil {
		log.Printf("Retrieve ftp file error: %v", err)
	}
	return resp
}

func copyFile(resp *ftp.Response, version, filename string) error {
	fmt.Println("... 正在读写镜像文件")

	buffer, err := ioutil.ReadAll(resp)
	if err != nil {
		log.Printf("Read ftp response error: %v", err)
		return err
	}

	writePath := ftpConfigs.FileRootPath + version + "/" + filename
	err = ioutil.WriteFile(writePath, buffer, 0644)
	if err != nil {
		log.Printf("Write ftp response error: %v", err)
		return err
	}

	fmt.Printf("*** 镜像文件 %s 已拷贝至文件夹 %s ***\n", filename, version)
	return nil
}

func downloadFile(filepath, version string) {
	conn := connectFtp()
	dir, filename := resolvePath(filepath)

	resp := retrFile(conn, dir, filename)
	defer resp.Close()

	err := copyFile(resp, version, filename)
	if err != nil {
		fmt.Printf("*** 镜像文件 %s 读写失败 ***\n", filename)
	} else {
		fmt.Println("*** 镜像文件下载完成 ***")
	}
}

func main() {
	downloadFile("ftp://10.151.12.52/01ALL/01FW/OpenPower/CP5466G2/BIOS/CP5466G2_BIOS_3.1.00_General_20200511/Checksum.txt", "")
}