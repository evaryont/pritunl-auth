package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-auth/database"
	"github.com/pritunl/pritunl-auth/google"
	"github.com/pritunl/pritunl-auth/utils"
)

func callbackGoogleGet(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)

	params := utils.ParseParams(c.Request)

	state := params.GetByName("state")
	code := params.GetByName("code")
	error := params.GetByName("error")

	switch error {
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
		c.Fail(500, err)
		return
	}

	c.Redirect(301, fmt.Sprintf("%s?state=%s&user=%s", tokn.RemoteCallback,
		tokn.RemoteState, acct.Id))
}
