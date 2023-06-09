package edaPkg

import (
	"bufio"
	"eda/src/common"
	"eda/src/config"
	"eda/src/dbops"
	"eda/src/zaplog"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type EdaPkg struct {
	Id          string             `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Components  []common.Component `bson:"components"`
	Lines       []common.Line      `bson:"lines"`

	db *dbops.DBOps
}

type CreateFileReq struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type CreateFileResp struct {
	FileId string `json:"id"`
}

type ErrorResp struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

// generate json format response string
func RespGen(errCode int, errMsg string) (string, error) {
	resp := ErrorResp{
		ErrCode: errCode,
		ErrMsg:  errMsg,
	}

	jsonResp, err := json.Marshal(resp)
	return string(jsonResp), err
}

func ExportFileRespGen(file dbops.File) (string, error) {
	jsonResp, err := json.MarshalIndent(file, "", "\t")
	return string(jsonResp), err
}

func GetFileRespGen(file dbops.File) (string, error) {
	jsonResp, err := json.Marshal(file)
	return string(jsonResp), err
}

func CreateFileRespGen(fileId string) (string, error) {
	resp := CreateFileResp{
		FileId: fileId,
	}
	jsonResp, err := json.Marshal(resp)
	return string(jsonResp), err
}

var Logger = zaplog.Logger

func (pkg *EdaPkg) New(db *dbops.DBOps, dbConfig *config.DBConfig) {
	pkg.db = db
	pkg.db.DBConfig = dbConfig
}

// ImportJsonFile need name and file
func (pkg *EdaPkg) ImportJsonFile(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "multipart/form-data;")
	file, _, err := r.FormFile("file")
	if err != nil {
		Logger.Error("ImportJsonFile", zap.Error(err))
		jsonResp, _ := RespGen(common.K_PARSE_FILE_ERROR, "Form file failed")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		Logger.Error("ImportJsonFile", zap.Error(err))
		jsonResp, _ := RespGen(common.K_PARSE_FILE_ERROR, "Read file failed")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
	err = pkg.db.ImportFile(content)
	if err != nil {
		Logger.Error("ImportJsonFile", zap.Error(err))
		return
	}
	jsonResp, _ := RespGen(0, "Success")
	Logger.Info("ImportJsonFile", zap.Any("file", file))
	fmt.Fprintf(w, jsonResp+"\n")
}

// ExportJsonFile export name and file
func (pkg *EdaPkg) ExportFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	oid := r.URL.Query().Get("oid")
	file := pkg.db.GetFile(oid)
	data := fmt.Sprintf("attachment; filename=%s.json", file.Title)
	w.Header().Set("Content-Disposition", data)

	resp := pkg.db.GetFile(oid)

	writer := bufio.NewWriter(w)
	jsonResp, _ := ExportFileRespGen(resp)
	writer.WriteString(jsonResp)
	writer.Flush()
	Logger.Info("ExportFile", zap.Any("_id", oid))
	jsonResp, _ = RespGen(0, "Success")
	fmt.Fprintf(w, jsonResp+"\n")
}

func (pkg *EdaPkg) GetFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	oid := r.URL.Query().Get("oid")
	resp := pkg.db.GetFile(oid)
	jsonResp, _ := GetFileRespGen(resp)
	Logger.Info("GetFile", zap.Any("_id", oid))
	fmt.Fprintf(w, jsonResp+"\n")
}

func (pkg *EdaPkg) CreateFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var req CreateFileReq
	err := json.Unmarshal(body, &req)
	if err != nil {
		jsonResp, _ := RespGen(common.K_REQUEST_COMMAND_ERROR, "Invalid json format")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
	id := pkg.db.CreateFile(req.Title, req.Desc)
	Logger.Info("CreateFile", zap.Any("_id", id), zap.Any("Title", req.Title),
		zap.Any("desc", req.Desc))
	jsonResp, _ := CreateFileRespGen(id)
	fmt.Fprintf(w, jsonResp+"\n")

}

func (pkg *EdaPkg) InsertComponent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	var component common.Component
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	err := bson.UnmarshalExtJSON(body, false, &component)
	if err != nil {
		Logger.Error("insert component bson.Unmarshal error",
			zap.String("body", string(body)), zap.Error(err))
		jsonResp, _ := RespGen(common.K_REQUEST_COMMAND_ERROR, "Invalid json format")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
	Logger.Info("InsertComponent", zap.Any("Component", component))
	pkg.db.InsertComponent(component)
	jsonResp, _ := RespGen(0, "Success")
	fmt.Fprintf(w, jsonResp+"\n")

}

func (pkg *EdaPkg) InsertLine(w http.ResponseWriter, r *http.Request) {
	var line common.Line
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	err := bson.UnmarshalExtJSON(body, false, &line)
	if err != nil {
		Logger.Error("insert line bson.Unmarshal error", zap.Error(err))
		jsonResp, _ := RespGen(common.K_REQUEST_COMMAND_ERROR, "Invalid json format")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
	Logger.Info("InsertLine", zap.Any("Line", line))
	pkg.db.InsertLine(line)
	jsonResp, _ := RespGen(0, "Success")
	fmt.Fprintf(w, jsonResp+"\n")
}

func (pkg *EdaPkg) UpdateComponent(w http.ResponseWriter, r *http.Request) {
	components := make([]common.Component, 2)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	err := bson.UnmarshalExtJSON(body, false, &components)
	if err != nil {
		Logger.Error("update component bson.Unmarshal error", zap.Error(err))
		jsonResp, _ := RespGen(common.K_REQUEST_COMMAND_ERROR, "Invalid json format")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
	Logger.Info("UpdateComponent", zap.Any("oldComponent", components[0]),
		zap.Any("newComponent", components[1]))
	pkg.db.UpdateComponent(components[0], components[1])
	jsonResp, _ := RespGen(0, "Success")
	fmt.Fprintf(w, jsonResp+"\n")
}

func (pkg *EdaPkg) UpdateLine(w http.ResponseWriter, r *http.Request) {
	lines := make([]common.Line, 2)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	err := bson.UnmarshalExtJSON(body, false, &lines)
	if err != nil {
		Logger.Error("update line bson.Unmarshal error", zap.Error(err))
		jsonResp, _ := RespGen(common.K_REQUEST_COMMAND_ERROR, "Invalid json format")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
	Logger.Info("UpdateLine", zap.Any("oldLine", lines[0]),
		zap.Any("newLine", lines[1]))
	pkg.db.UpdateLine(lines[0], lines[1])
	jsonResp, _ := RespGen(0, "Success")
	fmt.Fprintf(w, jsonResp+"\n")
}

func (pkg *EdaPkg) DeleteComponent(w http.ResponseWriter, r *http.Request) {
	var component common.Component
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	err := bson.UnmarshalExtJSON(body, false, &component)
	if err != nil {
		Logger.Error("delete component bson.Unmarshal error", zap.Error(err))
		jsonResp, _ := RespGen(common.K_REQUEST_COMMAND_ERROR, "Invalid json format")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
	Logger.Info("DeleteComponent", zap.Any("Component", component))
	pkg.db.DeleteComponent(component)
	jsonResp, _ := RespGen(0, "Success")
	fmt.Fprintf(w, jsonResp+"\n")
}

func (pkg *EdaPkg) DeleteLine(w http.ResponseWriter, r *http.Request) {
	var line common.Line
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	err := bson.UnmarshalExtJSON(body, false, &line)
	if err != nil {
		Logger.Error("delete line bson.Unmarshal error", zap.Error(err))
		jsonResp, _ := RespGen(common.K_REQUEST_COMMAND_ERROR, "Invalid json format")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
	Logger.Info("DeleteLine", zap.Any("Line", line))
	pkg.db.DeleteLine(line)
	jsonResp, _ := RespGen(0, "Success")
	fmt.Fprintf(w, jsonResp+"\n")
}

func (pkg *EdaPkg) File(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		pkg.GetFile(w, req)
	} else if req.Method == http.MethodPost {
		pkg.CreateFile(w, req)
	} else {
		jsonResp, _ := RespGen(common.K_REQUEST_COMMAND_ERROR, "Not supported request method")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
}

func (pkg *EdaPkg) Line(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		pkg.InsertLine(w, req)
	} else if req.Method == http.MethodDelete {
		pkg.DeleteLine(w, req)
	} else if req.Method == http.MethodPut {
		pkg.UpdateLine(w, req)
	} else {
		jsonResp, _ := RespGen(common.K_REQUEST_COMMAND_ERROR, "Not supported request method")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
}

func (pkg *EdaPkg) Component(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		pkg.InsertComponent(w, req)
	} else if req.Method == http.MethodDelete {
		pkg.DeleteComponent(w, req)
	} else if req.Method == http.MethodPut {
		pkg.UpdateComponent(w, req)
	} else {
		jsonResp, _ := RespGen(common.K_REQUEST_COMMAND_ERROR, "Not supported request method")
		fmt.Fprintf(w, jsonResp+"\n")
		return
	}
}
