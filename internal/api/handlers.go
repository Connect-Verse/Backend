package handlers

import (
	"crypto/sha256"
	"database/sql"
	"net/http"
	"encoding/hex"
	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/services"
	"github.com/connect-verse/internal/services/auth-service"
	"github.com/connect-verse/internal/utils"
	"github.com/gin-gonic/gin"
)



type Controller struct{
	userService  services.UserServices
    authService  authservice.AuthService
}

func NewControllerService(userService services.UserServices, authService authservice.AuthService) *Controller {
  return &Controller{
	authService: authService,
	userService: userService,
  }
}


func (c *Controller) Check(ctx *gin.Context){
	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   models.User{
			Email:"siofjioe",
			Id:"oisjfoiewi",
		},
	   }
	   
	   ctx.JSON(http.StatusOK, res)
}


func (c *Controller) Login(ctx *gin.Context){
	req := request.CreateUserRequest{}
	ctx.ShouldBindJSON(&req)

	
}

func (c *Controller) Signup(ctx *gin.Context){
	req := struct{
    Name     string `validate:"required,min=1,max=200" json:"name"`
	Email    string `validate:"required,min=1,max=50" json:"email"`
    Password string `validate:"required,min=1,max=20" json:"password"`
	}{}
	ctx.ShouldBindJSON(&req)
	h:= sha256.New()
	h.Write([]byte(req.Email))
	br:=h.Sum(nil)
	hashedPassword,err:= utils.HashPassword(req.Password)
	
	if err!=nil{
		ctx.JSON(http.StatusForbidden , response.ErrorResponse{Code:400,Message:"error occurred while creating the account"})
	}


	requiredReq:= request.CreateUserRequest{
		Name: sql.NullString{String:req.Name, Valid: true},
		Email: req.Email,
		Password: hashedPassword,
	}
	err= c.authService.SignUp(requiredReq,hex.EncodeToString(br))

	if err!=nil{
		ctx.JSON(http.StatusForbidden , response.ErrorResponse{Code:400,Message:"error occurred while creating the account"})
	}
    
	ctx.JSON(http.StatusOK, response.Response{
		Code :  200,
    Status :"ok",
    Data  : struct{
		email string
	}{
      email:req.Email,
	},
	})

}