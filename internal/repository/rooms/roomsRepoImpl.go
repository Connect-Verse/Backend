package rooms

import (
	"github.com/connect-verse/internal/models"
	"gorm.io/gorm"
)

type RoomsImplementation struct{
	db *gorm.DB
}

func  NewRoomImplementation(db *gorm.DB) (*RoomsImplementation){
	 
     return &RoomsImplementation{db}
}

func (r *RoomsImplementation) CreateRoom(room models.Rooms) error{
	result:= r.db.Create(&room)

	if result.Error!=nil {
		return result.Error
	}
	
	return nil
}

func (r *RoomsImplementation) DeleteRoom(roomId string) error{

	result:= r.db.Find(&models.Rooms{},roomId)

	if result.Error!=nil {
		return result.Error
	}
	
	return nil
}

// func (r *RoomsImplementation) JoinRoom(roomId string, userId string) error{

// 	result:= r.db.Find(&models.Rooms{}).Updates()
// 	if result.Error!=nil {
// 		return result.Error
// 	}
	
// 	return nil
// }

func (r *RoomsImplementation) LeaveRoom(roomId  string, userId string) error{
	
	result:= r.db.Find(&models.MetaUsers{}).Where("roomId= ? AND userID = ?", roomId, userId)

	if result.Error!=nil {
		return result.Error
	}
	
	return nil
}

func (r *RoomsImplementation) UserCreatedRoom(userId string) error{
	var rooms models.Rooms
	result:= r.db.Find(&rooms).Where("userId = ? ", userId)

	if result.Error!=nil {
		return result.Error
	}
	
	return nil
}

func (r *RoomsImplementation) AllRoom()  (rooms []models.Rooms, err error){
	var room []models.Rooms
	result:= r.db.Find(&room)

	if result.Error!=nil {
		return nil,result.Error
	}
	
	return room,nil
}
