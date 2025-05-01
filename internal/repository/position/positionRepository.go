package position

import (
	"github.com/connect-verse/internal/models"
	"gorm.io/gorm"
)

type PositionRepoImpl struct{
	db *gorm.DB
}

func NewPosRepoImpl(db *gorm.DB) *PositionRepoImpl{
	return &PositionRepoImpl{db :db}
}

func (p *PositionRepoImpl) SetPosition(position models.PlayerPosition) (models.PlayerPosition, error){
   err:= p.db.Create(position)
   if err!=nil{
	return models.PlayerPosition{},err.Error
   }
   return position,nil
}

func (p *PositionRepoImpl) FindPosition(metaId string) (models.PlayerPosition,error){
   var playerPosition models.PlayerPosition
   err := p.db.Where("meta_users_id",metaId).Find(&playerPosition)
   if err!=nil {
     return models.PlayerPosition{},err.Error
   }
   return playerPosition,nil
}

