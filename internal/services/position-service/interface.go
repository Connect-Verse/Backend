package positionservice

import "github.com/connect-verse/internal/models"

type PoService interface{
	FindPostion(models.PlayerPosition) (models.PlayerPosition,error)
	SetPosition(models.PlayerPosition) (models.PlayerPosition,error)
} 
