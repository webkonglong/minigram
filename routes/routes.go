package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/flyq/minigram/controllers"
)

// 对于某一确定小票，其电子小票和区块链小票id相同
func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todo", controllers.CreateTodo)
	router.GET("/v1/user/:userId", controllers.GetUser) // 某用户的小票记录
	router.GET("/v1/elerec/:recId", controllers.GetElerec) // Electronic receipt 某具体电子小票详细信息
	router.GET("/v1/blorec/:recId", controllers.GetBlorec) // blockchain receipt 某具体区块链小票详细信息
	router.DELETE("/todo/:todoId", controllers.DeleteTodo)
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
