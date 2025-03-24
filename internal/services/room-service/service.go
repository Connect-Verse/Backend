// package roomservice

// import (
// 	"errors"

// 	"github.com/connect-verse/internal/data/request"
// 	"github.com/connect-verse/internal/data/response"
// 	"github.com/connect-verse/internal/models"
// 	"github.com/connect-verse/internal/repository/Rooms"
// 	"github.com/connect-verse/internal/repository/user"
// 	"github.com/go-playground/validator/v10"
// )

// type RoomServiceImpl struct{
// 	RoomRepo     Rooms.RoomsRepository
// 	validate    *validator.Validate
// }

// func NewRoomServiceImpl(RoomService Rooms.RoomsRepository, validate *validator.Validate) (*RoomServiceImpl,error){
// 	if validate==nil{
//         return &RoomServiceImpl{},errors.New("no validator is provided")
// 	}

// 	return &RoomServiceImpl{
// 		RoomRepo: RoomService,
// 		validate: validate,
// 	},nil
// }


//   func (m *RoomServiceImpl) CreateRoom(Room request.RoomRequest) (response.RoomResponse,error){
     
	
// 	// findUser,err:= user.UserRepository.FindbyId(&user.UserImplementation{},Room.UserId)
	
// 	// if err!=nil {
// 	// 	return response.RoomResponse{},err
// 	// }
	
// 	result,err:= m.RoomRepo.CreateRoom(models.Rooms{
// 		Name: Room.Name,
// 		CreatedBy: Room.CreatedBy,
// 		MapId: Room.MapId,
// 		Maps:models.Maps{
// 		 Image: Room.Maps.Image,
// 		 Info: Room.Maps.Info,
// 		 Tiles: Room.Maps.Tiles,
// 		},
// 		UsersJoined: []models.User{
//          findUser,
// 		},
// 		MetaUsers: []models.MetaUsers{},
// 	 })

// 	 if err!=nil {
// 		return response.RoomResponse{},err
// 	 }

// 	 return response.RoomResponse{
// 		Id: result.Id,
// 		Image: result.Image,
// 		Info: result.Info,
// 		Tiles: result.Tiles,
// 	 },nil
//   }


//   func (m *RoomServiceImpl) DeleteRoom(RoompId string) (response.RoomResponse,error){
// 	result,err:= m.RoomRepo.DeleteRoom(RoompId)

// 	 if err!=nil {
// 		return response.RoomResponse{},err
// 	 }

// 	 return response.RoomResponse{
// 		Id: result.Id,
// 		Image: result.Image,
// 		Info: result.Info,
// 		Tiles: result.Tiles,
// 	 },nil
//   }


//   func (m *RoomServiceImpl) FindRoom(RoompId string) (response.RoomResponse,error){
// 	result,err:= m.RoomRepo.FindRoom(RoompId)

// 	 if err!=nil {
// 		return response.RoomResponse{},err
// 	 }

// 	 return response.RoomResponse{
// 		Id: result.Id,
// 		Image: result.Image,
// 		Info: result.Info,
// 		Tiles: result.Tiles,
// 	 },nil
//   }


//   func (m *RoomServiceImpl) FindAllRooms() ([]response.RoomResponse,error){
// 	result,err:= m.RoomRepo.FindAllRooms()
// 	Roomps:=[]response.RoomResponse{}
// 	 if err!=nil {
// 		return Roomps,err
// 	 }
// 	 for _,value:= range result {
// 		Roomp:=response.RoomResponse{
// 			Id: value.Id,
// 			Image: value.Image,
// 			Info: value.Info,
// 			Tiles: value.Tiles,
// 		 }
// 		Roomps= append(Roomps,Roomp)
// 	 }
// 	 return Roomps,nil
//   }

package roomservice