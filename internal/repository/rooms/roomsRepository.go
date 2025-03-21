package rooms

import (
	"github.com/connect-verse/internal/models"
)

type RoomsRepository interface{
	CreateRoom(user models.Rooms) error
	DeleteRoom(roomId string) (error)
	JoinRoom(roomId string, metaId string) (error)
	LeaveRoom(roomId string, metaId string) (error)
	UserCreatedRoom(userId string) (rooms []models.Rooms ,err error)
	AllRoom() (rooms []models.Rooms, err error)

	//rest join queries need to be added
}