package request

import (
	"database/sql"
	"time"
)

type CreateUserRequest struct {
Name     sql.NullString `validate:"required,min=1,max=200" json:"name"`
Email    string `validate:"required,min=1,max=50" json:"email"`
Password string `validate:"required,min=1,max=20" json:"password"`
}   
   
type MapRequest struct {
 Image          string `validate:"required,max=200,min=1" json:"image"`
 Tiles	        string `validate:"required,max=200,min=1" json:"tiles"`
 Info		    string `validate:"required,max=200,min=1" json:"info"`
}

type AvatarRequest struct{
	Name    string `validate:"required,max=200,min=1" json:"name"`
    Image   string `validate:"required,max=200,min=1" json:"image"`
	ExistedFrom time.Time `validate:"required,max=200,min=1" json:"existedFrom"`
}

type RoomRequest struct{
	Name    string `validate:"required,max=200,min=1" json:"name"`
    CreatedBy  string `validate:"required,max=200,min=1" json:"createdBy"`
	MapId string `validate:"required,max=200,min=1" json:"mapId"`
    UserIds []string  `validate:"max=200,min=1" json:"usersIds"`   
	MetaUsersIds []string  `validate:"max=200,min=1" json:"metaUsersIds"`
}

type MetaUser struct {
	Name    string `validate:"required,max=200,min=1" json:"name"`
    UserAvatarId  string `validate:"required,max=200,min=1" json:"userAvatarId"`
	RoomId string `validate:"required,max=200,min=1" json:"roomId"`
    UserId string  `validate:"required,max=200,min=1" json:"userId"`   
}