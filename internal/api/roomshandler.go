package handlers

import (
	"net/http"

	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/gin-gonic/gin"
)

type roomId struct{
	Id string `json:"id"`
}


func (c *Controller) CreateRoom(ctx *gin.Context) {
	req := request.RoomRequest{}
	ctx.ShouldBindJSON(&req)

	result, err := c.roomService.CreateRoom(req)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{
			Code:    400,
			Message: "internal server error occured while generating Room",
			Err:     err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, response.RoomResponse{
		Id:        result.Id,
		Name:      result.Name,
		CreatedBy: result.CreatedBy,
		MapId:     result.MapId,
		Maps:      result.Maps,
	})

}

func (c *Controller) DeleteRoom(ctx *gin.Context) {
	req := roomId{}
	ctx.ShouldBindJSON(&req)

	result, err := c.roomService.DeleteRoom(req.Id)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{
			Code:    400,
			Message: "internal server error occured while generating Room",
			Err:     err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, response.RoomResponse{
		Id:        result.Id,
		Name:      result.Name,
		CreatedBy: result.CreatedBy,
		MapId:     result.MapId,
		Maps:      result.Maps,
	})

}

func (c *Controller) MyRoom(ctx *gin.Context) {
	req:= roomId{}
	ctx.ShouldBindJSON(&req)

	result, err := c.roomService.MyRoom(req.Id)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{
			Code:    400,
			Message: "internal server error occured while generating Room",
			Err:     err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, result)

}

func (c *Controller) FindAllRooms(ctx *gin.Context) {
	req := request.RoomRequest{}
	ctx.ShouldBindJSON(&req)

	result, err := c.roomService.FindAllRooms()

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{
			Code:    400,
			Message: "internal server error occured while generating Room",
			Err:     err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, result)

}
