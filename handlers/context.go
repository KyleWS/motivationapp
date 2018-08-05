package handlers

import (
	"github.com/KyleWS/motivationapp/models/subscribe"
)

//Context holds context values

type Context struct {
	subscriberStore subscribe.MongoStore
}

func NewHandlerContext(subscriberStore subscribe.MongoStore) *Context {
	return &Context{
		subscriberStore: subscriberStore,
	}
}
