package utils

import (

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(pass string) (string,error){
   hashed,err:= bcrypt.GenerateFromPassword([]byte(pass),14)
   if err!=nil {
	return string(hashed),err
   } 
   return string(hashed),nil
}

func ComparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}