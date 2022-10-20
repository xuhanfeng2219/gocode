package vertify

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var jwtKey = []byte("www.baidu.com")
var str string

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func main() {
	r := gin.Default()
	r.GET("/set", Setting)
	r.GET("/get", Getting)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("run wrong!")
	}
}

func Setting(c *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "localhost",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
	}

	str = tokenStr
	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

func Getting(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "not enough Authorization"})
		c.Abort()
		return
	}

	token, claims, err := ParseToken(tokenStr)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "not enough Authorization"})
		c.Abort()
		return
	}

	fmt.Println(claims.UserId)
}

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
