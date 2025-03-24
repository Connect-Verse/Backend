package avatarservice

import (
	"errors"

	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/repository/avatars"
	"github.com/go-playground/validator/v10"
)

type AvatarServImpl struct{
	AvatarRepo   avatars.AvatarsRepository
	Validate     *validator.Validate
}

func NewAvatarServImpl(avatarRepo avatars.AvatarsRepository, validate *validator.Validate) (*AvatarServImpl,error) {
     if validate==nil {
		return nil,errors.New("no validate is provided")
	 }
	 return &AvatarServImpl{
		AvatarRepo: avatarRepo,
		Validate: validate,
	 },nil
}

func (m *AvatarServImpl) CreateAvatar(avatar request.AvatarRequest) (response.AvatarResponse,error){
	result,err:= m.AvatarRepo.CreateAvatar(models.Avatars{
	   Image: avatar.Image,
	   Name: avatar.Name,
	})

	if err!=nil {
	   return response.AvatarResponse{},err
	}

	return response.AvatarResponse{
	   Id: result.Id,
	   Image: result.Image,
	   Name: result.Name,
	   ExistedFrom: result.ExistedFrom,
	},nil
 }


 func (m *AvatarServImpl) DeleteAvatar(avatarId string) (response.AvatarResponse,error){
   result,err:= m.AvatarRepo.DeleteAvatar(avatarId)

	if err!=nil {
	   return response.AvatarResponse{},err
	}

	return response.AvatarResponse{
	   Id: result.Id,
	   Image: result.Image,
	   Name: result.Name,
	   ExistedFrom: result.ExistedFrom,
	},nil
 }


 func (m *AvatarServImpl) FindAvatar(avatarId string) (response.AvatarResponse,error){
   result,err:= m.AvatarRepo.FindAvatar(avatarId)

	if err!=nil {
	   return response.AvatarResponse{},err
	}

	return response.AvatarResponse{
	   Id: result.Id,
	   Image: result.Image,
	   Name: result.Name,
	   ExistedFrom: result.ExistedFrom,
	},nil
 }


 func (m *AvatarServImpl) FindAllAvatars() ([]response.AvatarResponse,error){
   result,err:= m.AvatarRepo.FindAllAvatar()
   avatars := []response.AvatarResponse{}
	if err!=nil {
	   return avatars,err
	}
	for _,value := range result {
	   avatar:=response.AvatarResponse{
		   Id: value.Id,
		   Image: value.Image,
		   Name: value.Name,
		   ExistedFrom: value.ExistedFrom,
		}
	   avatars= append(avatars,avatar)
	}
	return avatars,nil
 }


 func (m *AvatarServImpl) UpdateAvatars(avatar request.AvatarRequest) (response.AvatarResponse,error){
   result,err:= m.AvatarRepo.UpdateAvatar(models.Avatars{
	   Image: avatar.Image,
	   Name: avatar.Name,
	   ExistedFrom: avatar.ExistedFrom,
   })
   
	if err!=nil {
	   return response.AvatarResponse{},err
	}

	return response.AvatarResponse{
	   Id: result.Id,
	   Image: result.Image,
	   ExistedFrom: result.ExistedFrom,
	   Name: result.Name,
	},nil
 }




