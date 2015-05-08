package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-auth/database"
	"github.com/pritunl/pritunl-auth/google"
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
