package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/evaryont/pritunl-auth/google"
	"github.com/evaryont/pritunl-auth/handlers"
	"github.com/evaryont/pritunl-auth/saml"
	"net/http"
	"strconv"
	"time"
)

// Starts web server node
func App() {
	opts := getOpts()

	google.Init(opts.GoogleId, opts.GoogleSecret, opts.GoogleCallback)

	saml.SamlCallbackUrl = opts.SamlCallback

	router := gin.New()

	if opts.Debug {
		router.Use(gin.Logger())
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	handlers.Register(router)

	server := http.Server{
		Addr:           ":" + strconv.Itoa(opts.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 4096,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
