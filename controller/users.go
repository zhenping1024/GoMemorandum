package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct{
	Username string
	jwt.StandardClaims
}
func InitJWT( uname string)(s string){
	mySigningKey:=[]byte("usertest")
	c:=MyClaims{
		Username: uname,
		StandardClaims:jwt.StandardClaims{
			NotBefore: time.Now().Unix()-60,
			ExpiresAt: time.Now().Unix()+60*60*2,
			Issuer: "user2",
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	s,e:=token.SignedString(mySigningKey)
	if e!=nil{
		fmt.Println(e.Error())
		return
	}else{
		fmt.Println(s)
		return s
	}
}
func ParseJwt(s string)(username string,err error){
	mySigningKey:=[]byte("usertest")
	t,err:=jwt.ParseWithClaims(s,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey,nil
	})
	if err!=nil{
		fmt.Println(err.Error())
		return username,err
	}else{
		fmt.Println(t.Claims.(*MyClaims).Username)
		return t.Claims.(*MyClaims).Username,nil
	}
}