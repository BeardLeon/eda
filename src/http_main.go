package src

import (
	"fmt"
	server "golang-template/src/http_server"
	_ "golang-template/src/zaplog"
	"log"
	"net/http"
	"os"
)

func logOutput() {
	file := "./" + "log" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // set logfile
	log.SetPrefix("[logTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

func init() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd() error! \n")
	}
	dirConfig := dir + "/conf/server_config.xml"
	dbConfig := dir + "/conf/db_config.xml"
	log.Println("Directory of server_config file:", dirConfig)

	server := server.Server{}
	server.DirConfig = dirConfig
	server.DirDBConfig = dbConfig

	server.Init()
	server.RegisterHttpHandler()
	ServerAddr = server.Config.AddrConfig.IpAddr + ":" + server.Config.AddrConfig.HttpPort

	log.Println("Finish init() ! \n\n ")
}

var ServerAddr string

func main() {
	log.Println("main() serverAddress: ", ServerAddr, " \n\n ")

	err := http.ListenAndServe(ServerAddr, nil)
	if err != nil {
		fmt.Println("Listen to http requests failed", err)
	}
}
