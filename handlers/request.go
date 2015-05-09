package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-auth/database"
	"github.com/pritunl/pritunl-auth/google"
	"github.com/pritunl/pritunl-auth/user"
	"github.com/pritunl/pritunl-auth/utils"
)

type requestData struct {
	License  string `json:"license"`
	Callback string `json:"callback"`
	State    string `state:"state"`
	Secret   string `state:"secret"`
}

func requestPost(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)
	data := &requestData{}

	if !c.Bind(&data) {
		return
	}

	id, licenseHash, err := utils.DecrpytLicense(data.License)
	if err != nil {
		c.Fail(500, err)
		return
	}

	usr, err := user.FindUser(db, id)
	if err != nil {
		switch err.(type) {
		case *database.NotFoundError:
			c.Fail(404, err)
		default:
			c.Fail(500, err)
		}
		return
	}

	if usr.LicenseHash != licenseHash {
		c.Fail(401, err)
		return
	}

	if usr.Plan[:len(usr.Plan)-1] != "enterprise" {
		c.Fail(401, err)
		return
	}

	url, err := google.Request(db, data.State, data.Secret, data.Callback)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(200, struct {
		Url string `json:"url"`
	}{
		Url: url,
	})
}
