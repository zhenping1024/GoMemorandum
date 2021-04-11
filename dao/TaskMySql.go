package dao

import (
	"awesomeProject8/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func DbInit(){
	//连接数据库
	//var err error
	DB,_=gorm.Open("mysql","root:wzp614177@tcp(127.0.0.1:3306)/go_web?charset=utf8&parseTime=true")
	//if err!=nil{
	//	panic(err)
	//}
	//defer DB.Close()
	//创建表，自动迁移（把结构体和数据表进行对应
	//DB.CreateTable(&models.Memorandum{})
	DB.AutoMigrate(&models.User{},&models.Memorandum{})
	//CreatUser()
	//CreatTodolist()
	//SearchList()
	//AddList()
}
//添加一个事项
func AddList(u models.User,m models.Memorandum) (err error){
	//var u models.User
	if err=DB.First(&u,"id = ?",).Error;err!=nil{
		return err
	}else{
		err=DB.Debug().Model(&u).Association("Todolist").Append(&m).Error
	}
	return err
}
//查询用户所有事项
func SearchAll(u models.User, page int)(m []models.Memorandum, err error){
	//var m []models.Memorandum
	if err=DB.Model(&u).Association("Todolist").Find(&m).Error;err!=nil{
		return nil,err
	}else{
		total:=DB.Model(&u).Association("Todolist").Find(&m).Count()
		pagesize:=5
		nowpage :=total/pagesize
		if total%pagesize!=0{
			nowpage++
		}
		e:=errors.New("超出可查页数")

		if nowpage<page{
			return nil,e
		}else{
			DB.Model(&u).Limit(pagesize).Offset((page-1)*pagesize).Association("Todolist").Find(&m)
		}
		return m,nil
	}
}
//查询用户待办事项
func SearchTodo(u models.User,page int)(m []models.Memorandum,err error){
	//fmt.Println("user is",u.UserName)
	if err=DB.Model(&u).Where("finish_flag = ?","false").Association("Todolist").Find(&m).Error;err!=nil{
		return nil,err
	}else{
		total:=DB.Model(&u).Where("finish_flag = ?","false").Association("Todolist").Find(&m).Count()
		pagesize:=5
		nowpage :=total/pagesize
		if total%pagesize!=0{
			nowpage++
		}
		e:=errors.New("超出可查页数")
		if nowpage<page{
			return nil,e
		}else{
			DB.Model(&u).Where("finish_flag = ?","false").Association("Todolist").Find(&m)
		}
		return m,nil
	}
	//	return m,nil
	//}
}
//查询用户已完成事项
func SearchDone(u models.User,page int)(m []models.Memorandum,err error){
	if err=DB.Model(&u).Where("finish_flag = ?","true").Association("Todolist").Find(&m).Error;err!=nil{
		return nil,err
	}else{
		total:=DB.Model(&u).Where("finish_flag = ?","true").Association("Todolist").Find(&m).Count()
		pagesize:=5
		nowpage :=total/pagesize
		if total%pagesize!=0{
			nowpage++
		}
		e:=errors.New("超出可查页数")
		if nowpage<page{
			return nil,e
		}else{
			DB.Model(&u).Where("finish_flag = ?","true").Association("Todolist").Find(&m)
		}
		return m,nil
	}
}
//关键字查询事项
func SearchKey(u models.User,key string,page int)(m []models.Memorandum,err error){

	if err=DB.Model(&u).Where("title LIKE ?",key).Association("Todolist").Find(&m).Error;err!=nil{
		return nil,err
	}else{
		total:=DB.Model(&u).Where("title LIKE ?",key).Association("Todolist").Find(&m).Count()
		pagesize:=5
		nowpage :=total/pagesize
		if total%pagesize!=0{
			nowpage++
		}
		e:=errors.New("超出可查页数")
		if nowpage<page{
			return nil,e
		}else{
			DB.Model(&u).Where("title LIKE ?",key).Association("Todolist").Find(&m)
		}
		return m,nil
	}
}
//完成一事
func FinishOne(u models.User,id string)(m models.Memorandum,err error){
	var tmp models.Memorandum
	err=DB.Model(&u).Where("id = ?",id).Association("Todolist").Find(&tmp).Error
	if err!=nil{
		fmt.Println(err.Error())
		return m,err

	}else{
		//tmp.FinishFlag="true"
		if err=DB.Model(&tmp).Update("finish_flag","true").Error;err!=nil{
			return m,err
		}else{
			return tmp,nil
		}
	}

	//if err=DB.Model(&u).Where("id = ?",id).Association("Todolist").Replace(m,tmp).Error;err!=nil{
	//	return tmp,err
	//}else{
	//	return tmp,nil
	//}
}
//完成全部
func FinishAll(u models.User)(m []models.Memorandum,err error){
	if err= DB.Table("memorandums").Where("user_id = ? AND Finish_flag = ?",u.ID,"false").Update(map[string]interface{}{"finish_flag":"true"}).Error;err!=nil{
		fmt.Println(m)
		return m,err
	}else{
		return m,nil
	}
}
//未完一事
func UndoneOne(u models.User,id string)(m models.Memorandum,err error){
	var tmp models.Memorandum
	err=DB.Model(&u).Where("id = ?",id).Association("Todolist").Find(&tmp).Error
	if err!=nil{
		fmt.Println(err.Error())
		return m,err

	}else{
		//tmp.FinishFlag="true"
		if err=DB.Model(&tmp).Update("finish_flag","false").Error;err!=nil{
			return m,err
		}else{
			return tmp,nil
		}
	}

	//if err=DB.Model(&u).Where("id = ?",id).Association("Todolist").Replace(m,tmp).Error;err!=nil{
	//	return tmp,err
	//}else{
	//	return tmp,nil
	//}
}
//未完全部
func UndoneAll(u models.User)(m []models.Memorandum,err error){
	if err= DB.Table("memorandums").Where("user_id = ? AND Finish_flag = ?",u.ID,"true").Update(map[string]interface{}{"finish_flag":"false"}).Error;err!=nil{
		return m,err
	}else{
		return m,nil
	}
}
//删除一事
func DeleteOne(u models.User,id string)(m models.Memorandum,err error){
	var tmp models.Memorandum
	err=DB.Model(&u).Where("id = ?",id).Association("Todolist").Find(&tmp).Error
	if err!=nil{
		fmt.Println(err.Error())
		return m,err

	}else{
		//tmp.FinishFlag="true"
		if err=DB.Model(&u).Association("Todolist").Delete(tmp).Error;err!=nil{
			return tmp,err
		}else{
			return tmp,err
		}
	}

}
//删除所有
func DeleteAll(u models.User)(m []models.Memorandum,err error){
	var tmp []models.Memorandum
	err=DB.Model(&u).Association("Todolist").Find(&tmp).Error
	if err!=nil{
		fmt.Println(err.Error())
		return m,err

	}else{
		//tmp.FinishFlag="true"
		if err=DB.Model(&u).Association("Todolist").Delete(tmp).Error;err!=nil{
			return tmp,err
		}else{
			return tmp,err
		}
	}
}
func DeleteList(){

}
func SearchList(){
	var u models.User
	DB.Preload("Todolist").First(&u,"id = ?",1)
	fmt.Println("user is",u)
}
func CreatUser(){
	u:=models.User{
		Password:"123",
		UserName: "user1",
	}
	_=DB.Create(&u)
}
func CreatTodolist(){
	list:=models.Memorandum{
	Title: "test1",
	Content: "111111",
}
	u:=models.User{
		Password:"123",
		UserName: "user1",
		Todolist:[]models.Memorandum {list},
	}

	_=DB.Create(&u)
}