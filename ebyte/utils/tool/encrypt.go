package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashAndSalt
//
//	@Description: 将明文密钥加密
//	@param pwd	明文密钥
//	@return string	密文密钥
//	@return error	错误处理
func HashAndSalt(pwd string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// 验证密码
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
