package auth

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"

	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Encrypt encrytps ths plain text with bcrypt
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func CreateUUID() (string, error) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		// log.Errorf(err, "Cannot generate UUID:%s")
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return uuid, err
}

// hash plaintext with SHA-1
func EncryptSha1(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

// hash plaintext with md5
func EncryptMd5(source string) (cryptext string) {
	// 使用MD5加密
	signBytes := md5.Sum([]byte(source))
	// 把二进制转化为大写的十六进制
	cryptext = hex.EncodeToString(signBytes[:])
	return
}

// func VerifyEmailFormat(email string) bool {
// 	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
// 	reg := regexp.MustCompile(pattern)
// 	return reg.MatchString(email)
// }
