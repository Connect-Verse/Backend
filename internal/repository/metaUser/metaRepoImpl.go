package metauser

import "github.com/connect-verse/internal/models"


type MetaRepository interface{
	CreateMeta(metaUser models.MetaUsers) (models.MetaUsers,error)
	DeleteMetaUser(metaUserId string) (error)
	FindById(metaUserId string) (models.MetaUsers,error)
}