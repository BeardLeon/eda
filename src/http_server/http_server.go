package http_server

import (
	"eda/src/config"
	"eda/src/dbops"
	"eda/src/edaPkg"
	"eda/src/zaplog"
	"fmt"
	"net/http"
)

type Server struct {
	Config      config.Config
	DirConfig   string
	DBConfig    config.DBConfig
	DirDBConfig string
	EdaPkg      edaPkg.EdaPkg
	DBOps       dbops.DBOps
}

var Logger = zaplog.Logger

func (server *Server) Init() {
	server.DBConfig.LoadXMLDBConfig(server.DirDBConfig)
	server.Config.LoadXMLConfig(server.DirConfig)

	dbConfigInfo0 := server.DBConfig.DbBase

	dataSourceName := fmt.Sprintf("%s://%s:%s", dbConfigInfo0.DBType,
		dbConfigInfo0.IPAddress, dbConfigInfo0.Port)
	var db dbops.DBOps
	db.Init(dataSourceName, dbConfigInfo0.DBName, dbConfigInfo0.TableName)
	server.EdaPkg.New(&db, &server.DBConfig)
}

func (server *Server) RegisterHttpHandler() {
	// All the handler func map for request from client
	var requestHandlers = map[string]func(http.ResponseWriter, *http.Request){
		"/import":    server.importFile,
		"/export":    server.exportFile,
		"/line":      server.line,
		"/component": server.component,
		"/file":      server.file,
	}
	for k := range requestHandlers {
		http.HandleFunc(k, requestHandlers[k])
	}
}

func (server *Server) importFile(writer http.ResponseWriter, req *http.Request) {
	// Logger.Info("importJsonFile()", zap.Any("http.Request", *req))
	server.EdaPkg.ImportJsonFile(writer, req)
}

func (server *Server) exportFile(writer http.ResponseWriter, req *http.Request) {
	// Logger.Info("exportJsonFile()", zap.Any("http.Request", *req))
	server.EdaPkg.ExportFile(writer, req)
}

func (server *Server) file(writer http.ResponseWriter, req *http.Request) {
	server.EdaPkg.File(writer, req)
}

func (server *Server) line(writer http.ResponseWriter, req *http.Request) {
	// Logger.Info("line()", zap.Any("http.Request", req))
	server.EdaPkg.Line(writer, req)
}

func (server *Server) component(writer http.ResponseWriter, req *http.Request) {
	// Logger.Info("component()", zap.Any("http.Request", *req))
	server.EdaPkg.Component(writer, req)
}
