package v1

import (
	"fmt"
	"gin/Models"
	"gin/Models/db"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func C(c *gin.Context) {
	var data db.User
	err := c.BindJSON(&data)
	if err != nil{
		c.JSON(400,err.Error())
		return
	}
	err = Models.Insert(data)
	if err != nil {
		c.JSON(400,err.Error())
		return
	}
	c.JSON(200,data)
	return
}

func R(c *gin.Context) {
	data,err := Models.Get()
	if err != nil {
		c.JSON(400,err.Error())
		return
	}
	c.JSON(200,data)
	return
}

func U(c *gin.Context)  {
	var data db.User
	err := c.BindJSON(&data)
	if err != nil{
		log.Print(err)
		c.JSON(400,err.Error())
		return
	}
	err = Models.Update(data)
	if err != nil {
		log.Print(err)
		c.JSON(400,err.Error())
		return
	}
	c.JSON(200,"更新成功")
	return
}

func D(c *gin.Context)  {
	id,_ := strconv.Atoi(c.Query("id"))
	fmt.Print(id)
	err := Models.Delete(id)
	if err != nil {
		c.JSON(400,err.Error())
		return
	}
	c.JSON(200,"删除成功")
	return
}
