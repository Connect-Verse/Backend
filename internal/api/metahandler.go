package handlers

import (
	"log"
	"net/http"

	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/connect-verse/internal/utils"
	"github.com/gin-gonic/gin"
)


func (c *Controller) CreateMeta(ctx *gin.Context){
  req:= request.MetaUser{}
  ctx.ShouldBindJSON(&req);
  id,ok:=utils.ExtractUser(ctx)
  if !ok{
	ctx.JSON(http.StatusBadRequest,response.ErrorResponse{
		Code:400,
		Message: "server error occurred",
		Err: "unable to extract user in the meta created usee",
		
	})
  }
  req.UserId=id

  result,err:=c.metaService.CreateMeta(req)
  
  if err!=nil{
	ctx.JSON(http.StatusBadRequest,response.ErrorResponse{
		Code:400,
		Message: "server error occurred",
		Err: err.Error(),
	})
	return
  }

  ctx.JSON(http.StatusOK,result)
}

type metaId struct{
	Id string `json:"id"`
}

func (c *Controller) DeleteMetaUser(ctx *gin.Context){
	var req metaId
	ctx.ShouldBindJSON(&req);
    log.Print(req)
	err:=c.metaService.DeleteMetaUser(req.Id)
  
	if err!=nil{
	  ctx.JSON(http.StatusBadRequest,response.ErrorResponse{
		  Code:400,
		  Message: "server error occurred",
		  Err: err.Error(),
	  })
	  return
	}
  
	ctx.JSON(http.StatusOK,response.Response{
		Code: 200,
		Status:"deleted successfully",
	})
  }

  func (c *Controller) FindById(ctx *gin.Context){
	req:= metaId{}
	ctx.ShouldBindJSON(&req);
  
	result,err:=c.metaService.FindById(req.Id)
  
	if err!=nil{
	  ctx.JSON(http.StatusBadRequest,response.ErrorResponse{
		  Code:400,
		  Message: "server error occurred",
		  Err: err.Error(),
	  })
	  return
	}
  
	ctx.JSON(http.StatusOK,response.Response{
		Code: 200,
		Status: "found successfully",
		Data: result,
	})
  }
