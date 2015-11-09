package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-auth/database"
	"github.com/pritunl/pritunl-auth/google"
	"github.com/pritunl/pritunl-auth/saml"
	"github.com/pritunl/pritunl-auth/utils"
	"net/url"
)

func callbackGoogleGet(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)

	params := utils.ParseParams(c.Request)

	state := params.GetByName("state")
	code := params.GetByName("code")
	authErr := params.GetByName("error")

	switch authErr {
	case "":
		if state == "" || code == "" {
			c.AbortWithStatus(400)
			return
		}
	case "access_denied":
		// TODO Redirect to base callback url
		c.Redirect(301, "https://pritunl.com/")
		return
	default:
		c.AbortWithStatus(400)
		return
	}

	acct, tokn, err := google.Authorize(db, state, code)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	if tokn.Version == 1 {
		query := fmt.Sprintf("state=%s&username=%s", tokn.RemoteState,
			url.QueryEscape(acct.Id))

		hashFunc := hmac.New(sha512.New, []byte(tokn.RemoteSecret))
		hashFunc.Write([]byte(query))
		rawSignature := hashFunc.Sum(nil)
		sig := base64.URLEncoding.EncodeToString(rawSignature)

		url := fmt.Sprintf("%s?%s&sig=%s",
			tokn.RemoteCallback, query, url.QueryEscape(sig))

		c.Redirect(301, url)
	} else {
		hashFunc := hmac.New(sha256.New, []byte(tokn.RemoteSecret))
		hashFunc.Write([]byte(tokn.RemoteState + acct.Id))
		rawSignature := hashFunc.Sum(nil)
		sig := base64.URLEncoding.EncodeToString(rawSignature)

		c.Redirect(301, fmt.Sprintf("%s?state=%s&user=%s&sig=%s",
			tokn.RemoteCallback, tokn.RemoteState,
			url.QueryEscape(acct.Id), sig))
	}
}

func callbackSamlPost(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)

	state := c.PostForm("RelayState")
	respEncoded := c.PostForm("SAMLResponse")

	data, tokn, err := saml.Authorize(db, state, respEncoded)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	query := fmt.Sprintf("state=%s&username=%s&email=%s&org=%s&secondary=%s",
		tokn.RemoteState,
		url.QueryEscape(data.Username),
		url.QueryEscape(data.Email),
		url.QueryEscape(data.Org),
		url.QueryEscape(data.Secondary),
	)

	hashFunc := hmac.New(sha512.New, []byte(tokn.RemoteSecret))
	hashFunc.Write([]byte(query))
	rawSignature := hashFunc.Sum(nil)
	sig := base64.URLEncoding.EncodeToString(rawSignature)

	c.Redirect(301, fmt.Sprintf("%s?%s&sig=%s",
		tokn.RemoteCallback, query, sig))
}
