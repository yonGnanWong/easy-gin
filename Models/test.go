package Models

import (
	"gin/Models/db"
	"gin/Service/Database"
	"log"
)

//获取数据
func GetTestData() *db.User  {
	var test db.User
	// SELECT * FROM user Where id = 1
	result ,err := Database.Db.ID(1).Get(&test)
	if !result  {
		log.Print(err)
		return nil
	}else {
		return &test
	}
}

//插入数据
//func Insert(p *http.Request)  {
//	param := p.GetBody
//	var data db.User
//	_ = json.Unmarshal(p,data)
//
//}
