package handlers

import (
	"net/http"
	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateAvatar(ctx *gin.Context ) {
	req:=request.AvatarRequest{}  
	ctx.ShouldBindJSON(&req)
    
	result,err:= c.avatarService.CreateAvatar(req)

	if err!=nil {
		 ctx.JSON(http.StatusForbidden,response.ErrorResponse{
			Code: 400,
			Message: "internal server error occured while generating Avatar",
			Err:err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,response.AvatarResponse{
		Id: result.Id,
		Name: result.Name,
        Image: result.Image,
    	 ExistedFrom: result.ExistedFrom, 
	})

}

func (c *Controller) DeleteAvatar(ctx *gin.Context ) {
	req:=request.AvatarRequest{}  
	ctx.ShouldBindJSON(&req)
    
	result,err:= c.avatarService.CreateAvatar(req)

	if err!=nil {
		 ctx.JSON(http.StatusForbidden,response.ErrorResponse{
			Code: 400,
			Message: "internal server error occured while generating Avatar",
			Err:err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,response.AvatarResponse{
		Id: result.Id,
		Name: result.Name,
        Image: result.Image,
    	 ExistedFrom: result.ExistedFrom, 
	})

}

func (c *Controller) UpdateAvatar(ctx *gin.Context ) {
	req:=request.AvatarRequest{}  
	ctx.ShouldBindJSON(&req)
    
	result,err:= c.avatarService.CreateAvatar(req)

	if err!=nil {
		 ctx.JSON(http.StatusForbidden,response.ErrorResponse{
			Code: 400,
			Message: "internal server error occured while generating Avatar",
			Err:err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,response.AvatarResponse{
		Id: result.Id,
		Name: result.Name,
        Image: result.Image,
    	 ExistedFrom: result.ExistedFrom, 
	})

}

func (c *Controller) FindAvatar(ctx *gin.Context ) {
	req:=request.AvatarRequest{}  
	ctx.ShouldBindJSON(&req)
    
	result,err:= c.avatarService.CreateAvatar(req)

	if err!=nil {
		 ctx.JSON(http.StatusForbidden,response.ErrorResponse{
			Code: 400,
			Message: "internal server error occured while generating Avatar",
			Err:err.Error(),
		})
		
	}

	ctx.JSON(http.StatusOK,response.AvatarResponse{
		Id: result.Id,
		Name: result.Name,
        Image: result.Image,
    	 ExistedFrom: result.ExistedFrom, 
	})

}


func (c *Controller) FindAllAvatar(ctx *gin.Context ) {
	req:=request.AvatarRequest{}  
	ctx.ShouldBindJSON(&req)
    
	result,err:= c.avatarService.CreateAvatar(req)

	if err!=nil {
		 ctx.JSON(http.StatusForbidden,response.ErrorResponse{
			Code: 400,
			Message: "internal server error occured while generating Avatar",
			Err:err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,response.AvatarResponse{
		Id: result.Id,
		Name: result.Name,
        Image: result.Image,
    	 ExistedFrom: result.ExistedFrom, 
	})

}
