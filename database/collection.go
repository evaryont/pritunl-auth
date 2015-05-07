package database

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Collection struct {
	mgo.Collection
	Database *Database
}

func (c *Collection) FindOne(query interface{}, result interface{}) (
err error) {
	err = c.Find(query).One(result)
	if err != nil {
		err = ParseError(err)
		return
	}

	return
}

func (c *Collection) FindOneId(id interface{}, result interface{}) (
err error) {
	err = c.FindId(id).One(result)
	if err != nil {
		err = ParseError(err)
		return
	}

	return
}

func (c *Collection) Commit(id interface{}, data interface{}) (err error) {
	err = c.UpdateId(id, bson.M{
		"$set": data,
	})
	if err != nil {
		err = ParseError(err)
		return
	}

	return
}
