package v1

import (
	"fmt"
	"gin/Models"
	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	a := Models.GetTestData()
	fmt.Print(a)
	c.JSON(200,"this is get method")
	return
}

func Post(c *gin.Context) {
	c.JSON(200, "this is post method")
	return
}
