package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"net/http"
	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/services"
	"github.com/connect-verse/internal/services/auth-service"
	"github.com/connect-verse/internal/services/avatar-service"
	"github.com/connect-verse/internal/services/map-service"
	"github.com/connect-verse/internal/utils"
	"github.com/gin-gonic/gin"
)



type Controller struct{
	userService  services.UserServices
    authService  authservice.AuthService
	// roomService  roomService.RoomService
	mapService   mapservice.MapService
	avatarService avatarservice.AvatarService
}

func NewControllerService(avatarService avatarservice.AvatarService,userService services.UserServices, authService authservice.AuthService, mapService mapservice.MapService) *Controller {
  return &Controller{
	authService: authService,
	userService: userService,
	mapService : mapService,
	avatarService: avatarService,
  }
}


func (c *Controller) Check(ctx *gin.Context){
	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   models.User{
			Email:"siofjioe",
		},
	   }
	   
	   ctx.JSON(http.StatusOK, res)
}


func (c *Controller) Login(ctx *gin.Context){
	req := request.CreateUserRequest{}
	ctx.ShouldBindJSON(&req)

	result,err:=c.authService.Login(req)

	if err!=nil {
		ctx.JSON(http.StatusForbidden , response.ErrorResponse{Code:400,Message:"error occurred while creating the account"+err.Error()})
	}
    
	isMatching:=utils.ComparePassword(req.Password,result.Password)
    
	if isMatching {
		ctx.JSON(http.StatusOK , response.ErrorResponse{Code:200,Message:"succesfully logged in"})
	}else{
		ctx.JSON(http.StatusForbidden , response.ErrorResponse{Code:400,Message:"the password provided is not matching to the previous save password"})
	}

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

//jwt and session creation is to be managed

func (c *Controller) Verify(ctx *gin.Context) {
	req:= struct{
		VerificationId string  `json:"tokenId"`
		EmailIdentifier string  `json:"email"`
	}{}
	ctx.ShouldBindJSON(&req)
	err:=c.authService.Verify(req.EmailIdentifier,req.VerificationId)
   
    if err!=nil {
		ctx.JSON(http.StatusForbidden , response.ErrorResponse{Code:400,Message:"error occurred while verifying the account"})
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code :  200,
    Status :"ok",
    Data  : struct{
		email string
	}{
      email:req.EmailIdentifier,
	},
	})


}