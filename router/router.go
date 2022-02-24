/**
 * @Author: Administrator
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/2/24 16:10
 */

package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/gin-gonic/gin"
)

func InitSelfRouter(Router *gin.RouterGroup) {
	SelfRouter := Router.Group("/self")
	{
		SelfRouter.GET("/test", system.TestRouter)
	}
}
