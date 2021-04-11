package controller

import (
	"awesomeProject8/dao"
	"awesomeProject8/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
func Register(context*gin.Context){
	//1.接受前端所有注册信息
	//少一步数据验证
	//数据绑定
	var  acc models.User
	context.ShouldBind(&acc)

	var TmpAcc models.User
	err:=dao.DB.Where("user_name = ?",acc.UserName).First(&TmpAcc).Error
	//fmt.Println(acc.Name,TmpAcc.Name)
	if err!=nil{

		dao.DB.Create(&acc)
		context.JSON(http.StatusCreated,gin.H{
			"status":0,
			"message":"success",
			"data":acc,
		})
		//context.Redirect(http.StatusMovedPermanently,"/v1/api/login")
	}else{
		context.JSON(http.StatusOK,gin.H{
			"error":"用户已存在",
		})
		//context.Redirect(http.StatusMovedPermanently,"/v1/api/login")
	}
	////4.跳转登陆页面'
	//context.Redirect(http.StatusMovedPermanently,"127.0.0.1:9090/v1/api/login")

}
func Login(context*gin.Context){
	var  acc models.User
	//context.ShouldBind(&acc)
	acc.UserName=context.PostForm("username")
	acc.Password=context.PostForm("password")
	//查询用户是否存在
	var tmpuser models.User
	dao.DB.Debug().Where("user_name = ?",acc.UserName).First(&tmpuser)
	if tmpuser.UserName==acc.UserName{
		if tmpuser.Password==acc.Password{
			token:=InitJWT(tmpuser.UserName)
			//InitToken()
			context.JSON(http.StatusCreated,gin.H{
				"status":0,
				"message":"success",
				"data":token,
			})
		//context.JSON(http.StatusCreated,token)
		}else{
			context.JSON(200,"密码错误")
		}
	}else{
		context.JSON(200,"此用户不存在")
	}
}
//添加一个待办事项
func AddOne(c*gin.Context){
	//1.从请求中得到数据
	var m models.Memorandum
	var u models.User
	c.BindJSON(&m)
	s:=c.Request.Header.Get("token")
	//fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	fmt.Println("prase name is",username)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		//username:=c.Param("user")
		dao.DB.First(&u,"user_name = ?",username)
		fmt.Println(u)
		err:=dao.AddList(u,m)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			c.JSON(http.StatusCreated,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}

}
//查询所有事项
func FindAll(c*gin.Context){
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		fmt.Println("username is",username)
		//PraseToken(s)
		var u models.User
		spage:=c.Param("page")
		page,_:=strconv.Atoi(spage)
		dao.DB.Debug().First(&u,"user_name = ?",username)
		var m []models.Memorandum
		var err error
		m,err=dao.SearchAll(u,page)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			dao.PushList(m)
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}

}
//查询待办事项
func FindTodo(c*gin.Context){
	var m []models.Memorandum
	var u models.User
	var err error
	spage:=c.Param("page")
	page,_:=strconv.Atoi(spage)
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		//username:=c.Param("user")
		dao.DB.Debug().First(&u,"user_name = ?",username)
		m,err=dao.SearchTodo(u,page)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			dao.PushList(m)
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}
}
//查询已完成事件
func FindDone(c*gin.Context){
	var m []models.Memorandum
	var u models.User
	var err error
	spage:=c.Param("page")
	page,_:=strconv.Atoi(spage)
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		//username:=c.Param("user")
		dao.DB.Debug().First(&u,"user_name = ?",username)
		m,err=dao.SearchDone(u,page)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			dao.PushList(m)
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}
}
//关键字查询
func FindKey(c*gin.Context){
	var u models.User
	var err error
	spage:=c.Param("page")
	page,_:=strconv.Atoi(spage)
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		//username:=c.Param("user")
		dao.DB.Debug().First(&u,"user_name = ?",username)
		var m []models.Memorandum
		key:=c.Param("key")
		key="%"+key+"%"
		m,err=dao.SearchKey(u,key,page)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			dao.PushList(m)
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}
}
//查询历史记录
func History(c*gin.Context){
	var u models.User
	//var err error
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":e.Error(),
		})
	}else{
		dao.DB.Debug().First(&u,"user_name = ?",username)
		historylist,err:=dao.ListHistory(strconv.Itoa(int(u.ID)))
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":historylist,
			})
		}

	}

}
//完成一事项
func FinishOne(c*gin.Context){
	var err error
	var u models.User
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		//username:=c.Param("user")
		dao.DB.Debug().First(&u,"user_name = ?",username)
		id:=c.Param("id")
		var m models.Memorandum
		m,err=dao.FinishOne(u,id)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}

	//if err:=dao.DB.Where("id = ?",id).First(&m).Error;err!=nil{
	//	c.JSON(http.StatusOK,gin.H{
	//		"error":err.Error(),
	//	})
	//	//c.JSON(http.StatusOK,gin.H{
	//	//	"status":500,
	//	//	"message":"change error",
	//	//	"data":err.Error(),
	//	//})
	//}else{
	//	//m.FinishFlag="true"
	//	//c.BindJSON(&m)
	//	if err=dao.DB.Model(&m).Update("finish_flag","true").Error;err!=nil{
	//		c.JSON(http.StatusOK,gin.H{
	//			"error":err.Error(),
	//		})
	//		//c.JSON(http.StatusOK,gin.H{
	//		//	"status":500,
	//		//	"message":"change error",
	//		//	"data":err.Error(),
	//		//})
	//	}else{
	//		c.JSON(http.StatusOK,gin.H{
	//			"status":0,
	//			"message":"change success",
	//			"data":m,
	//		})
	//	}
	//
	//}
}
//完成全部
func FinishAll(c*gin.Context){
	var err error
	var u models.User
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		//username:=c.Param("user")
		dao.DB.First(&u,"user_name = ?",username)
		var m []models.Memorandum
		m,err=dao.FinishAll(u)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}
}
//未完一
func UndoneOne(c*gin.Context){
	var err error
	var u models.User
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		//username:=c.Param("user")
		dao.DB.Debug().First(&u,"user_name = ?",username)
		id:=c.Param("id")
		var m models.Memorandum
		m,err=dao.UndoneOne(u,id)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}

}
//未完全部
func UndoneAll(c*gin.Context){
	var err error
	var u models.User
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		//username:=c.Param("user")
		dao.DB.First(&u,"user_name = ?",username)
		var m []models.Memorandum
		m,err=dao.UndoneAll(u)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}

}
//删除一
func DeleteOne(c*gin.Context){
	var err error
	var u models.User
	id:=c.Param("id")
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		//username:=c.Param("user")
		var m models.Memorandum
		dao.DB.Debug().First(&u,"user_name = ?",username)
		m,err=dao.DeleteOne(u,id)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}

	//}
}
//删除所有
func DeleteAll(c*gin.Context){
	var err error
	var u models.User
	s:=c.Request.Header.Get("token")
	fmt.Println("get token" ,s)
	username,e:=ParseJwt(s)
	if e!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":e.Error(),
		})
	}else{
		//username:=c.Param("user")
		dao.DB.First(&u,"user_name = ?",username)
		var m []models.Memorandum
		m,err=dao.DeleteAll(u)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"status":0,
				"message":"success",
				"data":m,
			})
		}
	}

}