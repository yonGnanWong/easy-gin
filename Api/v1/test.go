package v1

import (
	"gin/Models"
	"github.com/gin-gonic/gin"
)

func C(c *gin.Context) {
	c.JSON(200, "ha")
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
