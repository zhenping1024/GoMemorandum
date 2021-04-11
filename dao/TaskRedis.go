package dao

import (
	"awesomeProject8/models"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

// 声明一个全局的rdb变量
var Rdb *redis.Client

// 初始化连接
func InitClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB

	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
//将查询事件标题存入数据库
func PushList(todotitle []models.Memorandum) {
	for i:=0;i< len(todotitle);i++{
		data,_:=json.Marshal(todotitle[i])
		result,err:=Rdb.LPush(strconv.Itoa(int(todotitle[i].UserId)),data).Result()
		if err!=nil {
			fmt.Println(err)
			return
		}
		fmt.Println(todotitle,result)
	}

	if lenlist:=ListLen(strconv.Itoa(int(todotitle[0].UserId)));lenlist>10{
		ListDelete(strconv.Itoa(int(todotitle[0].UserId)))
	}
}
//返回记录条数
func ListLen(key string)(listlen int64) {
	listlen,err:=Rdb.LLen(key).Result()
	if err!=nil {
		fmt.Println(err)
		return
	}
	return listlen
}
//删除多余记录
func ListDelete(key string) {
	result,err:=Rdb.LTrim(key,0,9).Result()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
//返回历史记录
func ListHistory(key string) (resultstruct []models.Memorandum,err error){
	result:=Rdb.LRange(key,0,9).Val()
	var showresult models.Memorandum
	//var resultstruct []models.Memorandum
	for i:=0;i< len(result);i++{
		_=json.Unmarshal([]byte(result[i]),&showresult)
		resultstruct=append(resultstruct,showresult)
		//fmt.Println("history is ",showresult)
	}
	//fmt.Println(resultstruct)
	return resultstruct,nil
}

