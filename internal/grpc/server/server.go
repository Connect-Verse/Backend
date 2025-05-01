package server

import (
	"context"
	"sync"
	pb "github.com/connect-verse/internal/grpc"
)

type Remoteserver struct {
	pb.UnimplementedRemoteServerServer
	Mu sync.Mutex
	
}

func (s *Remoteserver) SetPositions(ctx context.Context, position *pb.PlayerPosition) (*pb.QueryReply, error){
    
}


func (s *Remoteserver) CheckPositions(ctx context.Context , id *pb.MetaId) (*pb.PositionResponse, error) {

}  


