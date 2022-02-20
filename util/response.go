package util

import (
	"github.com/harekrishnamahto9872/todo-app-golang/models"
)

// ResMessage //
// response struct
type ResMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ResError //
// response struct
type ResError struct {
	Success bool  `json:"success"`
	Error   error `json:"message"`
}

// ResUser //
// response struct
type ResUser struct {
	Success bool           `json:"success"`
	Message models.UserRes `json:"message"`
}

// ResTask //
// response struct
type ResTask struct {
	Success bool        `json:"success"`
	Message models.Task `json:"message"`
}

// ResTasks //
// response struct
type ResTasks struct {
	Success bool          `json:"success"`
	Message []models.Task `json:"message"`
}
