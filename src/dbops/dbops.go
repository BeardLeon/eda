package dbops

import (
	"context"
	"eda/src/common"
	"eda/src/config"
	"eda/src/zaplog"
	"encoding/json"
	"fmt"

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

type File struct {
	Id          string             `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Components  []common.Component `bson:"components"`
	Lines       []common.Line      `bson:"lines"`
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
	err := ops.GetCli().UpdateOne(ctx, bson.M{"_id": line.OId},
		bson.M{"$push": bson.M{"lines": bson.M{
			"sx": line.StartY,
			"sy": line.StartY,
			"ex": line.EndX,
			"ey": line.EndY},
		}})
	if err != nil {
		Logger.Error("insert line", zap.Error(err),
			zap.Any("_id", line.OId), zap.Any("line", line))
	}
}

func (ops *DBOps) DeleteLine(line common.Line) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"_id": line.OId},
		bson.M{"$pull": bson.M{"lines": bson.M{
			"sx": line.StartY,
			"sy": line.StartY,
			"ex": line.EndX,
			"ey": line.EndY},
		}})
	if err != nil {
		Logger.Error("delete line", zap.Error(err),
			zap.Any("_id", line.OId), zap.Any("line", line))
	}
}

func (ops *DBOps) UpdateLine(preLine, curLine common.Line) {
	ops.DeleteLine(preLine)
	ops.InsertLine(curLine)
}

func (ops *DBOps) InsertComponent(component common.Component) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"_id": component.OId},
		bson.M{"$push": bson.M{"components": bson.M{
			"id":       component.Id,
			"name":     component.Name,
			"shape":    component.Shape,
			"pin":      component.Pin,
			"position": component.Position,
		}}})
	if err != nil {
		Logger.Error("insert component", zap.Error(err),
			zap.Any("_id", component.OId), zap.Any("component", component))
	}
}

func (ops *DBOps) DeleteComponent(component common.Component) {
	ctx := context.Background()
	err := ops.GetCli().UpdateOne(ctx, bson.M{"_id": component.OId},
		bson.M{"$pull": bson.M{"components": bson.M{
			"id":       component.Id,
			"name":     component.Name,
			"shape":    component.Shape,
			"pin":      component.Pin,
			"position": component.Position,
		}}})
	if err != nil {
		Logger.Error("delete component", zap.Error(err),
			zap.Any("_id", component.OId), zap.Any("component", component))
	}
}

func (ops *DBOps) UpdateComponent(preComponent, curComponent common.Component) {
	ops.DeleteComponent(preComponent)
	ops.InsertComponent(curComponent)
}

func (ops *DBOps) GetFile(oid string) File {
	ctx := context.Background()
	var file File
	err := ops.GetCli().Find(ctx, bson.M{"_id": oid}).One(&file)
	if err != nil {
		Logger.Error("GetFile Find error", zap.Error(err))
	}
	return file
}

func (ops *DBOps) CreateFile(title, description string) string {
	ctx := context.Background()
	id := common.ShordGuidGenerator()
	_, err := ops.GetCli().InsertOne(ctx, bson.M{
		"_id":         id,
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
		zap.Any("Object id ", id))
	return fmt.Sprintf("%v", id)
}

func (ops *DBOps) ImportFile(content []byte) error {
	ctx := context.Background()
	var file File
	err := json.Unmarshal(content, &file)
	if err != nil {
		Logger.Error("ImportFile json Unmarshal", zap.Error(err))
		return err
	}
	_, err = ops.GetCli().InsertOne(ctx, file)
	if err != nil {
		Logger.Error("ImportFile InsertOne", zap.Error(err))
		return err
	}
	return nil
}
