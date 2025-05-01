package positionservice

import "github.com/connect-verse/internal/models"

type PoService interface{
	SetPosition(models.PlayerPosition) (models.PlayerPosition,error)
	FindPosition(metaId string) (models.PlayerPosition,error)
} 
