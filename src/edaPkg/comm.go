package edaPkg

import (
	"eda/src/common"
	"eda/src/dbops"
	"eda/src/zaplog"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type EdaPkg struct {
	Id          string             `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Components  []common.Component `bson:"components"`
	Lines       []common.Line      `bson:"lines"`
}

var Logger = zaplog.Logger
var db = dbops.DBOps{}

func (pkg *EdaPkg) New() {
}

// ImportJsonFile need name and file
func (pkg *EdaPkg) ImportJsonFile(w http.ResponseWriter, r *http.Request) error {
	// TODO add string to db
	fmt.Println()
	return nil
}

// ExportJsonFile export name and file
func (pkg *EdaPkg) ExportJsonFile(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (pkg *EdaPkg) InsertComponent(w http.ResponseWriter, r *http.Request) {
	var component common.Component
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error("insert component ioutil.ReadAll error", zap.Error(err))
	}
	err = bson.Unmarshal(res, &component)
	if err != nil {
		Logger.Error("insert component bson.Unmarshal error", zap.Error(err))
	}
	db.InsertComponent(component)
}

func (pkg *EdaPkg) InsertLine(w http.ResponseWriter, r *http.Request) {
	var line common.Line
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error("insert line ioutil.ReadAll error", zap.Error(err))
	}
	err = bson.Unmarshal(res, &line)
	if err != nil {
		Logger.Error("insert line bson.Unmarshal error", zap.Error(err))
	}
	db.InsertLine(line)
}

func (pkg *EdaPkg) UpdateComponent(w http.ResponseWriter, r *http.Request) {
	var component [2]common.Component
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error("update component ioutil.ReadAll error", zap.Error(err))
	}
	err = bson.Unmarshal(res, &component)
	if err != nil {
		Logger.Error("update component bson.Unmarshal error", zap.Error(err))
	}
	db.UpdateComponent(component[0], component[1])
}

func (pkg *EdaPkg) UpdateLine(w http.ResponseWriter, r *http.Request) {
	var lines [2]common.Line
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error("update line ioutil.ReadAll error", zap.Error(err))
	}
	err = bson.Unmarshal(res, &lines)
	if err != nil {
		Logger.Error("update line bson.Unmarshal error", zap.Error(err))
	}
	db.UpdateLine(lines[0], lines[1])
}

func (pkg *EdaPkg) DeleteComponent(w http.ResponseWriter, r *http.Request) {
	var component common.Component
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error("delete component ioutil.ReadAll error", zap.Error(err))
	}
	err = bson.Unmarshal(res, &component)
	if err != nil {
		Logger.Error("delete component bson.Unmarshal error", zap.Error(err))
	}
	db.DeleteComponent(component)
}

func (pkg *EdaPkg) DeleteLine(w http.ResponseWriter, r *http.Request) {
	var line common.Line
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error("delete line ioutil.ReadAll error", zap.Error(err))
	}
	err = bson.Unmarshal(res, &line)
	if err != nil {
		Logger.Error("delete line bson.Unmarshal error", zap.Error(err))
	}
	db.DeleteLine(line)
}

func (pkg *EdaPkg) Line(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		pkg.InsertLine(w, req)
	} else if req.Method == http.MethodDelete {
		pkg.DeleteLine(w, req)
	} else if req.Method == http.MethodPut {
		pkg.UpdateLine(w, req)
	} else {
		return
	}
	return
}

func (pkg *EdaPkg) Component(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		pkg.InsertComponent(w, req)
	} else if req.Method == http.MethodDelete {
		pkg.DeleteComponent(w, req)
	} else if req.Method == http.MethodPut {
		pkg.UpdateComponent(w, req)
	} else {
		return
	}
	return
}
