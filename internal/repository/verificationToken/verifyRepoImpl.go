package verificationtoken

import (
	"github.com/connect-verse/internal/models"
	"gorm.io/gorm"
)

type VerifyImplementation struct{
	db *gorm.DB
}

func NewVerifyImplementation(DB *gorm.DB) *VerifyImplementation {
  return &VerifyImplementation{db : DB}
}


func (u *VerifyImplementation) FindbyEmail(email string) (users models.VerificationToken,err error){
	var token models.VerificationToken
    result:= u.db.Find(&token).Where("email_identifier=?",email)
	if result.Error != nil{
		return models.VerificationToken{}, result.Error
	}        

	return token , nil
}



func (u *VerifyImplementation) Create(token models.VerificationToken) (err error){
    result:= u.db.Create(&token)
	if result.Error != nil{
		return  result.Error
	}        

	return  nil
}


func (u *VerifyImplementation) FindbyId(tokenId string) ( models.VerificationToken, error){
    var token models.VerificationToken
	result:= u.db.First(&token, tokenId)
	if result.Error != nil{
		return models.VerificationToken{}, result.Error
	}        

	return token , nil
}


func (u *VerifyImplementation) Update(upd models.VerificationToken) ( models.VerificationToken ,  error){
	var token models.VerificationToken
    result:= u.db.Model(&token).Updates(upd)
	if result.Error != nil{
		return models.VerificationToken{},result.Error
	}        

	return token ,nil
}



func (u *VerifyImplementation) Delete(tokenId string) (error){
    var token models.VerificationToken
    result:= u.db.Delete(&token,tokenId)
	if result.Error != nil{
		return result.Error
	}        

	return nil
}


