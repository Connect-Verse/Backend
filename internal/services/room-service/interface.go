package roomservice

import (
	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
)

type RoomService interface{
    CreateRoom(Roomp request.RoomRequest) (response.RoomResponse,error)
	DeleteRoom(RoompId string) (response.RoomResponse,error)
	MyRoom(usersId string) ([]response.RoomResponse,error)
	FindAllRooms() ([]response.RoomResponse,error)
}