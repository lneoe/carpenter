package service

import (
	"crypto/sha256"
	"encoding/hex"
)

func PasswordSHA224(rawPwd string) string {
	h := sha256.New224()
	_, _ = h.Write([]byte(rawPwd))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
