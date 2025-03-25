package roomservice

import (
	"errors"
	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/repository/rooms"
	"github.com/go-playground/validator/v10"
     
    )

type RoomServiceImpl struct{
	RoomRepo     rooms.RoomsRepository
	validate    *validator.Validate
}

func NewRoomServiceImpl(RoomService rooms.RoomsRepository, validate *validator.Validate) (*RoomServiceImpl,error){
	if validate==nil{
        return &RoomServiceImpl{},errors.New("no validator is provided")
	}

	return &RoomServiceImpl{
		RoomRepo: RoomService,
		validate: validate,
	},nil
}


  func (m *RoomServiceImpl) CreateRoom(Room request.RoomRequest) (response.RoomResponse,error){  
	
	result,err:= m.RoomRepo.CreateRoom(models.Rooms{
		Name: Room.Name,
		CreatedBy: Room.CreatedBy,
		MapId: Room.MapId,
	 })


	 if err!=nil {
		return response.RoomResponse{},err
	 }

	 return response.RoomResponse{
		Id: result.Id,
		CreatedBy: result.CreatedBy,
		MapId: result.MapId,
		UsersJoined: result.UsersJoined,
		Maps:response.MapResponse{
			Id: result.Map.Id,
			Image: result.Map.Image,
			Info: result.Map.Info,
			Tiles: result.Map.Tiles,
		},
		MetaUsers:[]string{"Alice", "Bob"},
	 },nil
  }


  func (m *RoomServiceImpl) DeleteRoom(RoomId string) (response.RoomResponse,error){
	result,err:= m.RoomRepo.DeleteRoom(RoomId)

	 if err!=nil {
		return response.RoomResponse{},err
	 }

	 return response.RoomResponse{
		Id: result.Id,
		CreatedBy: result.CreatedBy,
		MapId: result.MapId,
		UsersJoined: result.UsersJoined,
		Maps:response.MapResponse{
			Id: result.Map.Id,
			Image: result.Map.Image,
			Info: result.Map.Info,
			Tiles: result.Map.Tiles,
		},
		MetaUsers:[]string{"Alice", "Bob"},
	 },nil
  }


  func (m *RoomServiceImpl) MyRoom(usersId string) ([]response.RoomResponse,error){
	result,err:= m.RoomRepo.MyRooms(usersId)

	 if err!=nil {
		return []response.RoomResponse{},err
	 }
     var Rooms []response.RoomResponse
	 for _,value:= range result {
		Room:=response.RoomResponse{
			Id: value.Id,
			CreatedBy: value.CreatedBy,
			MapId: value.MapId,
			UsersJoined: value.UsersJoined,
			Maps:response.MapResponse{
				Id: value.Map.Id,
				Image: value.Map.Image,
				Info: value.Map.Info,
				Tiles: value.Map.Tiles,
			},
			MetaUsers:[]string{"Alice", "Bob"},
		}
		Rooms= append(Rooms,Room)
	 }
	 return Rooms,nil
  }


  func (m *RoomServiceImpl) FindAllRooms() ([]response.RoomResponse,error){
	result,err:= m.RoomRepo.AllRoom()
	Rooms:=[]response.RoomResponse{}
	 if err!=nil {
		return Rooms,err
	 }
	 for _,value:= range result {
		Room:=response.RoomResponse{
			Id: value.Id,
			CreatedBy: value.CreatedBy,
			MapId: value.MapId,
			UsersJoined: value.UsersJoined,
			Maps:response.MapResponse{
				Id: value.Map.Id,
				Image: value.Map.Image,
				Info: value.Map.Info,
				Tiles: value.Map.Tiles,
			},
			MetaUsers:[]string{"Alice", "Bob"},
		}
		Rooms= append(Rooms,Room)
	 }
	 return Rooms,nil
  }

