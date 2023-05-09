package dbops

import (
	"context"
	"eda/src/common"
	"eda/src/config"
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
	InsertId interface{}
	cli      *qmgo.QmgoClient
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

func (ops *DBOps) InsertLine(line common.Line) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"_id": ops.InsertId},
		bson.M{"$push": bson.M{"lines": bson.M{
			"sx": line.StartY,
			"sy": line.StartY,
			"ex": line.EndX,
			"ey": line.EndY},
		}})
	if err != nil {
		Logger.Error("insert line", zap.Error(err),
			zap.Any("_id", ops.InsertId), zap.Any("line", line))
	}
}

func (ops *DBOps) DeleteLine(line common.Line) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"_id": ops.InsertId},
		bson.M{"$pull": bson.M{"lines": bson.M{
			"sx": line.StartY,
			"sy": line.StartY,
			"ex": line.EndX,
			"ey": line.EndY},
		}})
	if err != nil {
		Logger.Error("delete line", zap.Error(err),
			zap.Any("_id", ops.InsertId), zap.Any("line", line))
	}
}

func (ops *DBOps) UpdateLine(preLine, curLine common.Line) {
	ops.DeleteLine(preLine)
	ops.InsertLine(curLine)
}

func (ops *DBOps) InsertComponent(component common.Component) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"_id": ops.InsertId},
		bson.M{"$push": bson.M{"components": bson.M{
			"id":    component.Id,
			"name":  component.Name,
			"shape": component.Shape,
			"pin":   component.Pin,
		}}})
	if err != nil {
		Logger.Error("insert component", zap.Error(err),
			zap.Any("_id", ops.InsertId), zap.Any("component", component))
	}
}

func (ops *DBOps) DeleteComponent(component common.Component) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"_id": ops.InsertId},
		bson.M{"$pull": bson.M{"components": bson.M{
			"id":    component.Id,
			"name":  component.Name,
			"shape": component.Shape,
			"pin":   component.Pin,
		}}})
	if err != nil {
		Logger.Error("delete component", zap.Error(err),
			zap.Any("_id", ops.InsertId), zap.Any("component", component))
	}
}

func (ops *DBOps) UpdateComponent(preComponent, curComponent common.Component) {
	ops.DeleteComponent(preComponent)
	ops.InsertComponent(curComponent)
}

func (ops *DBOps) CreateFile(title, description string) {
	ctx := context.Background()
	res, err := ops.GetCli().InsertOne(ctx, bson.M{
		"title":       title,
		"description": description,
		"components":  bson.A{},
		"lines":       bson.A{},
	})
	if err != nil {
		Logger.Error("create file", zap.Error(err),
			zap.Any("title", title),
			zap.Any("description", description))
	}
	Logger.Info("Successfully create file. ",
		zap.Any("Object id ", res.InsertedID))
	ops.InsertId = res.InsertedID
}
