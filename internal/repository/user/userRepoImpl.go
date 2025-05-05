package user

import (
	"fmt"

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


func (u *UserImplementation) Create(user models.User) ( models.User,error){
	
    result:= u.db.Create(&user)
	if result.Error != nil{
		fmt.Print(result.Error)
		return  models.User{},result.Error
	}        

	return  user,nil
}


func (u *UserImplementation) FindbyId(userId string) (users models.User,err error){
    var user models.User
	result:= u.db.Where("Id=?", userId).Find(&user)
	if result.Error != nil{
		return models.User{}, result.Error
	}        

	return user , nil
}

func (u *UserImplementation) FindbyEmail(userEmail string) (users models.User,err error){
    var user models.User
	fmt.Print(userEmail,"email is ")
	result:= u.db.Preload("JoinedRooms").Where("email=?",userEmail).Find(&user)
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


