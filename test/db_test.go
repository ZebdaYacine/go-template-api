package test

import (
	"go-template-api/controller"
	"go-template-api/db"
	"testing"
)

func TestIsPhoneExist(t *testing.T) {
	db.ConnectDb()
	result := controller.IsPhoneExist("0658185867")
	expected := true
	if result != expected {
		t.Errorf("IsPhoneExist = %v; want %v", result, expected)
	}
}
