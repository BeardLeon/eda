package http_server

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"golang-template/src/config"
	"golang-template/src/dbops"
	"golang-template/src/http_handler"
	"golang-template/src/zaplog"
	"net/http"
)

type Server struct {
	Config      config.Config
	DirConfig   string
	DBConfig    config.DBConfig
	DirDBConfig string
	DB          *sql.DB
	EdaHandler  http_handler.EdaHandler
}

var Logger = zaplog.Logger

func (server *Server) Init() {
	server.DBConfig.LoadXMLDBConfig(server.DirDBConfig)
	server.Config.LoadXMLConfig(server.DirConfig)
	server.DB = server.DBConfig.ConnDB()

	dbConfigInfo0 := server.DBConfig.DbBase

	var dbOps dbops.DBOps
	dataSourceName := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", dbConfigInfo0.Username,
		dbConfigInfo0.Password, dbConfigInfo0.IPProtocol,
		dbConfigInfo0.IPAddress, dbConfigInfo0.Port, dbConfigInfo0.DBName)
	driverName := dbConfigInfo0.DBType

	dbOps.Init(driverName, dataSourceName, dbConfigInfo0.TableName)

}

func (server *Server) RegisterHttpHandler() {
	// All the handler func map for request from client
	var requestHandlers = map[string]func(http.ResponseWriter, *http.Request){
		"/import": server.importJsonFile,
	}
	for k := range requestHandlers {
		http.HandleFunc(k, requestHandlers[k])
	}
}

func (server *Server) importJsonFile(writer http.ResponseWriter, req *http.Request) {
	Logger.Info("importJsonFile()", zap.Any("http.Request", *req))
	server.EdaHandler.ImportJsonFile(writer, req)
}
