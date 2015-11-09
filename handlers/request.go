package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-auth/database"
	"github.com/pritunl/pritunl-auth/google"
	"github.com/pritunl/pritunl-auth/user"
)

type requestData struct {
	License  string `json:"license"`
	Callback string `json:"callback"`
	State    string `state:"state"`
	Secret   string `state:"secret"`
}

func requestGooglePost(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)
	data := &requestData{}

	err := c.Bind(&data)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	valid, err := user.CheckLicense(db, data.License)
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

	url, err := google.Request(db, data.State, data.Secret, data.Callback)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, struct {
		Url string `json:"url"`
	}{
		Url: url,
	})
}
