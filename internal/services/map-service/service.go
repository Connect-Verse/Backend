package mapservice

import (
	"errors"
	"fmt"

	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/repository/maps"
	"github.com/go-playground/validator/v10"
)

type MapServiceImpl struct{
	mapRepo     maps.MapsRepository
	validate    *validator.Validate
}

func NewMapServiceImpl(mapService maps.MapsRepository, validate *validator.Validate) (*MapServiceImpl,error){
	if validate==nil{
        return &MapServiceImpl{},errors.New("no validator is provided")
	}

	return &MapServiceImpl{
		mapRepo: mapService,
		validate: validate,
	},nil
}


  func (m *MapServiceImpl) CreateMap(mapp request.MapRequest) (response.MapResponse,error){
    
	err:= m.validate.Struct(mapp)
	if err!=nil{
		return response.MapResponse{},errors.New("the data provided is not matching with the schema")
	}
	
	result,err:= m.mapRepo.CreateMap(models.Maps{
		Image: mapp.Image,
		Tiles: mapp.Tiles,
		Info: mapp.Info,
	 })

	 if err!=nil {
		return response.MapResponse{},err
	 }

	 return response.MapResponse{
		Id: result.Id,
		Image: result.Image,
		Info: result.Info,
		Tiles: result.Tiles,
	 },nil
  }


  func (m *MapServiceImpl) DeleteMap(mappId string) (response.MapResponse,error){
	result,err:= m.mapRepo.DeleteMap(mappId)

	 if err!=nil {
		return response.MapResponse{},err
	 }

	 return response.MapResponse{
		Id: result.Id,
		Image: result.Image,
		Info: result.Info,
		Tiles: result.Tiles,
	 },nil
  }


  func (m *MapServiceImpl) FindMap(mappId string) (response.MapResponse,error){
	result,err:= m.mapRepo.FindMap(mappId)

	 if err!=nil {
		return response.MapResponse{},err
	 }

	 return response.MapResponse{
		Id: result.Id,
		Image: result.Image,
		Info: result.Info,
		Tiles: result.Tiles,
	 },nil
  }


  func (m *MapServiceImpl) FindALlMaps() ([]response.MapResponse,error){
	result,err:= m.mapRepo.FindAllMaps()
	mapps:=[]response.MapResponse{}
	 if err!=nil {
		return mapps,err
	 }
	 for _,value:= range result {
		mapp:=response.MapResponse{
			Id: value.Id,
			Image: value.Image,
			Info: value.Info,
			Tiles: value.Tiles,
		 }
		mapps= append(mapps,mapp)
	 }
	 fmt.Print(mapps)
	 return mapps,nil
  }


  func (m *MapServiceImpl) UpdateMaps(mapp request.MapRequest) (response.MapResponse,error){
	result,err:= m.mapRepo.UpdateMap(models.Maps{
		Image: mapp.Image,
		Info: mapp.Info,
		Tiles: mapp.Tiles,
	})
	
	 if err!=nil {
		return response.MapResponse{},err
	 }

	 return response.MapResponse{
		Id: result.Id,
		Image: result.Image,
		Tiles: result.Tiles,
		Info: result.Info,
	 },nil
  }

