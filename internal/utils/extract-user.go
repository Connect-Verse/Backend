package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ExtractUser(ctx *gin.Context) (string,bool){
	
	userId,ok :=ctx.Get("id")
	
	Id, ok := userId.(string)
	
	if !ok{
	fmt.Print("unable to assert the string type for this")// handle type assertion failure
	}
	
	return Id,ok
}