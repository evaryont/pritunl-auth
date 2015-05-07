package google

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pritunl/pritunl-auth/account"
	"github.com/pritunl/pritunl-auth/database"
	"github.com/pritunl/pritunl-auth/errortypes"
	"github.com/pritunl/pritunl-auth/oauth"
)

var (
	conf *oauth.Oauth2
)

func Init(clientId string, clientSecret string, callbackUrl string) {
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

func (g *GoogleClient) Init(db *database.Database) (err error) {
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

	g.acct.Id = data.Emails[1].Value

	return
}

func Request(db *database.Database, remoteState string) (
	url string, err error) {

	url, err = conf.Request(db, remoteState)
	if err != nil {
		return
	}

	return
}

func Authorize(db *database.Database, state string,
	code string) (acct *account.Account, err error) {

	coll := db.Accounts()

	auth, err := conf.Authorize(db, state, code)
	if err != nil {
		return
	}

	client := &GoogleClient{
		acct: auth.Account,
	}

	err = client.Init(db)
	if err != nil {
		return
	}

	err = coll.Insert(auth.Account)
	if err != nil {
		err = database.ParseError(err)
		return
	}

	return
}
