package user

import (
	

	"github.com/connect-verse/internal/models"
	"gorm.io/gorm"
)

type UserImplementation struct{
	db *gorm.DB
}

func NewUserImplementation(DB *gorm.DB) *UserImplementation {
  return &UserImplementation{db : DB}
}


func (u *UserImplementation) FindAll() (users []models.User,err error){
	var user []models.User
    result:= u.db.Find(&user)
	if result.Error != nil{
		return nil, result.Error
	}        

	return user , nil
}


func (u *UserImplementation) Create(user models.User) (err error){
	
    result:= u.db.Create(&user)
	if result.Error != nil{
		return  result.Error
	}        

	return  nil
}


func (u *UserImplementation) FindbyId(userId string) (users models.User,err error){
    var user models.User
	result:= u.db.First(&user, userId)
	if result.Error != nil{
		return models.User{}, result.Error
	}        

	return user , nil
}

func (u *UserImplementation) FindbyEmail(userEmail string) (users models.User,err error){
    var user models.User
	result:= u.db.Find(&user).Where("email=?",userEmail)
	if result.Error != nil{
		return models.User{}, result.Error
	}        
	return user,nil
}

func (u *UserImplementation) Update(upd models.User) (users models.User , err error){
	var user models.User
    result:= u.db.Model(&user).Updates(upd)
	if result.Error != nil{
		return models.User{},result.Error
	}        

	return user ,nil
}



func (u *UserImplementation) Delete(userId string) (err error){
    var user models.User
    result:= u.db.Delete(&user,userId)
	if result.Error != nil{
		return result.Error
	}        

	return nil
}


