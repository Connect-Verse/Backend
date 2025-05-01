package server

import (
	"context"
	"sync"

	pb "github.com/connect-verse/internal/grpc"
	"github.com/connect-verse/internal/models"
	positionservice "github.com/connect-verse/internal/services/position-service"
)

type Remoteserver struct {
	pb.UnimplementedRemoteServerServer
	PoService positionservice.PoService
	Mu sync.Mutex
	
}

func (s *Remoteserver) SetPositions(ctx context.Context, position *pb.PlayerPosition) (*pb.QueryReply, error){
    positionModel:= models.PlayerPosition{
		MetaUsersId: position.MetaId,
		X_cordinate: position.XPosition,
		Y_cordinate: position.YPosition,
	}
	_,err:= s.PoService.SetPosition(positionModel)
	if err!=nil {
		return &pb.QueryReply{
		Status: 400,     
        Respose:"service erro occurred while sending executing the function ",
		},err
	}

	return &pb.QueryReply{
		Status: 200,     
        Respose:"successfully saved the given position credentials ",
	}, nil
}


func (s *Remoteserver) CheckPositions(ctx context.Context , id *pb.MetaId) (*pb.PositionResponse, error) {
	result,err:= s.PoService.FindPosition(id.Id)
	if err!=nil {
		return &pb.PositionResponse{
		Status: 400,     
        Response:"service erro occurred while sending executing the function ",
		XPosition: result.X_cordinate,
		YPosition: result.Y_cordinate,
		MetaId: result.MetaUsersId,
		},err
	}

	return &pb.PositionResponse{
		Status: 400,     
        Response:"service erro occurred while sending executing the function ",
		XPosition: result.X_cordinate,
		YPosition: result.Y_cordinate,
		MetaId: result.MetaUsersId,
		},nil
}  


