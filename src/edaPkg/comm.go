package edaPkg

import (
	"fmt"
	"golang-template/src/zaplog"
)

type EdaPkg struct {
	Id          string `bson:"_id"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Schematics  string `bson:"schematics"`
}

var Logger = zaplog.Logger

func (pkg *EdaPkg) New() {

}

// ImportJsonFile need name and file
func (pkg *EdaPkg) ImportJsonFile(file string) error {
	// TODO add string to db
	fmt.Println(file)
	return nil
}

// ExportJsonFile export name and file
func (pkg *EdaPkg) ExportJsonFile() (string, error) {
	return "", nil
}

func (eh *EdaPkg) UpdateJsonFile() error {
	return nil
}
