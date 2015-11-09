package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-auth/database"
	"github.com/pritunl/pritunl-auth/google"
	"github.com/pritunl/pritunl-auth/user"
	"github.com/pritunl/pritunl-auth/utils"
)

func updateGoogleGet(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)

	params := utils.ParseParams(c.Request)

	usr := params.GetByName("user")
	license := params.GetByName("license")

	valid, err := user.CheckLicense(db, license)
	if err != nil {
		switch err.(type) {
		case *database.NotFoundError:
			c.AbortWithError(404, err)
		default:
			c.AbortWithError(500, err)
		}
		return
	}

	if !valid {
		c.AbortWithError(401, err)
		return
	}

	err = google.Update(db, usr)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.String(200, "")
}
