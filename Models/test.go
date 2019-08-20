package Models

import (
	"errors"
	"fmt"
	"gin/Models/db"
	"gin/Service/Database"
	"log"
)

//读取
func Get() (test db.User, err error) {
	// SELECT * FROM user Where id = 1
	_ ,err = Database.Db.ID(1).Get(&test)
	if err != nil {
		log.Print(err)
		err = errors.New("系统错误")
		return
	}
	return
}

//插入数据
func Insert(d db.User) (err error) {
	//redis 插入
	_,err = Database.Redis.Client.Set("ceshi","nihao",0).Result()
	if err != nil {
		log.Print(err)
		err = errors.New("插入失败")
		return
	}
	_, err = Database.Db.Insert(d)
	if err != nil {
		log.Print(err)
		err = errors.New("插入失败")
		return
	}
	return
}

//更新
func Update(d db.User) (err error) {
	var bean db.User
	fmt.Print(d)
	_, err = Database.Db.Select("id").Where("name = ?",d.Name).Get(&bean)
	if err != nil {
		log.Print(err)
		return
	}
	_, err = Database.Db.Id(bean.Id).Update(d)
	return
}
//删除
func Delete(id int) (err error) {
	var bean  db.User
	rows,err := Database.Db.Id(id).Delete(bean)
	if rows == 0 || err != nil {
		err = errors.New("删除失败")
	}
	return
}
