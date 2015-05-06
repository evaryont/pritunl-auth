// Commands available in cli.
package cmd

import (
	"github.com/Sirupsen/logrus"
	"math/rand"
	"os"
	"strconv"
)

type Options struct {
	Debug        bool
	Port         int
	Database     string
	GoogleId     string
	GoogleSecret string
}

func getOpts() (opts *Options) {
	opts = &Options{}

	debugStr := os.Getenv("DEBUG")
	if debugStr == "" {
		opts.Debug = true
	} else {
		opts.Debug, _ = strconv.ParseBool(debugStr)
	}

	portStr := os.Getenv("PORT")
	if portStr != "" {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("cmd: Failed to parse port")
			panic(err)
		}
		opts.Port = p
	} else {
		opts.Port = rand.Intn(55000) + 10000
	}

	opts.Database = os.Getenv("DB")
	opts.GoogleId = os.Getenv("GOOGLE_ID")
	opts.GoogleSecret = os.Getenv("GOOGLE_SECRET")

	return
}
