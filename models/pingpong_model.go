package models

import "gopkg.in/mgo.v2/bson"

// PingPong represents the ping pong model.
type PingPong struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Text string        `bson:"text" json:"text"`
}
