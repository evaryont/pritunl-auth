package database

import (
	"github.com/Sirupsen/logrus"
	"github.com/dropbox/godropbox/errors"
	"github.com/pritunl/pritunl-auth/constants"
	"github.com/pritunl/pritunl-auth/requires"
	"labix.org/v2/mgo"
	"os"
	"strings"
	"time"
)

var (
	MongoUrl string
	Session  *mgo.Session
)

type Database struct {
	session  *mgo.Session
	database *mgo.Database
}

func (d *Database) Close() {
	d.session.Close()
}

func (d *Database) getCollection(name string) (coll *Collection) {
	coll = &Collection{
		*d.database.C(name),
		d,
	}
	return
}

func (d *Database) Users() (coll *Collection) {
	coll = d.getCollection("users")
	return
}

func (d *Database) AuthUsers() (coll *Collection) {
	coll = d.getCollection("auth_users")
	return
}

func Connect() (err error) {
	Session, err = mgo.Dial(MongoUrl)
	if err != nil {
		err = &ConnectionError{
			errors.Wrap(err, "database: Connection error"),
		}
		return
	}

	Session.SetMode(mgo.Strong, true)

	return
}

func GetDatabase() (db *Database) {
	session := Session.Copy()

	var dbName string
	if x := strings.LastIndex(MongoUrl, "/"); x != -1 {
		dbName = MongoUrl[x+1:]
	} else {
		dbName = "pritunl"
	}

	database := session.DB(dbName)

	db = &Database{
		session:  session,
		database: database,
	}
	return
}

func addIndexes() (err error) {
	db := GetDatabase()
	defer db.Close()

	coll := db.AuthUsers()
	err = coll.EnsureIndex(mgo.Index{
		Key:        []string{"username"},
		Background: true,
	})
	if err != nil {
		err = &IndexError{
			errors.Wrap(err, "database: Index error"),
		}
	}

	return
}

func init() {
	MongoUrl = os.Getenv("DB")

	module := requires.New("database")

	module.Handler = func() {
		for {
			err := Connect()
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("database: Connection")
			} else {
				break
			}

			time.Sleep(constants.RetryDelay)
		}

		for {
			err := addIndexes()
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("database: Add indexes")
			} else {
				break
			}

			time.Sleep(constants.RetryDelay)
		}
	}
}
