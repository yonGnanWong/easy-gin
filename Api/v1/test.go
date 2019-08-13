package v1

import (
	"gin/Models"
	"github.com/gin-gonic/gin"
)

func C(c *gin.Context) {
	err := Models.Insert(c)
	if err != nil {
		c.JSON(400,err.Error())
	}
	c.JSON(200, "插入成功")
	return
}

func R(c *gin.Context) {
	a := Models.GetTestData()
	c.JSON(200,a)
	return
}

func U(c *gin.Context)  {
	c.JSON(200,"更新成功")
	return
}

func D(c *gin.Context)  {
	c.JSON(200,"删除成功")
	return
}
