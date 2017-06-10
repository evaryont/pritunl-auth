package handlers

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/dropbox/godropbox/errors"
	"github.com/gin-gonic/gin"
	"github.com/evaryont/pritunl-auth/database"
	"net/http"
)

// Limit size of request body
func Limiter(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1000000)
}

// Get database from session
func Database(c *gin.Context) {
	db := database.GetDatabase()
	c.Set("db", db)
	c.Next()
	db.Close()
}

// Recover panics
func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.WithFields(logrus.Fields{
				"client": c.Request.RemoteAddr,
				"error":  errors.New(fmt.Sprintf("%s", r)),
			}).Error("handlers: Handler panic")
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
	}()

	c.Next()
}

// Register all endpoint handlers
func Register(engine *gin.Engine) {
	engine.Use(Limiter)
	engine.Use(Recovery)

	dbGroup := engine.Group("")
	dbGroup.Use(Database)

	dbGroup.GET("/check", checkGet)
	dbGroup.POST("/request/google", requestGooglePost)
	dbGroup.POST("/request/saml", requestSamlPost)
	dbGroup.GET("/callback/google", callbackGoogleGet)
	dbGroup.POST("/callback/saml", callbackGoogleGet)
	dbGroup.GET("/update/google", updateGoogleGet)

	dbGroup.POST("/v1/request/google", requestGoogle2Post)
	dbGroup.POST("/v1/request/saml", requestSamlPost)
	dbGroup.GET("/v1/callback/google", callbackGoogleGet)
	dbGroup.POST("/v1/callback/saml", callbackSamlPost)
	dbGroup.GET("/v1/update/google", updateGoogleGet)
}
