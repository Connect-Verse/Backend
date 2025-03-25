package response

import (
	"database/sql"
	"time"

)

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Err     string        `json:"err"`
}

type UserResponse struct {
	Id  	string          `json:"id"`
	Name 	sql.NullString  `json:"name"`
	Email 	string			`json:"email"`
	CreatedAt time.Time		`json:"createdAt"`
	UpdatedAt time.Time	 	`json:"updatedAt"`
	CreatedRooms []RoomResponse  `json:"createdRoom"`
	JoinedRooms  []RoomResponse  `json:"joinedRoom"`
}

type MapResponse struct{
	Id   string     `json:"id"`
	Image string    `json:"image"`
 	Tiles string 	`json:"tiles"`
	Info  string    `json:"info"`
}

type RoomResponse struct{
	Id   string  	 	`json:"id"`
	Name string  	 	`json:"name"`
	CreatedBy string 	`json:"createdby"`
	MapId  string    	`json:"mapId"`
	UsersJoined any  	`json:"userJoined"`
	Maps     MapResponse `json:"maps"`
	MetaUsers []string   `json:"metaUsers"`
  }

type AvatarResponse struct {
	Id string  		`json:"id"`
	Name string 	`json:"name"`
	Image string 	`json:"image"`
	ExistedFrom time.Time `json:"existedFrom"`
}

type MetaUserResponse struct {
	Id string     `json:"id"`
	Name string    `json:"name"`
	UserAvatarId string `json:"userAvatarId"`
	Avatar AvatarResponse    `json:"avatar"`
	UserId string    `json:"userId"`
	RoomId string    `json:"roomId"`
	Room string    `json:"room"`
	Position string    `json:"position"`
}