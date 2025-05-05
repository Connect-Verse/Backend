package position

import (
	"fmt"

	"github.com/connect-verse/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PositionRepoImpl struct{
	db *gorm.DB
}

func NewPosRepoImpl(db *gorm.DB) *PositionRepoImpl{
	return &PositionRepoImpl{db :db}
}

func (p *PositionRepoImpl) SetPosition(position models.PlayerPosition) (models.PlayerPosition, error){
   fmt.Print(position,"pehlea ")
   result:= p.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&position)

   if result.Error!=nil {
		return models.PlayerPosition{},result.Error
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

