// Oauth1 and oauth2 clients.
package oauth

type Token struct {
	Id          string `bson:"_id"`
	RemoteState string `json:"remote_state"`
	OauthSecret string `bson:"oauth_secret"`
	Type        string `bson:"type"`
}
