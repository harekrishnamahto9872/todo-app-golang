package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//for responses from mongodb queries
type Task struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DueDate     bool               `json:"dueDate"`
	CreatedAt   int64              `json:"createdAt"`
	UpdatedAt   int64              `json:"updatedAt"`
	User        string             `json:"user"`
}

func (task *Task) SetCreatedAt() {
	task.CreatedAt = time.Now().Unix()
}

func (task *Task) SetUpdatedAt() {
	task.UpdatedAt = time.Now().Unix()
}

// For mongodb insert
type NewTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     bool   `json:"dueDate"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	User        string `json:"user"`
}

// SetCreatedAt //
func (task *NewTask) SetCreatedAt() {
	task.CreatedAt = time.Now().Unix()
}

// SetUpdatedAt //
func (task *NewTask) SetUpdatedAt() {
	task.UpdatedAt = time.Now().Unix()
}
