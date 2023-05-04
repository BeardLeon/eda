package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	common "golang-template/src/common/definition"
	"golang-template/src/config"
	"golang-template/src/zaplog"
	"time"
)

var Logger = zaplog.Logger

type DBConfigInfoStruct struct {
	EdaTableName string
}

var dbConfigInfo = DBConfigInfoStruct{
	EdaTableName: "",
}

type DBOps struct {
	mc []*sql.DB

	// DBInfo struct
	DBConfig *config.DBConfig
}

func (ops *DBOps) Init(driverName string, dataSourceName string,
	edaTableName string) {
	Logger.Debug("", zap.String("driverName", driverName),
		zap.String("dataSourceName", dataSourceName))

	for i := 0; i < common.F_NUM_DB_CONN; i++ {
		mc, dbErr := sql.Open(driverName, dataSourceName)
		if dbErr != nil {
			Logger.Fatal(
				"DB connection",
				zap.Any("driverName", driverName),
				zap.Any("dataSourceName", dataSourceName),
				zap.Any("error", dbErr))
		}
		mc.SetMaxOpenConns(2000)
		mc.SetMaxIdleConns(1000)
		mc.SetConnMaxLifetime(time.Minute * 60)
		ops.mc = append(ops.mc, mc)
	}

	dbConfigInfo.EdaTableName = edaTableName
}
