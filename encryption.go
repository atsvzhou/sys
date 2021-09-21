package sys

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

//Sha256 加密
func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// HmacSha256 加密
func HmacSha256(stringToSign string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(stringToSign))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}
