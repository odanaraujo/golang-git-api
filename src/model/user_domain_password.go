package model

import (
	"crypto/md5"
	"encoding/hex"
)

func (domain *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(domain.password))
	domain.password = hex.EncodeToString(hash.Sum(nil))
}
