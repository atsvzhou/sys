package sys

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
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

func NewMD5(str string) string {
	data := []byte(str)
	h := md5.New()
	h.Write(data)
	md5String := hex.EncodeToString(h.Sum(nil))
	return md5String
}

// GetFileMd5 获取文件的md5码
func GetFileMd5(filename string) (string, error) {
	// 文件全路径名
	pFile, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("打开文件失败，filename=%v, err=%v", filename, err)
		return "", err
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)

	return hex.EncodeToString(md5h.Sum(nil)), err
}
