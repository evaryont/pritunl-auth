package saml

import (
	"fmt"
	"github.com/pritunl/pritunl-auth/constants"
	"github.com/pritunl/pritunl-auth/utils"
	"os"
	"path/filepath"
)

func InitSignCert() (err error) {
	constants.SamlCertDir = filepath.Join("/", "tmp",
		fmt.Sprintf("pritunl_%s", utils.Uuid()))

	os.Mkdir(constants.SamlCertDir, 0700)

	err = utils.Exec(constants.SamlCertDir, "openssl", "req",
		"-days", "3650",
		"-newkey", "rsa:4096",
		"-nodes",
		"-subj", "/C=US/ST=New York/O=Pritunl/CN=pritunl.com",
		"-keyout", "saml.key",
		"-out", "saml.req",
	)
	if err != nil {
		return
	}

	err = utils.Exec(constants.SamlCertDir, "openssl", "x509",
		"-req",
		"-in", "saml.req",
		"-signkey", "saml.key",
		"-sha256",
		"-days", "3650",
		"-out", "saml.crt",
	)
	if err != nil {
		return
	}

	return
}
