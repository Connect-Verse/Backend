package maps

import (
	"github.com/connect-verse/internal/models"
	"gorm.io/gorm"
)

type MapsRepoImpl struct {
  db *gorm.DB
}

func NewMapsRepoImpl(db *gorm.DB) *MapsRepoImpl {
	return &MapsRepoImpl{db: db}
}

func (u *MapsRepoImpl) FindAllMaps() ([]models.Maps, error){
	var mapp []models.Maps
    result:= u.db.Find(&mapp)
	if result.Error != nil{
		return nil, result.Error
	}        

	return mapp , nil
}


func (u *MapsRepoImpl) CreateMap(mapp models.Maps) (models.Maps,error){
	
    result:= u.db.Create(&mapp)
	if result.Error != nil{
		return  models.Maps{},result.Error
	}        

	return  mapp,nil
}


func (u *MapsRepoImpl) FindMap(mappId string) (models.Maps,error){
    var mapp models.Maps
	result:= u.db.Where("id=?",mappId).First(&mapp)
	if result.Error != nil{
		return models.Maps{}, result.Error
	}        
	return mapp , nil
}



func (u *MapsRepoImpl) UpdateMap(updMap models.Maps) (models.Maps, error){
	var mapp models.Maps
    result:= u.db.Model(&mapp).Updates(updMap)
	if result.Error != nil{
		return models.Maps{},result.Error
	}        

	return mapp ,nil
}



func (u *MapsRepoImpl) DeleteMap(mappId string) (models.Maps, error){
    var mapp models.Maps
    result:= u.db.Where("id=?",mappId).Delete(&mapp)
	if result.Error != nil{
		return models.Maps{},result.Error
	}        
	return mapp,nil
}



