package db

import (
	"time"
)

type Course struct {
	Id        int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	UserId    int       `xorm:"comment('用户id') INT(11)"`
	Course    string    `xorm:"comment('课程') VARCHAR(20)"`
	CreatedAt time.Time `xorm:"TIMESTAMP"`
	UpdatedAt time.Time `xorm:"TIMESTAMP"`
	DeletedAt time.Time `xorm:"TIMESTAMP"`
}
