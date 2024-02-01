package models

// PingPong represents the ping pong model.
type PingPong struct {
	// ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Text string `bson:"text" json:"text"`
}

// Error implements error.
// func (*PingPong) Error() string {
// 	panic("unimplemented")
// }
