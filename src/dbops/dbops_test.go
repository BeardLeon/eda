package dbops

import (
	"eda/src/common"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

var dbops DBOps

var dataSourceName string = fmt.Sprintf("%s://%s:%s", "mongodb",
	"localhost", "27017")

var coll string = "final"

func TestInit(t *testing.T) {
	dbops.Init(dataSourceName, "eda", coll)
}

func TestInsertLine(t *testing.T) {
	line := common.Line{
		StartX: 2,
		StartY: 2,
		EndX:   2,
		EndY:   2,
	}
	dbops.InsertLine(line)
}

func TestDeleteLine(t *testing.T) {
	line := common.Line{
		StartX: 2,
		StartY: 2,
		EndX:   2,
		EndY:   2,
	}
	dbops.DeleteLine(line)
}

func TestUpdateLine(t *testing.T) {
	curLine := common.Line{
		StartX: 3,
		StartY: 3,
		EndX:   3,
		EndY:   3,
	}
	preLine := common.Line{
		StartX: 2,
		StartY: 2,
		EndX:   2,
		EndY:   2,
	}
	dbops.UpdateLine(preLine, curLine)
}

func TestInsertComponent(t *testing.T) {
	component := common.Component{
		Id:   8,
		Name: "Good test name",
		Shape: []common.Line{
			{StartX: 2,
				StartY: 2,
				EndX:   2,
				EndY:   2,
			},
			{StartX: 3,
				StartY: 3,
				EndX:   3,
				EndY:   3},
		},
		Pin: []common.Coordinate{
			{X: 0, Y: 0},
			{X: 3, Y: 3},
		},
	}
	dbops.InsertComponent(component)
}

func TestDeleteComponent(t *testing.T) {
	component := common.Component{
		Id:   8,
		Name: "Good test name",
		Shape: []common.Line{
			{StartX: 2,
				StartY: 2,
				EndX:   2,
				EndY:   2,
			},
			{StartX: 3,
				StartY: 3,
				EndX:   3,
				EndY:   3},
		},
		Pin: []common.Coordinate{
			{X: 0, Y: 0},
			{X: 3, Y: 3},
		},
	}
	dbops.DeleteComponent(component)
}

func TestUpdateComponent(t *testing.T) {
	preComponent := common.Component{
		Id:   8,
		Name: "Good test name",
		Shape: []common.Line{
			{StartX: 2,
				StartY: 2,
				EndX:   2,
				EndY:   2,
			},
			{StartX: 3,
				StartY: 3,
				EndX:   3,
				EndY:   3},
		},
		Pin: []common.Coordinate{
			{X: 0, Y: 0},
			{X: 4, Y: 4},
		},
	}
	curComponent := common.Component{
		Id:   8,
		Name: "Good test name",
		Shape: []common.Line{
			{StartX: 2,
				StartY: 2,
				EndX:   2,
				EndY:   2,
			},
			{StartX: 3,
				StartY: 3,
				EndX:   3,
				EndY:   3},
		},
		Pin: []common.Coordinate{
			{X: 5, Y: 5},
			{X: 4, Y: 4},
		},
	}
	dbops.UpdateComponent(preComponent, curComponent)
}

func TestCreateFile(t *testing.T) {
	dbops.CreateFile("soso title", "this is a soso desc")
}

func TestDBOps(t *testing.T) {
	dbops.Init(dataSourceName, "eda", coll)
	dbops.CreateFile("soso title", "this is a soso desc")
	component := common.Component{
		Id:   8,
		Name: "Good test name",
		Shape: []common.Line{
			{StartX: 2,
				StartY: 2,
				EndX:   2,
				EndY:   2,
			},
			{StartX: 3,
				StartY: 3,
				EndX:   3,
				EndY:   3},
		},
		Pin: []common.Coordinate{
			{X: 0, Y: 0},
			{X: 3, Y: 3},
		},
	}
	dbops.InsertComponent(component)

	line := common.Line{
		StartX: 2,
		StartY: 2,
		EndX:   2,
		EndY:   2,
	}
	dbops.InsertLine(line)
}

func TestMarshal(t *testing.T) {
	line := common.Line{
		StartX: 2,
		StartY: 2,
		EndX:   2,
		EndY:   2,
	}
	res, _ := bson.MarshalExtJSON(line, false, false)
	fmt.Println(string(res))
}
