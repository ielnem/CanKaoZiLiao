package main

import (
	"flag"
	"fmt"
)

var fileName string
var excelName string
var cmd string
var upPath string

func init() {
	flag.StringVar(&fileName, "f", "", "scp文件名")
	flag.StringVar(&cmd, "c", "", "ssh命令")
	flag.StringVar(&excelName, "e", "./servers.xlsx", "服务器excel文件名，默认./servers.xlsx")
	flag.StringVar(&upPath, "p", "/home/", "scp目标目录，默认/home/")

	flag.Parse()

}

func main() {

	if (fileName == "" && cmd == "") || (fileName != "" && cmd != "") {
		flag.Usage()
		return
	}

	serverList, err := readXlsx()
	if err != nil {
		fmt.Printf("read %s failed, err %s \n", excelName, err.Error())
	}

	for _, svr := range serverList {
		doServerJob(svr)
	}
}

func doServerJob(svr *Server) {
	if fileName != "" {
		bytesReader := ReadFile()
		err := doScpWork(bytesReader, svr)
		if err != nil {
			fmt.Printf("scp failed %s \n", err.Error())
		}
	}

	if cmd != "" {
		sshClient := createSSHClient(svr)
		defer sshClient.Close()
		runCMD(cmd, sshClient)
	}

}
