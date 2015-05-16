package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-auth/database"
	"github.com/pritunl/pritunl-auth/google"
	"github.com/pritunl/pritunl-auth/utils"
)

func updateGoogleGet(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)

	params := utils.ParseParams(c.Request)

	user := params.GetByName("user")

	err := google.Update(db, user)
	if err != nil {
		c.Fail(500, err)
	}

	c.String(200, "")
}
