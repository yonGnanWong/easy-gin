package db

import (
	"time"
)

type User struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Name      string    `xorm:"comment('姓名') VARCHAR(128)"`
	Age       int       `xorm:"comment('年龄') INT(3)"`
	Gender    string    `xorm:"comment('性别') VARCHAR(128)"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	DeletedAt time.Time `xorm:"TIMESTAMP"`
}
