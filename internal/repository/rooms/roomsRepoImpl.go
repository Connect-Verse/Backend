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

func (r *RoomsImplementation) CreateRoom(room models.Rooms) (models.Rooms,error){
	result:= r.db.Create(&room)

	if result.Error!=nil {
		return models.Rooms{},result.Error
	}
	
	return room,nil
}

func (r *RoomsImplementation) DeleteRoom(roomId string) (models.Rooms,error){
    var room models.Rooms
	result:= r.db.Where("id=?",roomId).Delete(&room)

	if result.Error!=nil {
		return room , result.Error
	}
	
	return room,nil
}

// func (r *RoomsImplementation) JoinRoom(roomId string, userId string) error{

// 	result:= r.db.Find(&models.Rooms{}).Updates()
// 	if result.Error!=nil {
// 		return result.Error
// 	}
	
// 	return nil
// }

// func (r *RoomsImplementation) LeaveRoom(roomId  string, userId string) error{
	
// 	result:= r.db.Find(&models.MetaUsers{}).Where("roomId= ? AND userID = ?", roomId, userId)

// 	if result.Error!=nil {
// 		return result.Error
// 	}
	
// 	return nil
// }

func (r *RoomsImplementation) MyRooms(userId string) ([]models.Rooms,  error){
	var rooms []models.Rooms
	result:= r.db.Find(&rooms).Where("user_id = ? ", userId)

	if result.Error!=nil {
		return rooms,result.Error
	}
	
	return rooms,nil
}

func (r *RoomsImplementation) AllRoom()  ( []models.Rooms,  error){
	var room []models.Rooms
	result:= r.db.Find(&room)

	if result.Error!=nil {
		return nil,result.Error
	}
	
	return room,nil
}
