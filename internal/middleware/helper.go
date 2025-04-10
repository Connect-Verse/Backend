package middleware

import (
	"fmt"
	"net/http"

	"github.com/connect-verse/internal/data/response"
	"github.com/connect-verse/internal/utils"
	"github.com/gin-gonic/gin"
)

func Middleware(ctx *gin.Context){
 
	cookie,err:=ctx.Cookie("token")

	if err!=nil{
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{Code: 400, Message: "cookies is not found try signing in again", Err: err.Error()})
	    return
	}

	data,err:= utils.VerifyToken(cookie)

	if err!=nil{

		ctx.JSON(http.StatusForbidden, response.ErrorResponse{Code: 400, Message: "unable to parse the provided cookies", Err: err.Error()})
		return
	}
	fmt.Print(data,"eror  occured")

	ctx.Set("name",data.Name)
	ctx.Set("email",data.Email)
	ctx.Set("id",data.Id)
    fmt.Print(data.Id)
	ctx.Next()

}

func CorsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE,HEAD")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}