// Global constants.
package constants

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

const (
	RetryDelay       = 3 * time.Second
	ErrLogDelay      = 3 * time.Minute
	KeepAliveTimeout = 2 * time.Minute
	SamlReq          = "saml.req"
	SamlKey          = "saml.key"
	SamlCert         = "saml.crt"
)

var (
	SamlCertDir string
	Version     string
	Key         []byte
	HashKey     []byte
)

func init() {
	ver := "unknown"

	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		return
	}

	pkgPath := path.Join(goPath, "src/github.com/evaryont/pritunl-auth")

	output, err := exec.Command("git", "-C", pkgPath,
		"rev-parse", "HEAD").Output()
	if err != nil {
		return
	}

	ver = strings.TrimSpace(string(output))
	if len(ver) > 8 {
		ver = ver[:8]
	}

	Version = ver
}
