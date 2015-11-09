// Oauth1 and oauth2 clients.
package oauth

type Token struct {
	Id             string `bson:"_id"`
	RemoteCallback string `bson:"remote_callback"`
	RemoteState    string `bson:"remote_state"`
	RemoteSecret   string `bson:"remote_secret"`
	OauthSecret    string `bson:"oauth_secret"`
	Type           string `bson:"type"`
	Version        int    `bson:"version"`
}
