package main

import (
	"awesomeProject8/controller"
	"awesomeProject8/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.DbInit()
	dao.InitClient()
   R:=gin.Default()

	ToDOGroup:=R.Group("/v1/api")
	{
		//注册
		ToDOGroup.POST("/register",controller.Register)
		//登录
		ToDOGroup.POST("/login",controller.Login)
		//创建一个新的待办事项
		ToDOGroup.POST("/todolist/creat",controller.AddOne)
		//查看所有事项
		ToDOGroup.GET("/todolist/all/:page",controller.FindAll)
		//查看待办事项
		ToDOGroup.GET("/todolist/todo/:page",controller.FindTodo)
		//查看已完成事项
		ToDOGroup.GET("/todolist/done/:page",controller.FindDone)
		//输入关键字查询
		ToDOGroup.GET("/todolist/get/:key/:page",controller.FindKey)
		//查询历史记录
		ToDOGroup.GET("todolist/history",controller.History)
		//修改一条待办——>完成
		ToDOGroup.PUT("/todolist/finish/:id",controller.FinishOne)
		//修改所有待办——>完成
		ToDOGroup.PUT("/todolist/finish",controller.FinishAll)
		//修改一条完成——>代办
		ToDOGroup.PUT("/todolist/undone/:id",controller.UndoneOne)
		//修改所有完成——>代办
		ToDOGroup.PUT("/todolist/undone",controller.UndoneAll)
		//删除一条
		ToDOGroup.DELETE("/todolist/delete/:id",controller.DeleteOne)
		//删除所有
		ToDOGroup.DELETE("/todolist/delete",controller.DeleteAll)
	}


   R.Run(":9090")
}
