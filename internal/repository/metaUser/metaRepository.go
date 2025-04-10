package metauser

import (
	"log"

	"github.com/connect-verse/internal/models"
	"gorm.io/gorm"
)

type MetaRepoImpl struct {
	db *gorm.DB
}

func NewMetaRepoImpl(db *gorm.DB) *MetaRepoImpl{
	return &MetaRepoImpl{db :db}
}


//this is same as joining a room
func (m *MetaRepoImpl) CreateMeta(metaUser models.MetaUsers) (models.MetaUsers,error){
  result := m.db.Create(&metaUser)

  if result.Error!=nil {
    return models.MetaUsers{},result.Error
  }

  return metaUser,nil
}

func (m *MetaRepoImpl) FindById(metaUserId string) (models.MetaUsers,error){
	var metaUser models.MetaUsers
	result := m.db.Preload("Room").Preload("UserAvatar").Preload("Position").Preload("Avatars").Where("id = ?",metaUserId).Find(&metaUser)
    log.Printf(metaUser.Id,metaUserId)
	if result.Error!=nil {
	  return models.MetaUsers{},result.Error
	}
  
	return metaUser,nil
}


//this is same as leaving a room
func (m *MetaRepoImpl) DeleteMetaUser(metaUserId string) (error){
	result := m.db.Where("id=?",metaUserId).Delete(&models.MetaUsers{})

	if result.Error!=nil {
	  return result.Error
	}
  
	return nil
}