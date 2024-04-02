package test

import (
	"fmt"
	"go-template-api/internal/controller"
	"go-template-api/internal/db"
	"go-template-api/internal/model"

	"testing"
	"time"
)

func Test_createTask(t *testing.T) {
	db.ConnectDb()
	got := controller.CreateTask(model.Task_Table{
		ID:      1,
		Code:    "task1",
		Cost:    1500.50,
		StartAt: time.Now(),
		UserID:  1, Status: "on process",
	})
	fmt.Printf("createTask() = %v, want %v", got.Code, 1)
}

func TestIsTaskExist(t *testing.T) {
	db.ConnectDb()
	got := controller.IsTaskExist("d5zADF")
	fmt.Printf("createTask() = %v, want %v", got, true)
}
