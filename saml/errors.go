package saml

import (
	"github.com/dropbox/godropbox/errors"
)

type SamlError struct {
	errors.DropboxError
}
