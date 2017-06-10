// Google oauth client.
package google

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/evaryont/pritunl-auth/account"
	"github.com/evaryont/pritunl-auth/database"
	"github.com/evaryont/pritunl-auth/errortypes"
	"github.com/evaryont/pritunl-auth/oauth"
	"labix.org/v2/mgo/bson"
)

var (
	conf *oauth.Oauth2
)

func Init(clientId, clientSecret, callbackUrl string) {
	conf = &oauth.Oauth2{
		Type:         "gmail",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		CallbackUrl:  callbackUrl,
		AuthUrl:      "https://accounts.google.com/o/oauth2/auth",
		TokenUrl:     "https://www.googleapis.com/oauth2/v3/token",
		Scopes: []string{
			"profile",
			"email",
		},
	}
	conf.Config()
}

type GoogleClient struct {
	acct *account.Account
}

func (g *GoogleClient) Update(db *database.Database) (err error) {
	client := conf.NewClient(g.acct)

	err = client.Refresh(db)
	if err != nil {
		return
	}

	data := struct {
		Emails []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"emails"`
	}{}

	err = client.GetJson("https://www.googleapis.com/plus/v1/people/me", &data)
	if err != nil {
		return
	}

	if len(data.Emails) < 1 {
		err = errortypes.UnknownError{
			errors.New("Unable to get email from profile"),
		}
		return
	}

	g.acct.Id = data.Emails[0].Value

	return
}

func Request(db *database.Database, remoteState, remoteSecret,
	remoteCallback string, version int) (url string, err error) {

	url, err = conf.Request(db, remoteState, remoteSecret,
		remoteCallback, version)
	if err != nil {
		return
	}

	return
}

func Authorize(db *database.Database, state string,
	code string) (acct *account.Account, tokn *oauth.Token, err error) {

	coll := db.Accounts()

	acct, tokn, err = conf.Authorize(db, state, code)
	if err != nil {
		return
	}

	client := &GoogleClient{
		acct: acct,
	}

	err = client.Update(db)
	if err != nil {
		return
	}

	_, err = coll.Upsert(&bson.M{
		"_id": acct.Id,
	}, acct)
	if err != nil {
		err = database.ParseError(err)
		return
	}

	return
}

func Update(db *database.Database, user string) (err error) {
	coll := db.Accounts()

	acct := &account.Account{}

	err = coll.FindOneId(user, acct)
	if err != nil {
		return
	}

	client := &GoogleClient{
		acct: acct,
	}

	err = client.Update(db)
	if err != nil {
		return
	}

	_, err = coll.Upsert(&bson.M{
		"_id": acct.Id,
	}, acct)
	if err != nil {
		err = database.ParseError(err)
		return
	}

	return
}
