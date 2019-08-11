package db

import (
	"time"
)

type Course struct {
	Id        int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	UserId    int       `json:"user_id" xorm:"comment('用户id') INT(11)"`
	Course    string    `json:"course" xorm:"comment('课程') VARCHAR(20)"`
	CreatedAt time.Time `json:"created_at" xorm:"TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" xorm:"TIMESTAMP"`
}
