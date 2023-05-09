package edaPkg

import (
	"eda/src/zaplog"
	"fmt"
	"net/http"
)

type EdaPkg struct {
	Id          string      `bson:"_id"`
	Title       string      `bson:"title"`
	Description string      `bson:"description"`
	Components  []Component `bson:"components"`
	Lines       []Line      `bson:"lines"`
}

type Component struct {
	Id    int          `bson:"id"`
	Name  string       `bson:"name"`
	Shape []Line       `bson:"shape"`
	Pin   []Coordinate `bson:"pin"`
}

type Line struct {
	StartX int `bson:"sx"`
	StartY int `bson:"sy"`
	EndX   int `bson:"ex"`
	EndY   int `bson:"ey"`
}
type Coordinate struct {
	X int `bson:"x"`
	Y int `bson:"y"`
}

var Logger = zaplog.Logger

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

}

func (pkg *EdaPkg) InsertLine(w http.ResponseWriter, r *http.Request) {

}

func (pkg *EdaPkg) UpdateComponent(w http.ResponseWriter, r *http.Request) {

}

func (pkg *EdaPkg) UpdateLine(w http.ResponseWriter, r *http.Request) {

}

func (pkg *EdaPkg) DeleteComponent(w http.ResponseWriter, r *http.Request) {

}

func (pkg *EdaPkg) DeleteLine(w http.ResponseWriter, r *http.Request) {

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
