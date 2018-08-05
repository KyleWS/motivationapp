package subscribe

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

//MongoStore implements Store for MongoDB
type MongoStore struct {
	//the mongo session
	session *mgo.Session
	//the database name to use
	dbname string
	//the collection name to use
	colname string
	//the Collection object for that dbname/colname
	collection *mgo.Collection
}

//NewMongoStore constructs a new MongoStore, given a live mgo.Session,
//a database name, and a collection name
func NewMongoStore(sess *mgo.Session, dbName string, collectionName string) *MongoStore {
	if sess == nil {
		// panic to see stacktrace
		panic("nil pointer passed for session")
	}

	//return a new MongoStore
	return &MongoStore{
		session:    sess,
		dbname:     dbName,
		colname:    collectionName,
		collection: sess.DB(dbName).C(collectionName),
	}
}

func (s *MongoStore) Insert(ns *NewSubscriber) (*Subscriber, error) {
	subscriber := ns.ToSubscriber()
	if err := s.collection.Insert(subscriber); err != nil {
		return nil, fmt.Errorf("error inserting subscriber: %v", err)
	}
	return subscriber, nil
}

func (s *MongoStore) GetAll() ([]*Subscriber, error) {
	subs := []*Subscriber{}
	if err := s.collection.Find(&Subscriber{}).Limit(100).All(&subs); err != nil {
		return nil, fmt.Errorf("error getting tasks: %v", err)
	}
	return subs, nil
}
