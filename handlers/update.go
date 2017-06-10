package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/evaryont/pritunl-auth/database"
	"github.com/evaryont/pritunl-auth/google"
	"github.com/evaryont/pritunl-auth/utils"
)

func updateGoogleGet(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)

	params := utils.ParseParams(c.Request)

	usr := params.GetByName("user")
	license := params.GetByName("license")

	err := google.Update(db, usr)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.String(200, "")
}
