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
		"-keyout", constants.SamlKey,
		"-out", constants.SamlReq,
	)
	if err != nil {
		return
	}

	err = utils.Exec(constants.SamlCertDir, "openssl", "x509",
		"-req",
		"-in", constants.SamlReq,
		"-signkey", constants.SamlKey,
		"-sha256",
		"-days", "3650",
		"-out", constants.SamlCert,
	)
	if err != nil {
		return
	}

	return
}
