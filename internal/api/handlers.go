package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/services"
	"github.com/connect-verse/internal/services/auth-service"
	"github.com/connect-verse/internal/services/avatar-service"
	"github.com/connect-verse/internal/services/map-service"
	"github.com/connect-verse/internal/services/meta-users"
	"github.com/connect-verse/internal/services/room-service"
	"github.com/connect-verse/internal/utils"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService   services.UserServices
	authService   authservice.AuthService
	roomService   roomservice.RoomService
	mapService    mapservice.MapService
	avatarService avatarservice.AvatarService
	metaService   metaservice.MetaService
}

func NewControllerService(metaService metaservice.MetaService,roomService roomservice.RoomService, avatarService avatarservice.AvatarService, userService services.UserServices, authService authservice.AuthService, mapService mapservice.MapService) *Controller {
	return &Controller{
		authService:   authService,
		userService:   userService,
		mapService:    mapService,
		avatarService: avatarService,
		roomService:   roomService,
		metaService:   metaService,
	}
}

func (c *Controller) Check(ctx *gin.Context) {
	res := response.Response{
		Code:   200,
		Status: "Ok",
		Data: models.User{
			Email: "siofjioe",
		},
	}

	ctx.JSON(http.StatusOK, res)
}


func (c *Controller) Login(ctx *gin.Context) {
	req := request.CreateUserRequest{}
	ctx.ShouldBindJSON(&req)

	result, err := c.authService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{Code: 400, Message: "error occurred while creating the account" + err.Error(), Err: err.Error()})
	}

	isMatching := utils.ComparePassword(req.Password, result.Password)


	if isMatching==nil {

		fmt.Print("logging the error",result,req)

		
		token,err := utils.GenerateToken(result.Id,result.Email,result.Name.String)

		if err!=nil{
			ctx.JSON(http.StatusForbidden, response.ErrorResponse{Code: 400, Message: "error occurred while creating the account", Err: err.Error()})
		}
	
		ctx.SetCookie("token", token, 360000, "/", "localhost", false, true)

		
		ctx.JSON(http.StatusOK, response.ErrorResponse{Code: 200, Message: "succesfully logged in"})
	} else {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{Code: 400, Message: "the password provided is not matching to the previous save password", Err: err.Error()})
	}

}

func (c *Controller) Signup(ctx *gin.Context) {
	req := struct {
		Name     string `validate:"required,min=1,max=200" json:"name"`
		Email    string `validate:"required,min=1,max=50" json:"email"`
		Password string `validate:"required,min=1,max=20" json:"password"`
	}{}
	ctx.ShouldBindJSON(&req)

	h := sha256.New()
	h.Write([]byte(req.Email))
	br := h.Sum(nil)

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{Code: 400, Message: "error occurred while creating the account", Err: err.Error()})
	}

	requiredReq := request.CreateUserRequest{
		Name:     sql.NullString{String: req.Name, Valid: true},
		Email:    req.Email,
		Password: hashedPassword,
	}

	result,err := c.authService.SignUp(requiredReq, hex.EncodeToString(br))

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{Code: 400, Message: "error occurred while creating the account", Err: err.Error()})
	}

	token,err := utils.GenerateToken(result.Id,result.Email,result.Name.String)

	if err!=nil{
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{Code: 400, Message: "error occurred while creating the account", Err: err.Error()})
	}

	ctx.SetCookie("token", token, 360000, "/", "localhost", false, true)


	ctx.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "ok",
		Data: struct {
			email string
		}{
			email: req.Email,
		},
	})

}

//jwt and session creation is to be managed

func (c *Controller) Verify(ctx *gin.Context) {
	req := struct {
		VerificationId  string `json:"tokenId"`
		EmailIdentifier string `json:"email"`
	}{}
	ctx.ShouldBindJSON(&req)
	err := c.authService.Verify(req.EmailIdentifier, req.VerificationId)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{Code: 400, Message: "error occurred while verifying the account", Err: err.Error()})
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "ok",
		Data: struct {
			email string
		}{
			email: req.EmailIdentifier,
		},
	})

}


func (c *Controller) FindByEmail (ctx *gin.Context){

	userEmail,ok :=ctx.Get("email")
	
	email, ok := userEmail.(string)
	
	if !ok{
	fmt.Print("unable to assert the string type for this")// handle type assertion failure
	}    
	result,err:=c.userService.FindbyEmail(email)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{Code: 400, Message: "error occurred while finding the users", Err: err.Error()})
	}

	ctx.JSON(http.StatusOK,result)

}


func (c *Controller) Logout(ctx *gin.Context){
	ctx.SetCookie("token", "", 360000, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK,response.Response{
	Code:http.StatusAccepted,
	Status:"200",
	})
}