package common

type Component struct {
	OId   string       `bson:"_id"`
	Id    int          `bson:"id"`
	Name  string       `bson:"name"`
	Shape []Line       `bson:"shape,omitempty"`
	Pin   []Coordinate `bson:"pin"`
}

type Line struct {
	OId    string `bson:"_id,omitempty"`
	StartX int    `bson:"sx"`
	StartY int    `bson:"sy"`
	EndX   int    `bson:"ex"`
	EndY   int    `bson:"ey"`
}
type Coordinate struct {
	X int `bson:"x"`
	Y int `bson:"y"`
}
