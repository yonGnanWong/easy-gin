package db

import (
	"time"
)

type User struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string    `json:"name" xorm:"comment('姓名') VARCHAR(128)"`
	Age       int       `json:"age" xorm:"comment('年龄') INT(3)"`
	Gender    string    `json:"gender" xorm:"comment('性别') VARCHAR(128)"`
	CreatedAt time.Time `json:"created_at" xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"comment('更新时间') TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" xorm:"TIMESTAMP"`
}
