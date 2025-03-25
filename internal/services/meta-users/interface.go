package metaservice

import (
	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
)

type MetaService interface{
	CreateMeta(metaUser request.MetaUser) (response.MetaUserResponse,error)
	DeleteMetaUser(metaUserId string) (error)
	FindById(metaUserId string) (response.MetaUserResponse,error)
}