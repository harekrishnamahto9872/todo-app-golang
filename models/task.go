package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task for responses from mongodb queries
type Task struct {
	ID                   primitive.ObjectID `bson:"_id" json:"_id"`
	Title                string             `json:"title"`
	Description          string             `json:"description"`
	DueDate              string             `json:"dueDate"`
	CreatedAt            int64              `json:"createdAt"`
	UpdatedAt            int64              `json:"updatedAt"`
	User                 string             `json:"user"`
	IsSubtask            bool               `json:"isSubtask"`
	ParentTaskID         string             `json:"parentTaskID"`
	HasReminder          bool               `json:"hasReminder"`
	NoOfHoursForReminder int64              `json:"noOfHoursForReminder"`
}

func (task *Task) SetCreatedAt() {
	task.CreatedAt = time.Now().Unix()
}

func (task *Task) SetUpdatedAt() {
	task.UpdatedAt = time.Now().Unix()
}

// NewTask For mongodb insert
type NewTask struct {
	Title                string `json:"title"`
	Description          string `json:"description"`
	DueDate              string `json:"dueDate"`
	CreatedAt            int64  `json:"createdAt"`
	UpdatedAt            int64  `json:"updatedAt"`
	User                 string `json:"user"`
	IsSubtask            bool   `json:"isSubtask"`
	ParentTaskID         string `json:"parentTaskID"`
	HasReminder          bool   `json:"hasReminder"`
	NoOfHoursForReminder int64  `json:"noOfHoursForReminder"`
}

// SetCreatedAt //
func (task *NewTask) SetCreatedAt() {
	task.CreatedAt = time.Now().Unix()
}

// SetUpdatedAt //
func (task *NewTask) SetUpdatedAt() {
	task.UpdatedAt = time.Now().Unix()
}
