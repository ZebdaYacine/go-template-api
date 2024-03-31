package controller

import (
	"go-template-api/common"
	"go-template-api/db"
	"go-template-api/model"
	"log"
)

func CreateNewUser(user model.User_Table) common.SqlQueryStatus {
	if IsEmailExist(user.Email) {
		return common.SqlQueryStatus{Message: "Email is already used", Code: 0, Err: nil}
	}
	if IsPhoneExist(user.Phone) {
		return common.SqlQueryStatus{Message: "Phone is already used", Code: 0, Err: nil}
	}
	result := db.Conn.Create(&user)
	if result.Error != nil {
		log.Println("Error creating user:", result.Error)
		return common.SqlQueryStatus{Message: "cannot create user", Code: 0, Err: result.Error}
	}
	return common.SqlQueryStatus{Message: "User created successfully", Code: 1, Err: nil}

}

func CheckCredentials(user model.User_Table) common.SqlQueryStatus {
	if !IsPhoneExist(user.Phone) {
		return common.SqlQueryStatus{Message: "Phone not found", Code: 0, Err: nil}
	}
	if !IsPwdCorrect(user.Phone, user.Password) {
		return common.SqlQueryStatus{Message: "Password incorrect", Code: 0, Err: nil}
	}
	return common.SqlQueryStatus{Message: "Jwt Processing...!", Code: 1, Err: nil}
}

func IsEmailExist(email string) bool {
	var mail model.User_Table
	db.Conn.Find(&mail, "email=?", email)
	return mail.ID != 0
}

func IsPhoneExist(phone string) bool {
	var phones []model.User_Table
	db.Conn.Find(&phones, "phone=?", phone)
	return len(phones) != 0
}

func IsPwdCorrect(phone, pwd string) bool {
	var pwds []model.User_Table
	db.Conn.Where("phone = ? AND password = ?", phone, pwd).Find(&pwds)
	return len(pwds) != 0
}
