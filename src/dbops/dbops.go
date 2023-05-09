package dbops

import (
	"context"
	"eda/src/config"
	"eda/src/edaPkg"
	"eda/src/zaplog"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

var Logger = zaplog.Logger

type DBConfigInfoStruct struct {
	EdaTableName string
	EdaCollName  string
}

var dbConfigInfo = DBConfigInfoStruct{
	EdaTableName: "",
	EdaCollName:  "",
}

type DBOps struct {
	cli *qmgo.QmgoClient
	// DBInfo struct
	DBConfig *config.DBConfig
}

func (ops *DBOps) Init(dataSourceName string, database string,
	collection string) {
	Logger.Debug("", zap.String("dataSourceName", dataSourceName),
		zap.String("database", database), zap.String("collection", collection))
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{
		Uri:      dataSourceName,
		Database: database,
		Coll:     collection})
	if err != nil {
		Logger.Error("mongodb init error", zap.Error(err))
	}
	ops.cli = cli
	dbConfigInfo.EdaTableName = collection
}

func (ops *DBOps) GetCli() *qmgo.QmgoClient {
	return ops.cli
}

func (ops *DBOps) InsertLine(title string, line edaPkg.Line) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"title": title},
		bson.M{"$push": bson.M{"lines": bson.M{
			"sx": line.StartY,
			"sy": line.StartY,
			"ex": line.EndX,
			"ey": line.EndY},
		}})
	if err != nil {
		Logger.Error("insert line", zap.Error(err),
			zap.Any("title", title), zap.Any("line", line))
	}
}

func (ops *DBOps) DeleteLine(title string, line edaPkg.Line) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"title": title},
		bson.M{"$pull": bson.M{"lines": bson.M{
			"sx": line.StartY,
			"sy": line.StartY,
			"ex": line.EndX,
			"ey": line.EndY},
		}})
	if err != nil {
		Logger.Error("delete line", zap.Error(err),
			zap.Any("title", title), zap.Any("line", line))
	}
}

func (ops *DBOps) UpdateLine(title string, preLine, curLine edaPkg.Line) {
	ops.DeleteLine(title, preLine)
	ops.InsertLine(title, curLine)
}

func (ops *DBOps) InsertComponent(title string, component edaPkg.Component) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"title": title},
		bson.M{"$push": bson.M{"component": bson.M{
			"id":    component.Id,
			"name":  component.Name,
			"shape": component.Shape,
			"pin":   component.Pin,
		}}})
	if err != nil {
		Logger.Error("insert component", zap.Error(err),
			zap.Any("title", title), zap.Any("component", component))
	}
}

func (ops *DBOps) DeleteComponent(title string, component edaPkg.Component) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"title": title},
		bson.M{"$pull": bson.M{"component": bson.M{
			"id":    component.Id,
			"name":  component.Name,
			"shape": component.Shape,
			"pin":   component.Pin,
		}}})
	if err != nil {
		Logger.Error("delete component", zap.Error(err),
			zap.Any("title", title), zap.Any("component", component))
	}
}

func (ops *DBOps) UpdateComponent(title string, preComponent, curComponent edaPkg.Component) {
	ops.DeleteComponent(title, preComponent)
	ops.InsertComponent(title, curComponent)
}
