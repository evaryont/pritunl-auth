// Oauth1 and oauth2 clients.
package oauth

type Token struct {
	Id             string `bson:"_id"`
	RemoteCallback string `json:"remote_callback"`
	RemoteState    string `json:"remote_state"`
	RemoteSecret   string `json:"remote_secret"`
	OauthSecret    string `bson:"oauth_secret"`
	Type           string `bson:"type"`
}
