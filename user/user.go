package user

import (
	"github.com/evaryont/pritunl-auth/database"
	"github.com/evaryont/pritunl-auth/utils"
)

type User struct {
	Id          string `bson:"_id"`
	Status      string `bson:"status"`
	Plan        string `bson:"plan"`
	LicenseHash string `bson:"license_hash"`
}

func FindUser(db *database.Database, id string) (usr *User, err error) {
	usr = &User{}
	coll := db.Users()

	err = coll.FindOneId(id, usr)
	if err != nil {
		return
	}

	return
}
