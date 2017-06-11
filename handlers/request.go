package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/evaryont/pritunl-auth/database"
	"github.com/evaryont/pritunl-auth/google"
	"github.com/evaryont/pritunl-auth/saml"
	"github.com/Sirupsen/logrus"
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

  logrus.WithFields(logrus.Fields{
  	"callback": data.Callback,
		"issuer_url": data.IssuerUrl,
  	"license": data.License,
		"secret": data.Secret,
		"sso_url": data.SsoUrl,
		"state": data.State,
  }).Info("request saml called")

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
