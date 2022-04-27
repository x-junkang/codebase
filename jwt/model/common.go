package model

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	ID string `json:"ID"`
	jwt.StandardClaims
}

var mySigningKey = []byte("AllYourBase")

func GenerateToken(id string) string {

	// Create the Claims
	claims := AuthClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 600,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v\n", ss, err)
	return ss

	// res, err := parseToken(ss)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%v %v", res.ID, res.StandardClaims.ExpiresAt)
}

func ParseToken(tokenString string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
