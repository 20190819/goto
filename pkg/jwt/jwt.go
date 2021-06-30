package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

type JwtServer struct {
}

type CustomClaims struct {
	Id      uint64
	ExpTime int64
	jwt.StandardClaims
}

func CreateToken(*JwtServer) (map[string]interface{}, error) {
	maxAage, _ := strconv.Atoi(viper.GetString("jwt.expire"))
}

func (*JwtServer) AuthToken(token string) bool {
	if token == "" {
		return false
	}

	content := strings.Split(token, "")
	if content[0] != "Bearer" {
		return false
	}

	tokenParse, _ := jwt.Parse(content[1], func(tokenStruct *jwt.Token) (interface{}, error) {
		if _, ok := tokenStruct.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tokenStruct.Header["alg"])
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if _, ok := tokenParse.Claims.(jwt.MapClaims); ok && tokenParse.Valid {
		return true
	} else {
		return false
	}
}
