package models

type task struct {
	title           string
	description     string
	due_date        string
	is_substask     bool
	parent_task_id  string
	user_id         string
	is_alert_active bool
	alert_time      string
}
