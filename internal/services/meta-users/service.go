package metaservice

import (
	"errors"
	"log"

	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/repository/metaUser"
	"github.com/go-playground/validator/v10"
)


type MetaServiceImpl struct{
	metaRepo   metauser.MetaRepository
	validate   *validator.Validate
}
func NewMetaServiceImpl(metaRepo metauser.MetaRepository, validate *validator.Validate) (*MetaServiceImpl,error) {
  if validate==nil{
	 return &MetaServiceImpl{},errors.New("validator is not mounted yet")
  }
  return &MetaServiceImpl{
	metaRepo: metaRepo,
	validate: validate,
  },nil
}


func (m *MetaServiceImpl) CreateMeta(metaUser request.MetaUser) (response.MetaUserResponse,error){
	log.Printf(metaUser.RoomId,metaUser.Name,metaUser.UserAvatarId,metaUser.UserId)
	result,err:= m.metaRepo.CreateMeta(models.MetaUsers{
		Name: metaUser.Name,
        UserAvatarId: metaUser.UserAvatarId,
		RoomId: metaUser.RoomId,
		UserId: metaUser.UserId,
	})
    if err!=nil {
		return response.MetaUserResponse{},err
	}

	return response.MetaUserResponse{
		Id           : result.Id,
		Name         : result.Name,
		UserAvatarId : result.UserAvatarId,
		UserId       : result.UserId,
		RoomId       : result.Id,
		Room         : result.Id,
		Position     : result.Id,
		Avatar       : response.AvatarResponse{
						Id: result.UserAvatar.Id,
						Image: result.UserAvatar.Image,
						Name: result.UserAvatar.Name,
						},
		
	},nil
}

func (m *MetaServiceImpl) DeleteMetaUser(metaUserId string) (error){
	err:= m.metaRepo.DeleteMetaUser(metaUserId)
    if err!=nil {
		return err
	}

	return nil
}

func (m *MetaServiceImpl) FindById(metaUserId string) (response.MetaUserResponse,error){
	result,err:= m.metaRepo.FindById(metaUserId)
    if err!=nil {
		return response.MetaUserResponse{},err
	}

	return response.MetaUserResponse{
		Id           : result.Id,
		Name         : result.Name,
		UserAvatarId : result.UserAvatarId,
		UserId       : result.UserId,
		RoomId       : result.Id,
		Room         : result.Id,
		Position     : result.Id,
		Avatar       : response.AvatarResponse{
						Id: result.UserAvatar.Id,
						Image: result.UserAvatar.Image,
						Name: result.UserAvatar.Name,
						},
		
	},nil
}