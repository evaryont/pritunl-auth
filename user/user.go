package user

import (
	"github.com/pritunl/pritunl-auth/database"
	"github.com/pritunl/pritunl-auth/utils"
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

func CheckLicense(db *database.Database, license string) (
	valid bool, err error) {

	id, licenseHash, err := utils.DecrpytLicense(license)
	if err != nil {
		return
	}

	usr, err := FindUser(db, id)
	if err != nil {
		return
	}

	if usr.LicenseHash != licenseHash ||
		usr.Plan[:len(usr.Plan)-1] != "enterprise" {

		return
	}

	valid = true

	return
}
