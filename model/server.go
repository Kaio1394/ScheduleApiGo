package model

type Server struct {
	Id  int    `json:"id" gorm:"primaryKey"`
	Tag string `json:"tag"`
	Ip  string `json:"ip"`
}

func (Server) TableName() string {
	return "t_server_register"
}
