package positionservice

import (
	"errors"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/repository/position"
	"github.com/go-playground/validator/v10"
)


type PositionServiceImpl struct{
	Pos position.PositionRepository
	Validate   *validator.Validate
}

func NewPosServiceImpl (validate *validator.Validate, pos position.PositionRepository) (*PositionServiceImpl,error){
    if validate!=nil {
		return &PositionServiceImpl{
		   Validate: validate,
		   Pos: pos,
		}, nil
	}
	return &PositionServiceImpl{},errors.New("no validator is provided in Position service implementation")
}

func (p *PositionServiceImpl) SetPosition(position models.PlayerPosition) (models.PlayerPosition,error){
   
	err:= p.Validate.Struct(position)
    
	if err!=nil {
		return models.PlayerPosition{},err
	}

   result,err:= p.Pos.SetPosition(position)
   if err!=nil {
      return result,err
   }
   return result,nil
}

func (p *PositionServiceImpl) FindPosition(metaId string) (models.PlayerPosition,error){
	result,err:= p.Pos.FindPosition(metaId)
	if err!=nil {
	   return result,err
	}
	return result,nil
}

