package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/dropbox/godropbox/errors"
	"github.com/pritunl/pritunl-auth/constants"
	"strings"
)

func GetKey(phrase string) (key []byte) {
	digest := sha256.Sum256([]byte(phrase))

	for i := 0; i < 5; i++ {
		newDigest := sha256.Sum256(digest[:])
		digest = newDigest
	}

	key = digest[:]

	return
}

func DecrpytLicense(license string) (licenseHashHex string, err error) {
	license = strings.Replace(license,
		"-----------BEGIN LICENSE-----------", "", 1)
	license = strings.Replace(license,
		"------------END LICENSE------------", "", 1)
	license = strings.Replace(license, " ", "", -1)
	license = strings.Replace(license, "\n", "", -1)

	input, err := hex.DecodeString(license)
	if err != nil {
		err = &ParseError{
			errors.Wrap(err, "utils.license: Hex parse error"),
		}
		return
	}

	n := len(input)
	if n < 32 || n%16 != 0 {
		err = &InvalidLicenseError{
			errors.New("utils.license: Invalid license length"),
		}
		return
	}

	ivHash := sha1.Sum(input[:16])
	iv := ivHash[:16]

	block, err := aes.NewCipher(constants.Key)
	if err != nil {
		err = &ParseError{
			errors.Wrap(err, "utils.license: Aes key error"),
		}
		return
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	input = input[16:]

	output := make([]byte, len(input))
	mode.CryptBlocks(output, input)

	data := strings.Split(strings.Replace(string(output), "\x00", "", -1), "&")
	if len(data) < 3 {
		err = &InvalidLicenseError{
			errors.New("utils.license: Invalid license"),
		}
		return
	}

	licenseKey, err := base64.StdEncoding.DecodeString(data[1])
	if err != nil {
		err = &InvalidLicenseError{
			errors.Wrap(err, "utils.license: Invalid license"),
		}
		return
	}

	hashFunc := hmac.New(sha256.New, constants.HashKey)
	hashFunc.Write(licenseKey)
	licenseHash := hashFunc.Sum(nil)
	licenseHashHex = hex.EncodeToString(licenseHash)

	return
}
