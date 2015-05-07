// Oauth account.
package account

import (
	"time"
)

type Account struct {
	Id            string    `bson:"_id"`
	Type          string    `bson:"type"`
	Oauth2AccTokn string    `json:"oauth2_acct_tokn"`
	Oauth2RefTokn string    `json:"oauth2_ref_tokn"`
	Oauth2Exp     time.Time `json:"oauth2_exp"`
}
