package Models

import (
	"encoding/json"
	"errors"
	"gin/Models/db"
	"gin/Service/Database"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
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
func Insert(c *gin.Context) (err error) {
	data,_ := ioutil.ReadAll(c.Request.Body)
	if data == nil {
		err = errors.New("数据不能为空")
		return
	}
	m := map[string]string{}
	_ = json.Unmarshal([]byte(data),&m)
	age,_ := strconv.Atoi(m["age"])
	person := db.User{
		Name:m["name"],
		Age:age,
		Gender:m["gender"],
	}
	_ , err = Database.Db.Insert(person)
	if err != nil {
		err = errors.New("插入失败")
		return
	}
	return
}
