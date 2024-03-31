package controller

import (
	"fmt"
	"go-template-api/common"
	"go-template-api/db"
	"go-template-api/model"
	"log"
)

func CreateNewUser(user model.User_Table) common.SqlQueryStatus {
	if isEmailExist(user.Email) {
		return common.SqlQueryStatus{Message: "Email is already used", Code: 0, Err: nil}
	}
	if isPhoneExist(user.Phone) {
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
	if !isPhoneExist(user.Phone) {
		return common.SqlQueryStatus{Message: "Phone not found", Code: 0, Err: nil}
	}
	if !isPwdCorrect(user.Phone, user.Password) {
		return common.SqlQueryStatus{Message: "Password incorrect", Code: 0, Err: nil}
	}
	return common.SqlQueryStatus{Message: "Jwt Processing...!", Code: 1, Err: nil}
}

func isEmailExist(email string) bool {
	var emails []string
	db.Conn.Where("email=?", email).Find(&emails)
	return len(emails) != 0
}

func isPhoneExist(phone string) bool {
	var phones string
	db.Conn.Where("phone=?", phone).Find(phones)
	fmt.Println(phones)
	return len(phones) != 0
}

func isPwdCorrect(phone, pwd string) bool {
	var pwds []string
	db.Conn.Where("phone = ? AND password = ?", phone, pwd).Find(&pwds)
	return len(pwds) != 0
}
