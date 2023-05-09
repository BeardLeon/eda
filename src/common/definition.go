package common

type Component struct {
	Id    int          `bson:"id"`
	Name  string       `bson:"name"`
	Shape []Line       `bson:"shape"`
	Pin   []Coordinate `bson:"pin"`
}

type Line struct {
	StartX int `bson:"sx"`
	StartY int `bson:"sy"`
	EndX   int `bson:"tx"`
	EndY   int `bson:"ty"`
}
type Coordinate struct {
	X int `bson:"x"`
	Y int `bson:"y"`
}
