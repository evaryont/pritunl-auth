package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/pritunl/pritunl-auth/constants"
	"github.com/pritunl/pritunl-auth/google"
	"github.com/pritunl/pritunl-auth/handlers"
	"github.com/pritunl/pritunl-auth/utils"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Starts web server node
func App() {
	opts := getOpts()

	google.Init(opts.GoogleId, opts.GoogleSecret, opts.GoogleCallback)

	constants.Key = utils.GetKey(os.Getenv("LICENSE_KEY"))
	constants.HashKey = utils.GetKey(os.Getenv("LICENSE_HASH_KEY"))

	var debug bool
	debugStr := os.Getenv("DEBUG")
	if debugStr == "" {
		debug = true
	} else {
		debug, _ = strconv.ParseBool(debugStr)
	}

	router := gin.New()

	if debug {
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
