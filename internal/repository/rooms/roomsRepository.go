package rooms

import (
	"github.com/connect-verse/internal/models"
)

type RoomsRepository interface{
	CreateRoom(user models.Rooms) (models.Rooms,error)
	DeleteRoom(roomId string) (models.Rooms,error)
	// JoinRoom(roomId string, metaId string) (error)
	// LeaveRoom(roomId string, metaId string) (error)
	MyRooms(userId string) ([]models.Rooms,error)
	AllRoom() ([]models.Rooms,  error)
    FindById(roomId string) (models.Rooms, error)
	//rest join queries need to be added
}