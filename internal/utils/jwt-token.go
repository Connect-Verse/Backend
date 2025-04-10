package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("notSoSecretKey")


func GenerateToken(id string,email string, name string) (string,error){
	 token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"email":      email,
		"id":         id,
		"name":       name,
	})
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

 return tokenString, nil
}

type middlewareData struct{
	Email string
	Id string
	Name string
}

func VerifyToken(tokenString string) (middlewareData,error){

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
	  }
	  return sampleSecretKey, nil

		})

		if err != nil {
			return middlewareData{}, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
          if ok && parsedToken.Valid  {
                username := claims["name"].(string)
				email := claims["email"].(string)
                id := claims["id"].(string)

                return middlewareData{
					Name:username,
					Email:email,
					Id:id,
				}, nil
          }

        
        return middlewareData{}, nil
}

