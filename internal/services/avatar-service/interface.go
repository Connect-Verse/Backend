package avatarservice

import (
	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
)

type AvatarService interface{

	CreateAvatar(avatar request.AvatarRequest) (response.AvatarResponse,error)
	DeleteAvatar(avatarId string) (response.AvatarResponse,error)
	FindAvatar(avatarId string) (response.AvatarResponse,error)
	FindAllAvatars() ([]response.AvatarResponse,error)
	UpdateAvatars(avatar request.AvatarRequest) (response.AvatarResponse,error)

}