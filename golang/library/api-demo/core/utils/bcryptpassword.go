package utils

import "golang.org/x/crypto/bcrypt"

// 因为 bcrypt 加密使用了随机的盐，所以同一个字串每次 hash 的结果也是不一样的
// 同一字串的加密结果都是等价

// BcryptPasswordHash 进行加密
func BcryptPasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// BcryptPasswordCheck 对比密码哈希值
func BcryptPasswordCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
