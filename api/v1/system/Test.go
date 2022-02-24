/**
 * @Author: Administrator
 * @Description:
 * @File:  Test
 * @Version: 1.0.0
 * @Date: 2022/2/24 16:05
 */

package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestRouter(c *gin.Context) {
	c.String(http.StatusOK, "success!")
}
