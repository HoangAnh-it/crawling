package utils

import (
	"crypto"
	"encoding/hex"
	"regexp"
	"strings"
)

func IsMd5Hash(md5 string) bool {
	if len(md5) != 32 || !isHex(md5) {
		return false
	}
	return true
}

func IsSha1Hash(md5 string) bool {
	if len(md5) != 40 || !isHex(md5) {
		return false
	}
	return true
}

func IsSh256Hash(md5 string) bool {
	if len(md5) != 64 || !isHex(md5) {
		return false
	}
	return true
}

func isHex(s string) bool {
	hexPattern := regexp.MustCompile("^[0-9a-fA-F]+$")
	return hexPattern.MatchString(strings.TrimSpace(s))
}

func HashMd5(s string) string {
	md5 := crypto.MD5.New()
	md5.Write([]byte(s))
	return hex.EncodeToString(md5.Sum(nil))
}

func HashSha1(s string) string {
	sha1 := crypto.SHA1.New()
	sha1.Write([]byte(s))
	return hex.EncodeToString(sha1.Sum(nil))
}

func HashSha256(s string) string {
	sha256 := crypto.SHA256.New()
	sha256.Write([]byte(s))
	return hex.EncodeToString(sha256.Sum(nil))
}
