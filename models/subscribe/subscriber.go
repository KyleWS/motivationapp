package subscribe

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Subscriber struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Email     string        `json:"email,omitempty"`
	Mode      string        `json:"mode,omitempty"`
	CreatedAt time.Time     `json:"createdAt,omitempty"`
}

type NewSubscriber struct {
	Email string `json:"email,omitempty"`
	Mode  string `json:"mode,omitempty"`
}

func (ns *NewSubscriber) ToSubscriber() *Subscriber {
	sub := &Subscriber{
		ID:        bson.NewObjectId(),
		Email:     ns.Email,
		Mode:      ns.Mode,
		CreatedAt: time.Now(),
	}
	return sub
}
