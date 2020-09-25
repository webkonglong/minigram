package routes

import (
	"net/http"

	controllers "github.com/flyq/minigram/controllers"
	"github.com/gin-gonic/gin"
)

// 对于某一确定小票，其电子小票和区块链小票id相同
func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.POST("/v1/user", controllers.CreateUser)        // 创建用户
	router.POST("/v1/elerec", controllers.CreateElerec)    // 创建电子小票
	router.POST("/v1/blorec", controllers.CreateBlorec)    // 创建区块链小票
	router.GET("/v1/user/:userId", controllers.GetUser)    // 某用户的小票记录
	router.GET("/v1/elerec/:recId", controllers.GetElerec) // Electronic receipt 某具体电子小票详细信息
	router.GET("/v1/blorec/:recId", controllers.GetBlorec) // blockchain receipt 某具体区块链小票详细信息
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
