package middlewares

// import (
// 	"time"

// 	"github.com/ghynmo/MVC_Learning/config"
// 	jwt "github.com/dgrijalva/jwt-go"
// )

// func createToken(userID int) (token string, err error){
// 	claims := jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["userid"] = userID
// 	claims["exp"] = time.Now().Add(time.Hour *1).Unix()

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(config.JWT_SECRET))
// }
