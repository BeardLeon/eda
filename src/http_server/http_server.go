package http_server

import (
	"database/sql"
	"eda/src/config"
	"eda/src/dbops"
	"eda/src/edaPkg"
	"eda/src/zaplog"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	Config      config.Config
	DirConfig   string
	DBConfig    config.DBConfig
	DirDBConfig string
	DB          *sql.DB
	EdaPkg      edaPkg.EdaPkg
}

var Logger = zaplog.Logger

func (server *Server) Init() {
	server.DBConfig.LoadXMLDBConfig(server.DirDBConfig)
	server.Config.LoadXMLConfig(server.DirConfig)
	server.DB = server.DBConfig.ConnDB()

	dbConfigInfo0 := server.DBConfig.DbBase

	var dbOps dbops.DBOps
	dataSourceName := fmt.Sprintf("%s://%s:%s", dbConfigInfo0.DBType,
		dbConfigInfo0.IPAddress, dbConfigInfo0.Port)

	dbOps.Init(dataSourceName, dbConfigInfo0.DBName, dbConfigInfo0.TableName)

}

func (server *Server) RegisterHttpHandler() {
	// All the handler func map for request from client
	var requestHandlers = map[string]func(http.ResponseWriter, *http.Request){
		"/import":    server.importJsonFile,
		"/export":    server.exportJsonFile,
		"/line":      server.line,
		"/component": server.component,
	}
	for k := range requestHandlers {
		http.HandleFunc(k, requestHandlers[k])
	}
}

func (server *Server) importJsonFile(writer http.ResponseWriter, req *http.Request) {
	Logger.Info("importJsonFile()", zap.Any("http.Request", *req))
	server.EdaPkg.ImportJsonFile(writer, req)
}

func (server *Server) exportJsonFile(writer http.ResponseWriter, req *http.Request) {
	Logger.Info("exportJsonFile()", zap.Any("http.Request", *req))
	server.EdaPkg.ExportJsonFile(writer, req)
}

func (server *Server) line(writer http.ResponseWriter, req *http.Request) {
	Logger.Info("line()", zap.Any("http.Request", *req))
	server.EdaPkg.Line(writer, req)
}

func (server *Server) component(writer http.ResponseWriter, req *http.Request) {
	Logger.Info("component()", zap.Any("http.Request", *req))
	server.EdaPkg.Component(writer, req)
}
