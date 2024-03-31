package test

import (
	"go-template-api/controller"
	"go-template-api/db"
	"testing"
)

func TestIsEmailExist(t *testing.T) {
	db.ConnectDb()
	resualt := controller.IsEmailExist("med@gmail.com")
	t.Errorf("IsEmailExist() = %v, want %v", resualt, true)

}
