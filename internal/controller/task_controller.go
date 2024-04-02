package controller

import (
	"go-template-api/common"
	"go-template-api/internal/db"
	"go-template-api/internal/model"

	"log"
)

func CreateTask(task model.Task_Table) common.SqlQueryStatus {
	result := db.Conn.Create(&task)
	if result.Error != nil {
		log.Println("Error creating Task:", result.Error)
		return common.SqlQueryStatus{Message: "cannot create Task", Code: 0, Err: result.Error}
	}
	return common.SqlQueryStatus{Message: "Task created successfully", Code: 1, Err: nil}
}

func IsTaskExist(taskCode string) bool {
	var task *model.Task_Table
	db.Conn.Find(&task, "code=?", taskCode)
	return task != nil
}
