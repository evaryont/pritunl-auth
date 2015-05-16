// Oauth account.
package account

import (
	"time"
)

type Account struct {
	Id            string    `bson:"_id"`
	Type          string    `bson:"type"`
	Oauth2AccTokn string    `bson:"oauth2_acc_tokn"`
	Oauth2RefTokn string    `bson:"oauth2_ref_tokn"`
	Oauth2Exp     time.Time `bson:"oauth2_exp"`
}
