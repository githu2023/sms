package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	mathrand "math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MD5V md5加密
func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// GenerateRandomString generates a cryptographically secure random string of specified length
func GenerateRandomString(length int) string {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

// GenerateRandomMerchantNo generates a random 6-digit merchant number
func GenerateRandomMerchantNo() string {
	r := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(100000 + r.Intn(900000))
}
