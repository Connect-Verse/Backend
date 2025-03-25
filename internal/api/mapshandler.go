package handlers

import (
	"net/http"

	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateMap(ctx *gin.Context) {

	req := request.MapRequest{}
	ctx.ShouldBindJSON(&req)

	result, err := c.mapService.CreateMap(req)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{
			Code:    400,
			Message: "internal server error occured while generating map",
			Err:     err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, response.MapResponse{
		Id:    result.Id,
		Tiles: result.Tiles,
		Info:  result.Info,
		Image: result.Info,
	})

}

func (c *Controller) DeleteMap(ctx *gin.Context) {
	req := request.MapRequest{}
	ctx.ShouldBindJSON(&req)

	result, err := c.mapService.CreateMap(req)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{
			Code:    400,
			Message: "internal server error occured while generating map",
			Err:     err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, response.MapResponse{
		Id:    result.Id,
		Tiles: result.Tiles,
		Info:  result.Info,
		Image: result.Info,
	})

}

func (c *Controller) UpdateMap(ctx *gin.Context) {
	req := request.MapRequest{}
	ctx.ShouldBindJSON(&req)

	result, err := c.mapService.CreateMap(req)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{
			Code:    400,
			Message: "internal server error occured while generating map",
			Err:     err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, response.MapResponse{
		Id:    result.Id,
		Tiles: result.Tiles,
		Info:  result.Info,
		Image: result.Info,
	})

}

func (c *Controller) FindMap(ctx *gin.Context) {
	req := request.MapRequest{}
	ctx.ShouldBindJSON(&req)

	result, err := c.mapService.CreateMap(req)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{
			Code:    400,
			Message: "internal server error occured while generating map",
			Err:     err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, response.MapResponse{
		Id:    result.Id,
		Tiles: result.Tiles,
		Info:  result.Info,
		Image: result.Info,
	})

}

func (c *Controller) FindAllMap(ctx *gin.Context) {
	req := request.MapRequest{}
	ctx.ShouldBindJSON(&req)

	result, err := c.mapService.CreateMap(req)

	if err != nil {
		ctx.JSON(http.StatusForbidden, response.ErrorResponse{
			Code:    400,
			Message: "internal server error occured while generating map",
			Err:     err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, response.MapResponse{
		Id:    result.Id,
		Tiles: result.Tiles,
		Info:  result.Info,
		Image: result.Info,
	})

}
