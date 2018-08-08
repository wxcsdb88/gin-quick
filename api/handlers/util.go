package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wxcsdb88/gin-quick/api/common"
)

// ResponseJSON response json
func ResponseJSON(c *gin.Context, responseData *common.ResponseData) {
	code := responseData.Code
	c.JSON(code, responseData)
}
