package middleware

import (
	"errors"
	//"fmt"
	"net/http"
	"time"

	//"errors"

	"todo_api/internal/config"
	"todo_api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

//var UserLogedIn models.User
func Login_Required(c *gin.Context){
	//fmt.Printf("I am in the Login_Required middleware")
	
	tokenString,err:=c.Cookie("Authorization")
	//fmt.Print(tokenString)
	if err!=nil{
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//2 Decode/validate it 
	SECRET_KEY:="poiuytrewlkjhgfdsamnbvcxz"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(SECRET_KEY), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		//log.Fatal(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//3 Check exp
		if float64(time.Now().Unix())> claims["exp"].(float64){
			c.AbortWithStatus(http.StatusRequestTimeout)
			
		}
		//4 find the user with token sub
		var user models.User
		config.DB.First(&user,claims["sub"])

		if user.ID==0{
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//Attach to request
		c.Set("user",user)

		//continue
		c.Next()

		//fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}


}

func UserInfo(c *gin.Context) (models.User,error){
	//fmt.Println("function user info")
	userInterface,exists:=c.Get("user")
	if !exists{
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
        return	models.User{},errors.New("invalid user type in context")
	}
	user:=userInterface.(models.User)


	return user,nil
	
}

func UserWithTodos(user *models.User) error {
    err := config.DB.
        Preload("Todos"). 
        First(user, user.ID).Error

    if err != nil {
        return err
    }
    return nil
}