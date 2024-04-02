package controller

import (
	"go-template-api/common"
	"go-template-api/internal/db"
	"go-template-api/internal/model"

	"go-template-api/utils"
)

func CreateNewUser(user *model.User_Table) common.SqlQueryStatus {
	if IsEmailExist(user.Email) {
		return common.SqlQueryStatus{Message: "Email is already used", Code: 0, Err: nil}
	}
	if IsPhoneExist(user.Phone) {
		return common.SqlQueryStatus{Message: "Phone is already used", Code: 0, Err: nil}
	}
	result := db.Conn.Create(&user)
	if result.Error != nil {
		return common.SqlQueryStatus{Message: "cannot create user", Code: 0, Err: result.Error}
	}
	user.Code = utils.GenerateNewCode(user.ID)
	UpdateCode(*user)
	return common.SqlQueryStatus{Message: "User created successfully", Code: 1, Err: nil}
}

func UpdateCode(user model.User_Table) error {
	result := db.Conn.Model(&user).Update("code", user.Code)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CheckCredentials(user *model.User_Table) common.SqlQueryStatus {
	if user.Email == "" {
		if !IsPhoneExist(user.Phone) {
			return common.SqlQueryStatus{Message: "Phone not found", Code: 0, Err: nil}
		}
		if !IsPwdCorrect("phone = ? AND password = ?", user.Phone, user.Password) {
			return common.SqlQueryStatus{Message: "Password incorrect", Code: 0, Err: nil}
		}
	} else if user.Phone == "" {
		if !IsEmailExist(user.Email) {
			return common.SqlQueryStatus{Message: "Email not found", Code: 0, Err: nil}
		}
		if !IsPwdCorrect("email = ? AND password = ?", user.Email, user.Password) {
			return common.SqlQueryStatus{Message: "Password incorrect", Code: 0, Err: nil}
		}
	}

	return common.SqlQueryStatus{Message: "Jwt Processing...!", Code: 1, Err: nil}
}

func IsEmailExist(email string) bool {
	var user model.User_Table
	db.Conn.Where("email = ?", email).Find(&user, "email = ?", email)
	return user.ID != 0
}

func IsPhoneExist(phne string) bool {
	var user model.User_Table
	db.Conn.Find(&user, "phone = ?", phne)
	return user.ID != 0
}

func IsPwdCorrect(condition, att, pwd string) bool {
	var user model.User_Table
	db.Conn.Where(condition, att, pwd).Find(&user)
	return user.ID != 0
}
