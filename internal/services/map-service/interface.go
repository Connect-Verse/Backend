package mapservice

import (
	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
)

type MapService interface{
	CreateMap(mapp request.MapRequest) (response.MapResponse,error)
	DeleteMap(mappId string) (response.MapResponse,error)
	FindMap(mappId string) (response.MapResponse,error)
	FindALlMaps() ([]response.MapResponse,error)
	UpdateMaps(mapp request.MapRequest) (response.MapResponse,error)

}