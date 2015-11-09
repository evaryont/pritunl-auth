package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-auth/database"
	"github.com/pritunl/pritunl-auth/google"
	"github.com/pritunl/pritunl-auth/saml"
	"github.com/pritunl/pritunl-auth/user"
)

type googleRequestData struct {
	License  string `json:"license"`
	Callback string `json:"callback"`
	State    string `json:"state"`
	Secret   string `json:"secret"`
}

func _requestGooglePost(c *gin.Context, version int) {
	db := c.MustGet("db").(*database.Database)
	data := &googleRequestData{}

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

	url, err := google.Request(db, data.State, data.Secret, data.Callback,
		version)
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

func requestGooglePost(c *gin.Context) {
	_requestGooglePost(c, 0)
}

func requestGoogle2Post(c *gin.Context) {
	_requestGooglePost(c, 1)
}

type samlRequestData struct {
	License   string `json:"license"`
	Callback  string `json:"callback"`
	State     string `json:"state"`
	Secret    string `json:"secret"`
	SsoUrl    string `json:"sso_url"`
	IssuerUrl string `json:"issuer_url"`
	Cert      string `json:"cert"`
}

func requestSamlPost(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)
	data := &samlRequestData{}

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

	sml := &saml.Saml{
		SsoUrl: data.SsoUrl,
		IssuerUrl: data.IssuerUrl,
		Cert: data.Cert,
	}

	err = sml.Init()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}


	resp, err := sml.Request(db, data.State, data.Secret, data.Callback)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.Writer.WriteHeader(200)
	c.Writer.Header().Set("Content-Type", "text/html;charset=utf-8")
	c.Writer.Write(resp.Bytes())
}
