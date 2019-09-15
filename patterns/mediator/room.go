package mediator

//中介者模式
//房子

import "gopkg.in/mgo.v2/bson"

type Room struct {
	Id   bson.ObjectId
	Name string
}

func NewRoom(name string) *Room {
	return &Room{bson.NewObjectId(), name}
}
