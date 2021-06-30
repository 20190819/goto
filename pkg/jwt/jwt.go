package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

type JwtServer struct {
}

type CustomClaims struct {
	ExpTime int64
	jwt.StandardClaims
}

func CreateToken(*JwtServer) (map[string]interface{}, error) {
	maxAage, _ := strconv.Atoi(viper.GetString("jwt.expire"))
	expTime := time.Now().Add(time.Duration(maxAage) * time.Second).Unix()
	customClaims := CustomClaims{
		ExpTime: expTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenStr, err := token.SignedString(viper.GetString("jwt.secret"))
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	result["expTime"] = expTime
	result["token"] = tokenStr
	return result, nil
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
