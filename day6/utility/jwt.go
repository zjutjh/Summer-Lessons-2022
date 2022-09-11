package utility

import (
	"class-schedule/config"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtData struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func GenerateStandardJwt(jwtData *JwtData) string {
	claims := jwtData
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(time.Duration(config.Config.Jwt.Expires) * time.Hour)).Unix(),
		Issuer:    config.Config.Jwt.Issuer,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.Config.Jwt.Secret))
	if err != nil {
		log.Fatalln("Jwt Error", err)
		panic(err)
	}
	return token
}

func ParseToken(token string) (string, error) {
	jwtSecret := []byte(config.Config.Jwt.Secret)
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtData{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JwtData); ok && tokenClaims.Valid {
			return claims.ID, err
		}
	}
	return "", err
}
