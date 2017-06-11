package saml

import (
	"fmt"
	"github.com/evaryont/pritunl-auth/constants"
	"github.com/evaryont/pritunl-auth/utils"
	"os"
	"path/filepath"
)

func InitSignCert() (err error) {
	fullPathKey  := filepath.Join(constants.SamlCertDir, constants.SamlKey)
	fullPathReq  := filepath.Join(constants.SamlCertDir, constants.SamlReq)
	fullPathCert := filepath.Join(constants.SamlCertDir, constants.SamlCert)

	if _, err := os.Stat(fullPathCert); err == nil {
		fmt.Printf("SAML certificate %s exists already\n", fullPathCert)
		return nil
	}

	fmt.Printf("Generating new SAML certificate %s\n", fullPathCert)
	os.Mkdir(constants.SamlCertDir, 0700)

	err = utils.Exec(constants.SamlCertDir, "openssl", "req",
		"-days", "3650",
		"-newkey", "rsa:4096",
		"-nodes",
		"-subj", "/C=US/ST=New York/O=Pritunl/CN=pritunl.com",
		"-keyout", fullPathKey,
		"-out", fullPathReq,
	)
	if err != nil {
		return
	}

	err = utils.Exec(constants.SamlCertDir, "openssl", "x509",
		"-req",
		"-in", fullPathReq,
		"-signkey", fullPathKey,
		"-sha256",
		"-days", "3650",
		"-out", fullPathCert,
	)
	if err != nil {
		return
	}

	return
}

func GetCertPath() string {
	return filepath.Join(constants.SamlCertDir,
		fmt.Sprintf("%s.crt", utils.Uuid()))
}
