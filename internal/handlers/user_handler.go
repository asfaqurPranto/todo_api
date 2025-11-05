package handlers

import (
	//"fmt"

	"net/http"
	"time"
	"todo_api/internal/config"
	"todo_api/internal/middleware"
	"todo_api/internal/models"

	//"todo_api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context){
	var LoginJson models.LoginInfo

	err:=c.BindJSON(&LoginJson)
	if err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
	}

	var user models.User
	config.DB.First(&user,"Email=?",LoginJson.Email)
	if user.ID==0{		//user not found
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid email no user found with this email"})
		return
	}
	//user found (now compare input password with saved hashed pass)
	err=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(LoginJson.Password))
	if err!=nil{		//doesnot match
		c.JSON(http.StatusBadRequest,gin.H{"error":"wrong password"})
		return
	}
	//password matched now generate token for a timespan(1 month)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": user.ID,
	"exp": time.Now().Add(time.Hour*24*30).Unix(),

	})	

	//now send the the encoded token
	SECRET_KEY:="poiuytrewlkjhgfdsamnbvcxz"
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	//c.JSON(http.StatusOK,gin.H{"Token":tokenString})
	//thats not a good way of sending token

	// Better ways are sending in cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",tokenString, 3600*24*30,"","",false,true )

}

func Register(c *gin.Context){

	//get email,pass,name from frontend and bind
	var RegData models.RegisterInfo
	err:=c.BindJSON(&RegData)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	//fmt.Println(RegData)

	//Generate hash pasword
	hash,err:=bcrypt.GenerateFromPassword([]byte(RegData.Password),10)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Failed to hash password"})
		return
	}

	//so we will be using hashed password , not the actual one
	user:=models.User{Name:RegData.Name,Email:RegData.Email,Password:string(hash)}
	result:=config.DB.Create(&user)

	if result.Error!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":result.Error.Error()})
		return
	}


	c.JSON(http.StatusAccepted,gin.H{"message":"User Created Successfully"})
	



}

func Info(c *gin.Context){
	//fmt.Println("function user info")
	user,err:=middleware.UserInfo(c)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	_=middleware.UserWithTodos(&user)
	c.JSON(http.StatusOK, gin.H{
		"Name":user.Name,
		"email":user.Email,
		"Todos":user.Todos,
	})
	

}