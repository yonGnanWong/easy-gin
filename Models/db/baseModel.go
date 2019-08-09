package db

import "time"

type baseModel struct {
	Id int64
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
}
