package user

import "GOHUB/pkg/database"

// 判断email是否被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email =?", email).Count(&count)
	return count > 0
}

//判断手机号是否被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone=?", phone).Count(&count)
	return count > 0
}
