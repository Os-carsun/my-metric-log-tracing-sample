package calculater

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func plus(c *gin.Context) {
	a, err := strconv.Atoi(c.Query("a"))
	if err != nil {
		c.JSON(400, "a should be a num")
		return
	}

	b, err := strconv.Atoi(c.Query("b"))
	if err != nil {
		c.JSON(400, "b should be a num")
		return
	}
	c.JSON(200, a+b)
}

func sub(c *gin.Context) {
	a, err := strconv.Atoi(c.Query("a"))
	if err != nil {
		c.JSON(400, "a should be a num")
		return
	}

	b, err := strconv.Atoi(c.Query("b"))
	if err != nil {
		c.JSON(400, "b should be a num")
		return
	}
	c.JSON(200, a-b)
}
