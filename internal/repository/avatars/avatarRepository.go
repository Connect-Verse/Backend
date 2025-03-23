package avatars

import (
	"github.com/connect-verse/internal/models"
	"gorm.io/gorm"
)

type AvatarRepoImpl struct {
  db *gorm.DB
}

func NewAvatarRepoImpl(db *gorm.DB) *AvatarRepoImpl {
	return &AvatarRepoImpl{db: db}
}

func (u *AvatarRepoImpl) FindAllAvatar() ([]models.Avatars, error){
	var avatar []models.Avatars
    result:= u.db.Find(&avatar)
	if result.Error != nil{
		return nil, result.Error
	}        

	return avatar , nil
}


func (u *AvatarRepoImpl) CreateAvatar(avatar models.Avatars) (models.Avatars,error){
	
    result:= u.db.Create(&avatar)
	if result.Error != nil{
		return  models.Avatars{},result.Error
	}        

	return  avatar,nil
}


func (u *AvatarRepoImpl) FindAvatar(avatarId string) (models.Avatars,error){
    var avatar models.Avatars
	result:= u.db.Where("id=?",avatarId).First(&avatar)
	if result.Error != nil{
		return models.Avatars{}, result.Error
	}        
	return avatar , nil
}



func (u *AvatarRepoImpl) UpdateAvatar(updMap models.Avatars) (models.Avatars, error){
	var avatar models.Avatars
    result:= u.db.Model(&avatar).Updates(updMap)
	if result.Error != nil{
		return models.Avatars{},result.Error
	}        

	return avatar ,nil
}



func (u *AvatarRepoImpl) DeleteAvatar(avatarId string) (models.Avatars, error){
    var avatar models.Avatars
    result:= u.db.Where("id=?",avatarId).Delete(&avatar)
	if result.Error != nil{
		return models.Avatars{},result.Error
	}        
	return avatar,nil
}



