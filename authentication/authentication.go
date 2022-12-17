package authentication

import (
	"errors"
	"time"
  "gorm.io/gorm"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("Vzd2uNBkUm7JxdYRhWsy7qNlsikga0wN")
var Claims *JWTClaim
var UserScope *gorm.DB

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))

	claims:= &JWTClaim{
		Email: email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
		err = errors.New("token expired")
		return
	}
	Claims = claims

	return
}
