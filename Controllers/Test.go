package Controllers

import (
	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context)  {
	c.JSON(200,"hello word!")
	return
}

func Post(c *gin.Context)  {
	c.JSON(200,"this is post method")
	return
}
