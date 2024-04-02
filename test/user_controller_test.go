package test

import (
	"go-template-api/internal/controller"
	"go-template-api/internal/db"
	"go-template-api/internal/model"
	"log"
	"testing"
)

func TestIsEmailExist(t *testing.T) {
	db.ConnectDb()
	resualt := controller.IsEmailExist("med@gmail.com")
	log.Printf("IsEmailExist() = %v, want %v", resualt, true)
}

func TestIsPwdCorrect(t *testing.T) {
	db.ConnectDb()
	resualt := controller.IsPwdCorrect("0658185867", "45JiohD8")
	log.Printf("TestIsPwdCorrect() = %v, want %v", resualt, true)
}

func TestUpdateCode(t *testing.T) {
	db.ConnectDb()
	resualt := controller.UpdateCode(model.User_Table{ID: 1, Code: ""})
	log.Printf("TestUpdateCode() = %v, want %v", resualt, nil)
}
