package model

type Job struct {
	Id          int    `json:"id gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	CmdExecute  bool   `json:"cmdExecute" gorm:"column:cmd_execute"`
	Script      string `json:"script" gorm:"column:script"`
	Date        string `json:"date" gorm:"column:date"`
	Hold        bool   `json:"hold" gorm:"column:hold"`
	ServerId    int    `json:"server_id" gorm:"column:server_id"`
}

func (Job) TableName() string {
	return "t_job_register"
}
